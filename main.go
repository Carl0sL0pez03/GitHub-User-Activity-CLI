package main

import (
	"fmt"
	"github-activity/display"
	"github-activity/github"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	username := os.Args[1]

	events, err := github.FetchGitHubActivity(username)
	if err != nil {
		fmt.Println("Error fetching Github activity", err)
	}

	display.DisplayActivity(events)
}
