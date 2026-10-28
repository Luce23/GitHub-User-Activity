package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	defaultBaseURL     = "https://api.github.com"
	defaultHTTPTimeout = 10 * time.Second
)

type GitHubClient struct {
	httpClient *http.Client
	baseURL    string
}

func newGitHubClient() *GitHubClient {
	return &GitHubClient{
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultHTTPTimeout,
		},
	}
}

func (client *GitHubClient) newActivityRequest(username string) (*http.Request, error) {
	endpoint := fmt.Sprintf("%s/users/%s/events/public", client.baseURL, username)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("create GitHub request: %w", err)
	}

	request.Header.Set("Accept", "application/vnd.github+json")
	request.Header.Set("User-Agent", "github-user-activity-cli")
	request.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	return request, nil
}
