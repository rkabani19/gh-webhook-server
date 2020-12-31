package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// TODO: change to struct
	webhookData := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&webhookData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("acquired webhook payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s: %v\n", k, v)
	}
}
