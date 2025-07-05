package urlshortner

import(
	"fmt"
	"net/http"
	"encoding/json"
)

func SetupRoute() {
	http.HandleFunc("/shorturls", UrlHandler)
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	type Message struct {
		Url       string `json:"url"`
		Validity  int    `json:"validity"`
		Shortcode string `json:"shortcode"`
	}

	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Print(msg.Url, msg.Validity, msg.Shortcode)
	fmt.Fprintln(w, "test done")
}