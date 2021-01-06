package main

import (
	"fmt"
	"flag"
	"errors"
	"net/http"
	"net/url"
)

func main() {
	// Provide Webhook URL via flag
	webhookFlag := flag.String("slackurl", "", "Slack webhook URL")
	webhookURL := *webhookFlag

	if webhookURL == "" {
		errors.New("No URL for Slack webhook provided. Flag -h to see usage")
	}

	// POSTing
	resp, err := http.PostForm(webhookURL, url.Values{"text": {"VM is up!"}})
}