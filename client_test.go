package audited_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/getaudited/audited-go"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string
		baseAPI     string
		apiToken    string
		httpClient  *http.Client
		timeout     time.Duration
	}{
		{
			name:     "new_client",
			baseAPI:  "http://localhost:8080/api",
			apiToken: "super-secure-token",
		},
		{
			name:     "new_client_with_timeout",
			baseAPI:  "http://localhost:8080/api",
			apiToken: "super-secure-token",
			timeout:  time.Second * 5,
		},
		{
			name:       "new_client_with_http_client",
			baseAPI:    "http://localhost:8080/api",
			apiToken:   "super-secure-token",
			httpClient: &http.Client{},
		},
		{
			name:        "error_missing_base_api",
			baseAPI:     "",
			apiToken:    "super-secure-token",
			expectedErr: "audited-go: BaseAPI cannot be empty",
		},
		{
			name:        "error_missing_api_token",
			baseAPI:     "http://localhost:8080/api",
			apiToken:    "",
			expectedErr: "audited-go: APIToken cannot be empty",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			client, err := audited.NewClient(audited.Config{
				BaseAPI:    tc.baseAPI,
				APIToken:   tc.apiToken,
				HttpClient: tc.httpClient,
				Timeout:    tc.timeout,
			})

			if tc.expectedErr != "" {
				require.Nil(t, client)
				require.ErrorContains(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, client)
		})
	}
}

func TestClient_CreateEvent(t *testing.T) {
	event := audited.Event{
		Action:     "user.created",
		Actor:      audited.Actor{},
		Context:    audited.Context{},
		SourceID:   "",
		Targets:    nil,
		Version:    0,
		Metadata:   nil,
		OccurredAt: time.Time{},
	}

	auditedServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		require.NoError(t, err)

		var eventReceived audited.Event
		err = json.Unmarshal(body, &eventReceived)
		require.NoError(t, err)
		require.Equal(t, event, eventReceived)

		w.WriteHeader(http.StatusCreated)
	}))
	defer auditedServer.Close()

	client, err := audited.NewClient(audited.Config{
		BaseAPI:    auditedServer.URL,
		APIToken:   "super-secure-token",
		HttpClient: nil,
		Timeout:    0,
	})
	require.NoError(t, err)

	err = client.CreateEvent(context.Background(), event)
	require.NoError(t, err)
}

func TestClient_CreateEvent_Errors(t *testing.T) {
	testCases := []struct {
		name        string
		buildServer func() *httptest.Server
		expectedErr error
	}{
		{
			name: "error_unauthorized",
			buildServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusUnauthorized)
				}))
			},
			expectedErr: audited.ErrUnauthorized,
		},
		{
			name: "error_not_found",
			buildServer: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}))
			},
			expectedErr: audited.ErrEventTypeNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			auditedServer := tc.buildServer()
			defer auditedServer.Close()

			client, err := audited.NewClient(audited.Config{
				BaseAPI:    auditedServer.URL,
				APIToken:   "super-secure-token",
				HttpClient: nil,
				Timeout:    0,
			})
			require.NoError(t, err)

			err = client.CreateEvent(context.Background(), audited.Event{})
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
