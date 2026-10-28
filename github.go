package main

import (
	"net/http"
	"time"
)

const defaultHTTPTimeout = 10 * time.Second

type GitHubClient struct {
	httpClient *http.Client
}

func newGitHubClient() *GitHubClient {
	return &GitHubClient{
		httpClient: &http.Client{
			Timeout: defaultHTTPTimeout,
		},
	}
}
