package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Python Run: received a request")
	cmd := exec.Command("/bin/sh", "script.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

	cmd.Stdout = &out
	fmt.Printf("Output: %q\n", out.String())

func main() {
	log.Print("Python Run: starting server...")
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
	}
	log.Printf("Python Run: listening on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}