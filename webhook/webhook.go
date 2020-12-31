package webhook

import (
	"io"
	"log"
	"net/http"
	"os"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(os.Stdout, r.Body); err != nil {
		log.Println(err)
		return
	}
}
