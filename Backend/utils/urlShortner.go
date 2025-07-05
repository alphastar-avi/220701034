package urlshortner

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

func SetupRoute() {
    http.HandleFunc("/shorturls", UrlHandler)
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    type Request struct {
        Url       string `json:"url"`
        Validity  int    `json:"validity"`
        Shortcode string `json:"shortcode"`
    }

    var req Request
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    expiry := time.Now().Add(time.Duration(req.Validity) * time.Minute).UTC().Format(time.RFC3339)
    shortLink := fmt.Sprintf("https://%s/%s", r.Host, req.Shortcode)

    resp := map[string]string{
        "shortLink": shortLink,
        "expiry":    expiry,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}