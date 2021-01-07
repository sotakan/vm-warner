package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Provide Webhook URL via flag
	webhookFlag := flag.String("slackurl", "", "Slack webhook URL")
	vmName := flag.String("instance-name", "unnamed-instance", "Name of instance for text body")
	flag.Parse()
	webhookURL := *webhookFlag

	if webhookURL == "" {
		log.Fatal("No URL for Slack webhook provided. Flag -h to see usage")
	}

	// POSTing
	body := strings.NewReader(fmt.Sprintf(`payload={"text": "%v has been left running! Make sure to shutdown your VMs"}`, *vmName))
	req, err := http.NewRequest("POST", webhookURL, body)
	if err != nil {
		log.Fatal("Error with request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error POSTing")
	}
	defer resp.Body.Close()
}
