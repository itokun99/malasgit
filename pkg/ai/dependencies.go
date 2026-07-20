package ai

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jesseduffield/lazygit/pkg/config"
)

const aiAPIKeyEnvVar = "AI_API_KEY"

func NewClientFromUserConfig(cfg config.AIConfig) (*Client, error) {
	if !cfg.Enabled {
		return nil, errors.New("AI commit generation is disabled (ai.enabled = false)")
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("AI not configured: %w", err)
	}

	endpoint := strings.TrimRight(cfg.Endpoint, "/chat/completions") + "/chat/completions"
	apiKey := ResolveAPIKey(aiAPIKeyEnvVar, cfg.ApiKey)
	if apiKey == "" {
		return nil, errors.New("AI API key is empty: set " + aiAPIKeyEnvVar + " or ai.apiKey in config")
	}

	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second
	if cfg.TimeoutSeconds <= 0 {
		timeout = 30 * time.Second
	}

	return New(Options{
		Endpoint: endpoint,
		Model:    cfg.Model,
		ApiKey:   apiKey,
		Timeout:  timeout,
	})
}
