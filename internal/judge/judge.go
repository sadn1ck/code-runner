package judge

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/raydwaipayan/test-runner/runner/run"
)

// Submission is the incoming request format
type Submission struct {
	Code string
	ID   uint64
}

// SubmitHandler initializes the server
func SubmitHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parsedBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))
	var sub Submission
	err = json.Unmarshal(parsedBody, &sub)
	if err != nil {
		panic(err)
	}
	log.Println(sub.Code)
	status := handleResponseCode(sub.Code)
	log.Println(status)
	switch req.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST called"}`))
	default:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "!POST"}`))
	}
}

func handleResponseCode(Code string) string {
	// fetch these from question storage with appropriate calls (more endpoints!)
	in := []string{""}
	out := []string{"Hey"}
	log.Println("Reaches here")
	res := run.Evaluate(Code, "cpp", "a.cpp", in, out, 1, 1, 500*1024*1024)
	log.Println(res)
	return "AC"
}
