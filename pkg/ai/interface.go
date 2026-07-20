package ai

import "context"

// CommitGenerator is the narrow surface that the UI helpers depend on.
// Defining it as an interface (rather than depending on the concrete *Client)
// lets callers inject a fake in tests without spinning up an httptest
// server, and keeps the helpers' contract explicit.
type CommitGenerator interface {
	GenerateCommitMessage(ctx context.Context, diff string) (Result, error)
}

// Compile-time guarantee: the production *Client satisfies this interface.
var _ CommitGenerator = (*Client)(nil)
