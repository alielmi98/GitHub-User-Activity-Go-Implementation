package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alielmi98/GitHub-User-Activity-Go-Implementation/github"
	"github.com/alielmi98/GitHub-User-Activity-Go-Implementation/utils"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		fmt.Println("Please provide exactly one argument: the GitHub username.")
		os.Exit(1)
	}

	username := args[0]

	events, err := github.GetEvents(username)
	if err != nil {
		log.Fatalf("Error fetching events: %v", err)
	}

	if len(events) == 0 {
		fmt.Printf("No recent events found for user %s.\n", username)
		return
	}

	for _, event := range events {
		fmt.Println(utils.FormatEvent(event))
	}
}
