package hash

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"

	"github.com/ory/kratos/driver/config"
)

var ErrUnknownHashAlgorithm = fmt.Errorf("unknown hash algorithm")
var ErrUnknownHashFormat = fmt.Errorf("unknown hash format")
var ErrEmptyHashCompare = fmt.Errorf("empty hash provided")
var ErrEmptyPasswordCompare = fmt.Errorf("empty password provided")

var hashSeparator = []byte("$")
var bcryptLegacyPrefix = regexp.MustCompile(`^2[abzy]?$`)

func Compare(ctx context.Context, cfg *config.Config, password []byte, hash []byte) error {
	algorithm, realHash, err := ParsePasswordHash(hash)
	if err != nil {
		return errors.WithStack(err)
	}

	switch {
	case bytes.Equal(algorithm, Argon2AlgorithmId):
		return CompareArgon2id(ctx, cfg, password, realHash)
	case bytes.Equal(algorithm, BcryptAlgorithmId):
		return CompareBcrypt(ctx, cfg, password, realHash)
	case bytes.Equal(algorithm, BcryptAESAlgorithmId):
		return CompareBcryptAes(ctx, cfg, password, realHash)
	default:
		return ErrUnknownHashAlgorithm
	}
}

func CompareBcrypt(_ context.Context, _ *config.Config, password []byte, hash []byte) error {
	if err := validateBcryptPasswordLength(password); err != nil {
		return errors.WithStack(err)
	}

	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func IsPasswordHash(hash []byte) bool {
	_, _, err := ParsePasswordHash(hash)
	return err == nil
}

func ParsePasswordHash(input []byte) (algorithm, hash []byte, err error) {
	hashParts := bytes.SplitN(input, hashSeparator, 3)
	if len(hashParts) != 3 {
		err = errors.WithStack(ErrUnknownHashFormat)
		return
	}

	hash = append(hashSeparator, hashParts[2]...)
	switch {
	case bytes.Equal(hashParts[1], Argon2AlgorithmId):
		algorithm = Argon2AlgorithmId
		return
	case bcryptLegacyPrefix.Match(hashParts[1]):
		hash = input
		fallthrough
	case bytes.Equal(hashParts[1], BcryptAlgorithmId):
		algorithm = BcryptAlgorithmId
		return
	case bytes.Equal(hashParts[1], BcryptAESAlgorithmId):
		algorithm = BcryptAESAlgorithmId
		return
	default:
		err = ErrUnknownHashAlgorithm
		return
	}
}

// aes256Decrypt decrypts data using 256-bit AES-GCM.  This both hides the content of
// the data and provides a check that it hasn't been altered. Expects input
// form nonce|ciphertext|tag where '|' indicates concatenation.
func aes256Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

func CompareBcryptAes(_ context.Context, cfg *config.Config, password, hash []byte) error {
	if len(hash) == 0 {
		return errors.WithStack(ErrEmptyHashCompare)
	}
	if len(password) == 0 {
		return errors.WithStack(ErrEmptyPasswordCompare)
	}

	decoded := make([]byte, hex.DecodedLen(len(hash[1:])))
	_, err := hex.Decode(decoded, hash[1:])
	if err != nil {
		return errors.WithStack(err)
	}

	var lastError error
	var aesDecrypted []byte
	hasherCfg, err := cfg.HasherBcryptAES()
	if err != nil {
		return errors.WithStack(err)
	}
	for i := range hasherCfg.Key {
		aesDecrypted, lastError = aes256Decrypt(decoded[:], &hasherCfg.Key[i])
		if lastError == nil {
			break
		}
	}

	if lastError != nil {
		return errors.WithStack(lastError)
	}

	sh := sha3.New512()
	sh.Write(password)
	err = bcrypt.CompareHashAndPassword(aesDecrypted, sh.Sum(nil))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func CompareArgon2id(_ context.Context, _ *config.Config, password, hashPart []byte) error {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeArgon2idHash(hashPart)
	if err != nil {
		return errors.WithStack(err)
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey(password, salt, p.Iterations, uint32(p.Memory), p.Parallelism, p.KeyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return nil
	}
	return ErrMismatchedHashAndPassword
}

func decodeArgon2idHash(encodedHash []byte) (p *config.Argon2, salt, hash []byte, err error) {
	parts := strings.Split(string(encodedHash), string(hashSeparator))
	if len(parts) != 5 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(parts[1], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = new(config.Argon2)
	_, err = fmt.Sscanf(parts[2], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(parts[3])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
