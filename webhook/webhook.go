package webhook

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
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
		fmt.Println("Push event.")
	case *github.PullRequestEvent:
		fmt.Println("Pull Request event.")
	case *github.WatchEvent:
		if e.Action != nil && *e.Action == "starred" {
			fmt.Printf("%s starred repository %s\n", *e.Sender.Login, *e.Repo.FullName)
		}
	default:
		log.Printf("Event type %s unknown\n", e)
	}
}
