package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("templates")))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/answer", receiveAnswer)

	fmt.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func receiveAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Answer string `json:"answer"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid answer", http.StatusBadRequest)
		return
	}

	if data.Answer != "Yes" && data.Answer != "No" {
		http.Error(w, "Invalid answer", http.StatusBadRequest)
		return
	}

	fmt.Println("She answered:", data.Answer)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Answer received successfully!")
}
