package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/julienschmidt/httprouter"
)

func SendNotification() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var id map[string]string
		json.NewDecoder(r.Body).Decode(&id)
		fmt.Println(id["id"])
		valueType := reflect.TypeOf(id["id"])
		fmt.Println(valueType)
		notification := map[string]interface{}{
			"app_id":   "a97f9aed-da4e-4961-b351-9e3073068021",
			"contents": map[string]string{"en": fmt.Sprintf("Hello %s", id["id"])},
			// "included_segments": []string{"All"},
			"include_external_user_ids": []string{id["id"]},
		}

		// Convert the notification to JSON
		jsonData, err := json.Marshal(notification)
		if err != nil {
			fmt.Println("Failed to marshal notification:", err)
			return
		}

		// Create the HTTP request
		req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Failed to create request:", err)
			return
		}
		// Set the OneSignal REST API key as an Authorization header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Basic OTQ1YWU5YzgtM2RmMC00YzIzLTlhNjQtYjY3ZTk3NWI5N2Fi")

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Failed to send request:", err)
			return
		}
		defer resp.Body.Close()

		// Check the response status
		if resp.StatusCode == http.StatusOK {
			fmt.Println("Notification sent successfully")
		} else {
			fmt.Println("Failed to send notification. Status code:", resp.StatusCode)
		}
		json.NewEncoder(w).Encode(id)

	}
}
