package github

import "time"

/* Represent a single event fetched from the GitHub API */
type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
}
