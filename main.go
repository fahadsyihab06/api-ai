package main

import (
    "encoding/json"
    "net/http"
    "strings"
    "time"

    "github.com/jdkato/prose/v2"
)

type Request struct {
    Message string `json:"message"`
}

type Response struct {
    Reply string `json:"reply"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var req Request
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    userMessage := strings.ToLower(req.Message)
    var reply string

    // Contoh sederhana untuk jenis pertanyaan yang berbeda
    if strings.Contains(userMessage, "hello") || strings.Contains(userMessage, "hi") {
        reply = "Hello! How can I assist you today?"
    } else if strings.Contains(userMessage, "weather") {
        reply = "I can't check the weather right now, but I hope it's sunny where you are!"
    } else if strings.Contains(userMessage, "time") {
        currentTime := time.Now().Format("3:04 PM")
        reply = "The current time is " + currentTime
    } else if strings.Contains(userMessage, "date") {
        currentDate := time.Now().Format("Monday, January 2, 2006")
        reply = "Today's date is " + currentDate
    } else {
        // Jika tidak ada match, gunakan NLP untuk mencari entitas
        doc, _ := prose.NewDocument(req.Message)
        for _, ent := range doc.Entities() {
            reply += "I found an entity: " + ent.Text + " of type " + ent.Label + "\n"
        }
        if reply == "" {
            reply = "I'm sorry, I didn't quite understand that."
        }
    }

    res := Response{Reply: strings.TrimSpace(reply)}
    json.NewEncoder(w).Encode(res)
}

func main() {
    http.HandleFunc("/chat", handler)
    http.ListenAndServe(":8080", nil)
}
