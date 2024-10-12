package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
}

type Actor struct {
	Login string `json:"login"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action  string   `json:"action"`
	Commits []Commit `json:"commits,omitempty"`
}

type Commit struct {
	Message string `json:"message"`
}

func main() {
	args := os.Args
	argsLen := len(args)

	if argsLen < 1 {
		fmt.Println("Please provide a GitHub username")
		os.Exit(1)
	}

	username := args[1]
	fmt.Println("Username:", username)

	events, err := getEventsFromGithub(username)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, event := range *events {
		if event.Type == "WatchEvent" {
			fmt.Printf("\t# %s %s to watch %s\n", event.Actor.Login, event.Payload.Action, event.Repo.Name)
		} else if event.Type == "PushEvent" {
			commitCount := len(event.Payload.Commits)

			fmt.Printf("\t# %s pushed %d commit to %s\n", event.Actor.Login, commitCount, event.Repo.Name)
		} else if event.Type == "CreateEvent" {
			fmt.Printf("\t# %s created %s\n", event.Actor.Login, event.Repo.Name)
		} else {
			fmt.Printf("\t# %s did %s to %s\n", event.Actor.Login, event.Payload.Action, event.Repo.Name)
		}
	}
}

func getEventsFromGithub(username string) (*[]Event, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/users/"+username+"/events", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var events []Event

	err = json.Unmarshal(respBody, &events)
	if err != nil {
		return nil, err
	}

	return &events, nil
}
