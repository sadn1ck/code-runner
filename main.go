package main

import (
	"log"
	"net/http"
	"strconv"

	judge "github.com/sadn1ck/code-runner/internal/judge"
)

func main() {
	http.HandleFunc("/", judge.SubmitHandler)
	PORT := 3000
	log.Printf("Server starting at port %d.....", PORT)

	SERVE := ":" + strconv.Itoa(PORT)
	log.Fatal(http.ListenAndServe(SERVE, nil))
}
