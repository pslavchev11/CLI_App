package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleProcessRequest(w http.ResponseWriter, r *http.Request) {

	// Ensure that the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body into a map
	decoder := json.NewDecoder(r.Body)
	var requestData map[string]interface{}
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Extract the process name from the request data
	processName, ok := requestData["process"].(string)
	if !ok {
		http.Error(w, "Process name not provided", http.StatusBadRequest)
		return
	}

	 // Execute the requested process
	responseData, err := executeProcess(processName)
	if err != nil {
		http.Error(w, "Error executing process", http.StatusInternalServerError)
		return
	}

	//response := map[string]string{"message": "the process executed successfully"}

	// Encode the response as JSON
	jsonResponse, err := json.Marshal(responseData)
	if err != nil{
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to incdicate the JSON content
	w.Header().Set("Content-Type", "application/json")

	// HTTP status code and response body
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

// Define a function to execute the requested process
func executeProcess(processName string) (map[string]interface{}, error) {
    // Implement the logic to execute the requested process based on the process name
    switch processName {
    case "process1":
        // Implement process1 logic here
        response := map[string]interface{}{
            "message": "Process 1 executed successfully",
        }
        return response, nil
    case "process2":
        // Implement process2 logic here
        response := map[string]interface{}{
            "message": "Process 2 executed successfully",
        }
        return response, nil
    default:
        // Handle unsupported process name
        return nil, fmt.Errorf("unsupported process: %s", processName)
    }
}