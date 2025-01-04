package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GitHubEvent data structure for GitHub events
type GitHubEvent struct {
	Type      string      `json:"type"`
	Repo      Repo        `json:"repo"`
	Payload   interface{} `json:"payload"` // Using interface{} for flexibility
	CreatedAt string      `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

// GetEvents retrieves recent events for a GitHub user.
func GetEvents(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	print(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API error: status code %d", resp.StatusCode)
	}

	var events []GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("json decoding failed: %w", err)
	}

	return events, nil
}
