package main

import (
	"log"
	"net/http"

	judge "github.com/sadn1ck/code-runner/internal/judge"
)

func main() {
	http.HandleFunc("/", judge.SubmitHandler)
	log.Println("Server starting at port 3000.....")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
