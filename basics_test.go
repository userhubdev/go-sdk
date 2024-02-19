package userhub_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/userhubdev/go-sdk/adminapi"
	"github.com/userhubdev/go-sdk/adminv1"
	"github.com/userhubdev/go-sdk/code"
	"github.com/userhubdev/go-sdk/option"
	"github.com/userhubdev/go-sdk/userapi"
)

func getEnv(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		t.Skipf("%s not configured", name)
	}
	return value
}

func TestAdminApi(t *testing.T) {
	t.Parallel()

	adminApi, err := adminapi.New(getEnv(t, "TEST_ADMIN_KEY"))
	require.NoError(t, err)

	res, err := adminApi.Users().List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestUserApi(t *testing.T) {
	t.Parallel()

	userApi, err := userapi.New(getEnv(t, "TEST_USER_KEY"), "")
	require.NoError(t, err)

	session, err := userApi.Session().Get(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, session)
}

func TestModel(t *testing.T) {
	t.Parallel()

	user := &adminv1.User{
		Id:          "usr_1",
		DisplayName: "Jane Doe",
		CreateTime:  time.Now().UTC(),
		AccountConnections: []*adminv1.AccountConnection{
			{ExternalId: "cus_1"},
		},
	}
	require.Equal(t, "Jane Doe", user.DisplayName)

	encodedJson, err := json.Marshal(user)
	require.NoError(t, err)

	data := map[string]any{}
	err = json.Unmarshal(encodedJson, &data)
	require.NoError(t, err)
	require.Contains(t, data, "displayName")
	require.Equal(t, "Jane Doe", data["displayName"])

	var decodedUser *adminv1.User
	err = json.Unmarshal(encodedJson, &decodedUser)
	require.NoError(t, err)

	require.Equal(t, user.Id, decodedUser.Id)
	require.Equal(t, user.DisplayName, decodedUser.DisplayName)
	require.Equal(t, user.CreateTime, decodedUser.CreateTime)
	require.Equal(t, len(user.AccountConnections), len(decodedUser.AccountConnections))
	require.Equal(t, user.AccountConnections[0].ExternalId, decodedUser.AccountConnections[0].ExternalId)
}

func TestApiGet(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer sk_test", r.Header.Get("authorization"))
		require.Equal(t, "/admin/v1/users/usr_1", r.URL.Path)
		require.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(`{"id": "usr_1", "displayName": "Jane Doe"}`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	user, err := adminApi.Users().Get(context.Background(), "usr_1", nil)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "usr_1", user.Id)
	require.Equal(t, "Jane Doe", user.DisplayName)
}

func TestApiPost(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer sk_test", r.Header.Get("authorization"))
		require.Equal(t, "/admin/v1/users", r.URL.Path)
		require.Equal(t, "POST", r.Method)

		var data map[string]any
		err := json.NewDecoder(r.Body).Decode(&data)
		require.NoError(t, err)

		require.Contains(t, data, "displayName")
		require.Len(t, data, 1)
		require.Equal(t, "Jane Doe", data["displayName"])

		_, _ = w.Write([]byte(`{"id": "usr_1", "displayName": "Jane Doe"}`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	user, err := adminApi.Users().Create(context.Background(), &adminapi.UserCreateInput{
		DisplayName: "Jane Doe",
	})
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "usr_1", user.Id)
	require.Equal(t, "Jane Doe", user.DisplayName)
}

func TestApiPost_Empty(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer sk_test", r.Header.Get("authorization"))
		require.Equal(t, "/admin/v1/users", r.URL.Path)
		require.Equal(t, "POST", r.Method)

		var data map[string]any
		err := json.NewDecoder(r.Body).Decode(&data)
		require.NoError(t, err)

		require.Empty(t, data)

		_, _ = w.Write([]byte(`{"id": "usr_1"}`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	user, err := adminApi.Users().Create(context.Background(), &adminapi.UserCreateInput{
		DisplayName: "",
	})
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "usr_1", user.Id)
	require.Empty(t, user.DisplayName)
}

func TestApiPatch(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer sk_test", r.Header.Get("authorization"))
		require.Equal(t, "/admin/v1/users/usr_1", r.URL.Path)
		require.Equal(t, "PATCH", r.Method)

		var data map[string]any
		err := json.NewDecoder(r.Body).Decode(&data)
		require.NoError(t, err)

		require.Contains(t, data, "displayName")
		require.Len(t, data, 1)
		require.Equal(t, "Jane Doe", data["displayName"])

		_, _ = w.Write([]byte(`{"id": "usr_1", "displayName": "Jane Doe"}`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	user, err := adminApi.Users().Update(context.Background(), "usr_1", &adminapi.UserUpdateInput{
		DisplayName: adminapi.Some("Jane Doe"),
	})
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "usr_1", user.Id)
	require.Equal(t, "Jane Doe", user.DisplayName)
}

func TestApiPatch_Empty(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer sk_test", r.Header.Get("authorization"))
		require.Equal(t, "/admin/v1/users/usr_1", r.URL.Path)
		require.Equal(t, "PATCH", r.Method)

		var data map[string]any
		err := json.NewDecoder(r.Body).Decode(&data)
		require.NoError(t, err)

		require.Contains(t, data, "displayName")
		require.Len(t, data, 1)
		require.Equal(t, "", data["displayName"])

		_, _ = w.Write([]byte(`{"id": "usr_1"}`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	user, err := adminApi.Users().Update(context.Background(), "usr_1", &adminapi.UserUpdateInput{
		DisplayName: adminapi.Some(""),
	})
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "usr_1", user.Id)
	require.Empty(t, user.DisplayName)
}

func TestApiDelete(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, "Bearer sk_test", r.Header.Get("authorization"))
		require.Equal(t, "/admin/v1/users/usr_1", r.URL.Path)
		require.Equal(t, "DELETE", r.Method)
		_, _ = w.Write([]byte(`{"id": "usr_1"}`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	user, err := adminApi.Users().Delete(context.Background(), "usr_1", nil)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, "usr_1", user.Id)
}

func TestApiError(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(404)
		_, _ = w.Write([]byte(`{"code":"NOT_FOUND", "message":"Not Found", "metadata":{}}`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	_, err = adminApi.Users().Get(context.Background(), "usr_1", nil)
	require.EqualError(t, err, "userhub: Not Found (call: admin.users.get, apiCode: NOT_FOUND)")

	var apiError *adminapi.Error
	require.True(t, errors.As(err, &apiError))
	require.Equal(t, code.NotFound, apiError.ApiCode())
	require.Equal(t, "Not Found", apiError.Message())
	require.Equal(t, 404, apiError.StatusCode())
}

func TestApiRateLimited(t *testing.T) {
	if os.Getenv("CI") == "" {
		t.Skip("Skipping slow test")
	}

	t.Parallel()

	var attempts int

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++

		w.WriteHeader(429)
		_, _ = w.Write([]byte(`Rate limited`))
	}))
	defer ts.Close()

	adminApi, err := adminapi.New("sk_test", option.WithBaseUrl(ts.URL))
	require.NoError(t, err)

	startTime := time.Now()

	_, err = adminApi.Users().Get(context.Background(), "usr_1", nil)
	require.EqualError(t, err, "userhub: API call rate limited (call: admin.users.get, apiCode: RESOURCE_EXHAUSTED)")

	endTime := time.Now()

	var apiError *adminapi.Error
	require.True(t, errors.As(err, &apiError))
	require.Equal(t, code.ResourceExhausted, apiError.ApiCode())
	require.Equal(t, "API call rate limited", apiError.Message())
	require.Equal(t, 429, apiError.StatusCode())

	require.Equal(t, 6, attempts)
	require.WithinDuration(t, startTime.Add(3*time.Second), endTime, time.Second)
}
