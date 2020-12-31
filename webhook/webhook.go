package webhook

import (
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte("my-secret-key"))
	if err != nil {
		log.Printf("error reading request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	HandleEvent(event)
}

func HandleEvent(event interface{}) {
	switch e := event.(type) {
	case *github.PushEvent:
		log.Println("Push event.")
	case *github.PullRequestEvent:
		log.Println("Pull Request event.")
	case *github.WatchEvent:
		if e.Action != nil && *e.Action == "started" {
			log.Printf("%s starred repository %s\n", *e.Sender.Login, *e.Repo.FullName)
		}
	default:
		log.Printf("Event type %s unknown\n", e)
	}
}
