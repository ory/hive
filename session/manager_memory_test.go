package session_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/viper"

	"github.com/ory/herodot"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/hive/driver/configuration"
	"github.com/ory/hive/identity"
	"github.com/ory/hive/internal"
	. "github.com/ory/hive/session"
)

func init() {
	internal.RegisterFakes()
}

func fakeIdentity(t *testing.T, reg Registry) *identity.Identity {
	i := &identity.Identity{
		ID:              uuid.New().String(),
		TraitsSchemaURL: "file://./stub/identity.schema.json",
		Traits:          json.RawMessage(`{}`),
	}

	viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/identity.schema.json")

	_, err := reg.IdentityPool().Create(context.Background(), i)
	require.NoError(t, err)
	return i
}

func TestSessionManager(t *testing.T) {
	conf, reg := internal.NewMemoryRegistry(t)
	sm := NewManagerMemory(conf, reg)

	_, err := sm.Get(context.Background(), "does-not-exist")
	require.Error(t, err)

	var gave Session
	require.NoError(t, faker.FakeData(&gave))
	gave.Identity = fakeIdentity(t, reg)

	require.NoError(t, sm.Create(context.Background(), &gave))

	got, err := sm.Get(context.Background(), gave.SID)
	require.NoError(t, err)
	assert.EqualValues(t, &gave, got)

	require.NoError(t, sm.Delete(context.Background(), gave.SID))
	_, err = sm.Get(context.Background(), gave.SID)
	require.Error(t, err)
}

func TestSessionManagerHTTP(t *testing.T) {
	conf, reg := internal.NewMemoryRegistry(t)
	sm := NewManagerMemory(conf, reg)
	h := herodot.NewJSONWriter(nil)

	var s Session
	require.NoError(t, faker.FakeData(&s))
	s.Identity = fakeIdentity(t, reg)

	router := httprouter.New()
	router.GET("/set", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		require.NoError(t, sm.Create(context.Background(), &s))
		require.NoError(t, sm.SaveToRequest(context.Background(), &s, w, r))
	})

	router.GET("/set-direct", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, err := sm.CreateToRequest(context.Background(), s.Identity, w, r)
		require.NoError(t, err)
	})

	router.GET("/clear", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		require.NoError(t, sm.PurgeFromRequest(context.Background(), w, r))
	})

	router.GET("/get", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		s, err := sm.FetchFromRequest(context.Background(), r)
		if errors.Cause(err) == ErrNoActiveSessionFound {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			assert.NoError(t, err)
			return
		}

		h.Write(w, r, s)
	})

	ts := httptest.NewServer(router)
	defer ts.Close()

	t.Run("case=exists", func(t *testing.T) {
		c := &http.Client{}
		c.Jar, _ = cookiejar.New(nil)

		res, err := c.Get(ts.URL + "/set")
		require.NoError(t, err)
		defer res.Body.Close()
		require.EqualValues(t, http.StatusOK, res.StatusCode)

		res, err = c.Get(ts.URL + "/get")
		require.NoError(t, err)
		defer res.Body.Close()
		require.EqualValues(t, http.StatusOK, res.StatusCode)

		var got Session
		require.NoError(t, json.NewDecoder(res.Body).Decode(&got))
	})

	t.Run("case=exists-direct", func(t *testing.T) {
		c := &http.Client{}
		c.Jar, _ = cookiejar.New(nil)

		res, err := c.Get(ts.URL + "/set-direct")
		require.NoError(t, err)
		defer res.Body.Close()
		require.EqualValues(t, http.StatusOK, res.StatusCode)

		res, err = c.Get(ts.URL + "/get")
		require.NoError(t, err)
		defer res.Body.Close()
		require.EqualValues(t, http.StatusOK, res.StatusCode)

		var got Session
		require.NoError(t, json.NewDecoder(res.Body).Decode(&got))
	})

	t.Run("case=exists_clears", func(t *testing.T) {
		c := &http.Client{}
		c.Jar, _ = cookiejar.New(nil)

		res, err := c.Get(ts.URL + "/set")
		require.NoError(t, err)
		defer res.Body.Close()
		require.EqualValues(t, http.StatusOK, res.StatusCode)

		res, err = c.Get(ts.URL + "/clear")
		require.NoError(t, err)
		defer res.Body.Close()
		require.EqualValues(t, http.StatusOK, res.StatusCode)

		res, err = c.Get(ts.URL + "/get")
		require.NoError(t, err)
		defer res.Body.Close()
		assert.EqualValues(t, http.StatusUnauthorized, res.StatusCode)
	})

	t.Run("case=not-exist", func(t *testing.T) {
		c := &http.Client{}
		c.Jar, _ = cookiejar.New(nil)

		res, err := c.Get(ts.URL + "/get")
		require.NoError(t, err)
		defer res.Body.Close()
		assert.EqualValues(t, http.StatusUnauthorized, res.StatusCode)
	})
}
