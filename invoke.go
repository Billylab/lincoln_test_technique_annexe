package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Python Run: received a request")
	cmd := exec.Command("/bin/sh", "script.sh")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	// https://www.alexedwards.net/blog/golang-response-snippets
	// https://stackoverflow.com/questions/29237411/how-to-convert-type-bytes-buffer-to-use-as-byte-in-argument-to-w-write
	// https://zetcode.com/golang/exec-command/
	w.Write(out.Bytes()) 
	}

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