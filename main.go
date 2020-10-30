package main

import (
	"log"
	"net/http"
	"strconv"

	judge "github.com/sadn1ck/code-runner/internal/judge"
)

func main() {
	http.HandleFunc("/submit", judge.SubmitHandler)
	http.HandleFunc("/status", judge.StatusHandler)
	PORT := 3000
	log.Printf("Server starting at port %d.....", PORT)

	SERVE := ":" + strconv.Itoa(PORT)
	log.Fatal(http.ListenAndServe(SERVE, nil))
}
