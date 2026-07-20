package ai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCommitMessage_Success(t *testing.T) {
	cannedResponse := `{
		"choices": [
			{
				"message": {"role":"assistant","content":"feat: add login form\n\nAdds a login form with email and password fields."}
			}
		]
	}`

	var capturedBody string
	var capturedAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(buf)
		capturedBody = string(buf)
		capturedAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(cannedResponse))
	}))
	defer srv.Close()

	client, err := New(Options{
		Endpoint: srv.URL,
		Model:    "gpt-4o-mini",
		ApiKey:   "sk-test",
		Timeout:  5 * time.Second,
	})
	assert.NoError(t, err)

	result, err := client.GenerateCommitMessage(context.Background(), "diff --git a/foo b/foo\n+hello")
	assert.NoError(t, err)
	assert.Equal(t, "feat: add login form", result.Summary)
	assert.Equal(t, "Adds a login form with email and password fields.", result.Description)
	assert.Contains(t, capturedBody, `"gpt-4o-mini"`)
	assert.Equal(t, "Bearer sk-test", capturedAuth)
}

func TestGenerateCommitMessage_NoStagedDiff(t *testing.T) {
	client, err := New(Options{
		Endpoint: "http://localhost:1",
		Model:    "gpt-4o-mini",
		ApiKey:   "sk-test",
	})
	assert.NoError(t, err)

	_, err = client.GenerateCommitMessage(context.Background(), "")
	assert.ErrorIs(t, err, ErrEmptyDiff)
}

func TestGenerateCommitMessage_EmptyResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"choices":[]}`))
	}))
	defer srv.Close()

	client, err := New(Options{
		Endpoint: srv.URL,
		Model:    "gpt-4o-mini",
		ApiKey:   "sk-test",
		Timeout:  5 * time.Second,
	})
	assert.NoError(t, err)

	_, err = client.GenerateCommitMessage(context.Background(), "diff --git ...")
	assert.ErrorIs(t, err, ErrEmptyResult)
}

func TestGenerateCommitMessage_4xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":{"message":"unauthorized"}}`))
	}))
	defer srv.Close()

	client, err := New(Options{
		Endpoint: srv.URL,
		Model:    "gpt-4o-mini",
		ApiKey:   "sk-test",
		Timeout:  5 * time.Second,
	})
	assert.NoError(t, err)

	_, err = client.GenerateCommitMessage(context.Background(), "diff --git ...")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "401")
	}
}

func TestGenerateCommitMessage_5xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"message":"boom"}}`))
	}))
	defer srv.Close()

	client, err := New(Options{
		Endpoint: srv.URL,
		Model:    "gpt-4o-mini",
		ApiKey:   "sk-test",
		Timeout:  5 * time.Second,
	})
	assert.NoError(t, err)

	_, err = client.GenerateCommitMessage(context.Background(), "diff --git ...")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "500")
	}
}

func TestGenerateCommitMessage_Timeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	}))
	defer srv.Close()

	client, err := New(Options{
		Endpoint: srv.URL,
		Model:    "gpt-4o-mini",
		ApiKey:   "sk-test",
		Timeout:  100 * time.Millisecond,
	})
	assert.NoError(t, err)

	_, err = client.GenerateCommitMessage(context.Background(), "diff --git ...")
	if assert.Error(t, err) {
		assert.True(t,
			strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "deadline"),
			"expected timeout/deadline error, got: %v", err)
	}
}

func TestTruncateDiff(t *testing.T) {
	short := "abc"
	assert.Equal(t, short, truncateDiff(short, 100))

	big := strings.Repeat("x", 200)
	out := truncateDiff(big, 50)
	assert.Len(t, out, 50)
}

func TestResolveAPIKey_EnvPriority(t *testing.T) {
	t.Setenv("LAZYGIT_AI_API_KEY", "from-env")
	got := ResolveAPIKey("LAZYGIT_AI_API_KEY", "from-config")
	assert.Equal(t, "from-env", got)
}

func TestResolveAPIKey_ConfigFallback(t *testing.T) {
	t.Setenv("LAZYGIT_AI_API_KEY", "")
	got := ResolveAPIKey("LAZYGIT_AI_API_KEY", "from-config")
	assert.Equal(t, "from-config", got)
}

func TestResolveAPIKey_AllEmpty(t *testing.T) {
	t.Setenv("LAZYGIT_AI_API_KEY", "")
	got := ResolveAPIKey("LAZYGIT_AI_API_KEY", "")
	assert.Equal(t, "", got)
}

func TestNew_RequiresEndpoint(t *testing.T) {
	_, err := New(Options{Model: "x", ApiKey: "y"})
	assert.Error(t, err)
}

func TestNew_RequiresModel(t *testing.T) {
	_, err := New(Options{Endpoint: "http://x", ApiKey: "y"})
	assert.Error(t, err)
}
