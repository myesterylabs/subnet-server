package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	// "strings"
)

// Arguments to hold additional arguments for the command
type Arguments struct {
	Args []string `json:"args"`
}

var authToken = new(string)

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != *authToken {
			http.Error(w, "Invalid auth token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func main() {
	authToken = flag.String("authToken", "098790879089789", "Auth token for the server")
	flag.Parse()
	// Define a command-line flag for the port, defaulting to 8818
	var port = flag.Int("port", 8818, "Port on which the server will listen")
	flag.Parse()

	// Setup the HTTP server routes
	http.HandleFunc("/", middleware(handler))

	// Create the address string using the provided port
	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Server starting on port %d...\n", *port)

	// Start listening and serving HTTP requests
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests for this handler
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// Decode JSON body for additional arguments
	var args Arguments
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		http.Error(w, "Error parsing JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the command with the arguments
	var baseCommand = "ipc-cli"
	command := exec.Command(baseCommand, args.Args...)
	output, err := command.CombinedOutput()
	if err != nil {
		var response = struct {
			Output string `json:"output"`
		}{
			Output: string(([]byte(err.Error()))),
		}

		json.NewEncoder(w).Encode(response)
	} else {
		// w.Write(output)
		// write as json
		w.Header().Set("Content-Type", "application/json")
		var response = struct {
			Output string `json:"output"`
		}{
			Output: string(output),
		}

		json.NewEncoder(w).Encode(response)
	}
}
