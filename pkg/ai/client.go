// Package ai talks to OpenAI-compatible chat-completions endpoints to produce
// commit-message drafts from a staged diff. The client is transport-only: it
// does not know about the UI, the config layer, or the i18n catalog, so it
// stays trivially testable and reusable.
package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Sentinel errors callers can match with errors.Is.
var (
	// ErrEmptyDiff is returned when there is no staged diff to summarize.
	ErrEmptyDiff = errors.New("ai: empty diff")
	// ErrEmptyResult is returned when the endpoint replies successfully but
	// produces no message content (e.g. the model refused the request).
	ErrEmptyResult = errors.New("ai: empty result from endpoint")
	// ErrMissingAPIKey is returned by New when neither the env var nor the
	// configured key is set.
	ErrMissingAPIKey = errors.New("ai: no API key configured (set the env var or ai.apiKey)")
)

// Options is the immutable configuration passed to New. Construct it from
// pkg/config.AIConfig (after ResolveAPIKey has merged the env override).
type Options struct {
	Endpoint string
	Model    string
	ApiKey   string
	Timeout  time.Duration
}

// Result is the parsed output of the chat-completions response. The model is
// prompted to emit the first line as the subject and any subsequent lines as
// the body, with a single blank line separating them.
type Result struct {
	Summary     string
	Description string
}

// Client is a thin wrapper around an http.Client that knows how to build and
// parse OpenAI-compatible chat-completions requests for commit summaries.
type Client struct {
	endpoint string
	model    string
	apiKey   string
	httpc    *http.Client
}

// New validates the endpoint, model, and API key, and returns a Client with a
// bounded HTTP timeout. If Timeout is zero, a 30-second default is used.
func New(opts Options) (*Client, error) {
	if opts.Endpoint == "" {
		return nil, errors.New("ai: endpoint is required")
	}
	if opts.Model == "" {
		return nil, errors.New("ai: model is required")
	}
	if opts.ApiKey == "" {
		return nil, ErrMissingAPIKey
	}
	timeout := opts.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	return &Client{
		endpoint: opts.Endpoint,
		model:    opts.Model,
		apiKey:   opts.ApiKey,
		httpc:    &http.Client{Timeout: timeout},
	}, nil
}

// GenerateCommitMessage sends the staged diff to the configured endpoint and
// returns the parsed subject/body. The supplied context can cancel the request
// or shorten its deadline.
//
// Errors:
//   - ErrEmptyDiff when diff is empty after trimming.
//   - A wrapped error for transport / HTTP-status failures.
//   - ErrEmptyResult when the response carries no message content.
func (c *Client) GenerateCommitMessage(ctx context.Context, diff string) (Result, error) {
	if strings.TrimSpace(diff) == "" {
		return Result{}, ErrEmptyDiff
	}

	body, err := json.Marshal(map[string]any{
		"model": c.model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": systemPrompt,
			},
			{
				"role":    "user",
				"content": diff,
			},
		},
		"temperature": 0.2,
	})
	if err != nil {
		return Result{}, fmt.Errorf("ai: marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(body))
	if err != nil {
		return Result{}, fmt.Errorf("ai: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	resp, err := c.httpc.Do(req)
	if err != nil {
		return Result{}, fmt.Errorf("ai: request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read up to 4 KiB of the body so error messages are useful without
		// risking unbounded memory if the server streams back megabytes.
		raw, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return Result{}, fmt.Errorf("ai: endpoint returned status %d: %s", resp.StatusCode, string(raw))
	}

	var parsed chatResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return Result{}, fmt.Errorf("ai: decode response: %w", err)
	}
	if len(parsed.Choices) == 0 || parsed.Choices[0].Message.Content == "" {
		return Result{}, ErrEmptyResult
	}

	return splitSummary(parsed.Choices[0].Message.Content), nil
}

// systemPrompt instructs the model to follow the conventional-commit subject
// style and to use the body only when the diff is non-trivial.
const systemPrompt = "You write git commit messages. Reply with a single line " +
	"summary in the imperative mood (no trailing period) followed by an " +
	"optional body separated by a blank line. Keep the summary under 72 " +
	"characters. Do not wrap the output in code fences. Reply only with the " +
	"commit message."

// chatResponse is the minimal subset of an OpenAI chat-completions response we
// consume. Unknown fields are ignored by encoding/json.
type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// splitSummary parses "subject\n\nbody" into a Result. If the content has no
// blank line separator the whole content becomes the summary.
func splitSummary(content string) Result {
	content = strings.TrimRight(content, "\n")
	subject, body, _ := strings.Cut(content, "\n\n")
	return Result{
		Summary:     strings.TrimSpace(subject),
		Description: strings.TrimSpace(body),
	}
}

// TruncateDiff caps a diff at maxBytes. Callers use it to pre-truncate large
// diffs to AIConfig.MaxDiffSize before sending them to the endpoint.
func TruncateDiff(diff string, maxBytes int) string {
	if maxBytes <= 0 || len(diff) <= maxBytes {
		return diff
	}
	return diff[:maxBytes]
}

// ResolveAPIKey returns the API key, preferring a non-empty environment
// variable over the config-supplied value. This lets users keep secrets out
// of their dotfiles while still allowing the config file to hold a
// non-secret default.
//
// A missing env var falls through to the config value; an empty-but-set env
// var also falls through (so LAZYGIT_AI_API_KEY= in the shell doesn't
// accidentally clobber the config value). All empty returns "".
func ResolveAPIKey(envName, configValue string) string {
	if v, ok := os.LookupEnv(envName); ok && v != "" {
		return v
	}
	return configValue
}
