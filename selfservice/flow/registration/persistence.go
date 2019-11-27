package registration

import (
	"context"
	"testing"

	"github.com/gobuffalo/uuid"
	"github.com/stretchr/testify/require"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/x"
)

type RequestPersister interface {
	CreateRegistrationRequest(context.Context, *Request) error
	GetRegistrationRequest(context.Context, uuid.UUID) (*Request, error)
	UpdateRegistrationRequest(context.Context, uuid.UUID, identity.CredentialsType, *RequestMethod) error
}

type RequestPersistenceProvider interface {
	RegistrationRequestPersister() RequestPersister
}

func TestRequestPersister(p RequestPersister) func(t *testing.T) {
	return func(t *testing.T) {
		// nbr := func() *Request {
		// 	return &Request{
		// 		ID:             uuid.New().String(),
		// 		IssuedAt:       time.Now().UTC().Round(time.Second),
		// 		ExpiresAt:      time.Now().Add(time.Hour).UTC().Round(time.Second),
		// 		RequestURL:     "https://www.ory.sh/request",
		// 		RequestHeaders: http.Header{"Content-Type": {"application/json"}},
		// 		// Disable Active as this value is initially empty (NULL).
		// 		// Active:         identity.CredentialsTypePassword,
		// 		Methods: map[identity.CredentialsType]*CredentialsRequest{
		// 			identity.CredentialsTypePassword: {
		// 				Method: identity.CredentialsTypePassword,
		// 				Config: password.NewRequestMethodConfig(),
		// 			},
		// 			identity.CredentialsTypeOIDC: {
		// 				Method: identity.CredentialsTypeOIDC,
		// 				Config: oidc.NewRequestMethodConfig(),
		// 			},
		// 		},
		// 	}
		// }
		//
		// assertUpdated := func(t *testing.T, expected, actual Request) {
		// 	assert.EqualValues(t, identity.CredentialsTypePassword, actual.Active)
		// 	assert.EqualValues(t, "bar", actual.Methods[identity.CredentialsTypeOIDC].Config.(*oidc.RequestMethodConfig).Action)
		// 	assert.EqualValues(t, "foo", actual.Methods[identity.CredentialsTypePassword].Config.(*password.RequestMethodConfig).Action)
		// }

		t.Run("case=should error when the registration request does not exist", func(t *testing.T) {
			_, err := p.GetRegistrationRequest(context.Background(), x.NewUUID())
			require.Error(t, err)
		})

		// r := LoginRequest{Request: nbr()}
		// require.NoError(t, p.CreateLoginRequest(context.Background(), &r))
		//
		// g, err := p.GetLoginRequest(context.Background(), r.ID)
		// require.NoError(t, err)
		// assert.EqualValues(t, r, *g)
		//
		// require.NoError(t, p.UpdateLoginRequest(context.Background(), r.ID, identity.CredentialsTypeOIDC, &oidc.RequestMethod{Action: "bar"}))
		// require.NoError(t, p.UpdateLoginRequest(context.Background(), r.ID, identity.CredentialsTypePassword, &password.RequestMethod{Action: "foo"}))
		//
		// g, err = p.GetLoginRequest(context.Background(), r.ID)
		// require.NoError(t, err)
		// assertUpdated(t, *r.Request, *g.Request)
	}
}
