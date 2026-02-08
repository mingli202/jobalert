package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var port = 8000

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from go server!")
	})

	fmt.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
