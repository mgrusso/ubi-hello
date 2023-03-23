package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got / request\n")
	io.WriteString(w, "Hello from inside RHEL UBI!\n")	
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}