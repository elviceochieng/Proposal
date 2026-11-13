package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("templates")))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/answer", receiveAnswer)

	fmt.Println("Server running at http://localhost:8080")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil)

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

	err = sendEmail(data.Answer)
	if err != nil {
		fmt.Println("Email notification failed:", err)
	} else {
		fmt.Println("Email notification sent successfully!")
	}

	fmt.Fprintln(w, "Your answer has been received! ❤️")
}

func sendEmail(answer string) error {
	apiKey := os.Getenv("RESEND_API_KEY")

	if apiKey == "" {
		return fmt.Errorf("RESEND_API_KEY is missing")
	}

	emailData := map[string]interface{}{
		"from":    "Proposal <onboarding@resend.dev>",
		"to":      []string{"delivered@resend.dev"},
		"subject": "Someone answered your question ❤️",
		"text":    "Her answer was: " + answer,
	}

	jsonData, err := json.Marshal(emailData)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		"POST",
		"https://api.resend.com/emails",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "Bearer "+apiKey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("Resend returned %s: %s", response.Status, string(body))
	}

	fmt.Println("Email request successful:", string(body))

	return nil
}
