package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/process-single", ProcessSingle)
	http.HandleFunc("/process-concurrent", ProcessConcurrent)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
