package helpers

import (
	"context"
	"errors"
	"testing"

	"github.com/jesseduffield/lazygit/pkg/ai"
	"github.com/stretchr/testify/assert"
)

func TestTryRemoveHardLineBreaks(t *testing.T) {
	scenarios := []struct {
		name           string
		message        string
		autoWrapWidth  int
		expectedResult string
	}{
		{
			name:           "empty",
			message:        "",
			autoWrapWidth:  7,
			expectedResult: "",
		},
		{
			name:           "all line breaks are needed",
			message:        "abc\ndef\n\nxyz",
			autoWrapWidth:  7,
			expectedResult: "abc\ndef\n\nxyz",
		},
		{
			name:           "some can be unwrapped",
			message:        "123\nabc def\nghi jkl\nmno\n456\n",
			autoWrapWidth:  7,
			expectedResult: "123\nabc def ghi jkl mno\n456\n",
		},
	}
	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			actualResult := TryRemoveHardLineBreaks(s.message, s.autoWrapWidth)
			assert.Equal(t, s.expectedResult, actualResult)
		})
	}
}

type fakeCommitGenerator struct {
	summary     string
	description string
	err         error
}

func (f *fakeCommitGenerator) GenerateCommitMessage(ctx context.Context, diff string) (ai.Result, error) {
	if f.err != nil {
		return ai.Result{}, f.err
	}
	return ai.Result{Summary: f.summary, Description: f.description}, nil
}

func TestApplyDraft_WritesBothFields(t *testing.T) {
	fake := &fakeCommitGenerator{summary: "fix bug", description: "more detail"}
	var gotSummary, gotDescription string

	err := applyDraft(fake, "diff content", context.Background(),
		func(summary, body string) {
			gotSummary = summary
			gotDescription = body
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, "fix bug", gotSummary)
	assert.Equal(t, "more detail", gotDescription)
}

func TestApplyDraft_PropagatesError(t *testing.T) {
	boom := errors.New("network down")
	fake := &fakeCommitGenerator{err: boom}
	called := false

	err := applyDraft(fake, "diff", context.Background(),
		func(_, _ string) { called = true },
	)
	assert.ErrorIs(t, err, boom)
	assert.False(t, called, "apply callback must not run on error")
}
