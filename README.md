# GitHub-User-Activity-CLI

## Overview

This project is a simple command-line interface (CLI) tool that fetches and displays the recent activity of a GitHub user. It interacts with the GitHub API to retrieve user events and displays them in a human-readable format in the terminal.

## Features

- Fetch the recent activity of a specified GitHub user.
- Display various types of GitHub events such as push events, issues, and star events.
- Handle errors gracefully, such as invalid usernames or API failures.

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Carl0sL0pez03/GitHub-User-Activity-CLI
    cd GitHub-User-Activity-CLI
    ```

2. Build the project using the provided `Makefile`:
    ```sh
    make build
    ```

This will create an executable named `github-activity`.

## Usage

Run the CLI with a GitHub username as an argument:

```sh
./github-activity <username>
```

# Challenge roadmap

- https://roadmap.sh/projects/github-user-activity