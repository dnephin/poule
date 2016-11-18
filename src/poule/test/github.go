package test

import (
	"poule/gh"
	"poule/test/mocks"
	"time"

	"github.com/google/go-github/github"
)

// TestClient is a mocked implementation of a GitHub client.
type TestClient struct {
	MockIssues       mocks.IssuesService
	MockPullRequests mocks.PullRequestsService
	MockRepositories mocks.RepositoriesService
	MockSearch       mocks.SearchService
}

// Issues returns the issue service instance.
func (t *TestClient) Issues() gh.IssuesService {
	return &t.MockIssues
}

// PullRequests returns the pull request service instance.
func (t *TestClient) PullRequests() gh.PullRequestsService {
	return &t.MockPullRequests
}

// Repositories returns the repository service instance.
func (t *TestClient) Repositories() gh.RepositoriesService {
	return &t.MockRepositories
}

// Search returns the search service instance.
func (t *TestClient) Search() gh.SearchService {
	return &t.MockSearch
}

// MakeLabel is a helper to create a GitHub label.
func MakeLabel(name string) github.Label {
	return github.Label{
		Name: github.String(name),
	}
}

// MakeStatus is a helper to create a GitHub repository status.
func MakeStatus(context, status string, createdAt time.Time) *github.RepoStatus {
	return &github.RepoStatus{
		Context:   &context,
		CreatedAt: &createdAt,
		State:     &status,
	}
}
