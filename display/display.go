package display

import (
	"fmt"
	"github-activity/github"

)

const (
	EventPush   = "PushEvent"
	EventIssues = "IssuesEvent"
	EventWatch  = "WatchEvent"
)

/* Outputs the GitHub events to the terminal */
func DisplayActivity(events []github.GitHubEvent) {
	for _, event := range events {
		switch event.Type {
		case EventPush:
			fmt.Printf("Pushed to %s at %s\n", event.Repo.Name, event.CreatedAt)
		case EventIssues:
			fmt.Printf("Opened an issue in %s at %s\n", event.Repo.Name, event.CreatedAt)
		case EventWatch:
			fmt.Printf("Starred %s at %s\n", event.Repo.Name, event.CreatedAt)
		default:
			fmt.Printf("Performed %s in %s at %s\n", event.Type, event.Repo.Name, event.CreatedAt)
		}
	}
}
