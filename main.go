package main

import (
	"log"
	"net/http"

	"github.com/rkabani19/gh-webhook-server/webhook"
)

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", webhook.HandleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
