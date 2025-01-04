package utils

import (
	"fmt"

	"github.com/alielmi98/GitHub-User-Activity-Go-Implementation/github"
)

// FormatEvent formats a GitHub event into a human-readable string.
func FormatEvent(event github.GitHubEvent) string {
	switch event.Type {
	case "PushEvent":
		payload := event.Payload.(map[string]interface{})
		commits, ok := payload["commits"].([]interface{})
		if !ok {
			return "- Error processing PushEvent"
		}
		commitCount := len(commits)
		return fmt.Sprintf("- Pushed %d commits to %s", commitCount, event.Repo.Name)
	case "IssuesEvent":
		payload := event.Payload.(map[string]interface{})
		action, ok := payload["action"].(string)
		if !ok {
			return "- Error processing IssuesEvent"
		}
		return "- " + action + " an issue in " + event.Repo.Name
	case "IssueCommentEvent":
		return "- New comment on an issue in " + event.Repo.Name
	case "WatchEvent":
		return "- Starred " + event.Repo.Name
	default:
		return "- " + event.Type + " event in " + event.Repo.Name
	}
}
