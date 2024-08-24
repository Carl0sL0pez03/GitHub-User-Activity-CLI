package main

import (
	"github-activity/github"
	"testing"
)

func TestFetchGitHubActivity(t *testing.T) {
	/* User not exits */
	_, err := github.FetchGitHubActivity("notExistUserInGithub")
	if err == nil {
		t.Error("Expected and error for notExistUserInGithub user, got nil")
	}

	/* User valid */
	_, err = github.FetchGitHubActivity("kamranahmedse")
	if err != nil {
		t.Errorf("Expected not error for valid user, got %v", err)
	}

}
