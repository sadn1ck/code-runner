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
	status := handleResponseCode(sub.Code)
	var response string
	if status == "OK" {
		response = "{\"message\": \"AC\"}"
	} else {
		response = "{\"message\": \"" + status + "\"}"
	}
	log.Println(response)
	switch req.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Invalid request"}`))
	}
}

func handleResponseCode(Code string) string {
	// fetch these from question storage with appropriate calls (more endpoints!)
	in := []string{"2"}
	out := []string{"4"}
	res := run.Evaluate(Code, "cpp", "a.cpp", in, out, 1, 1, 500*1024*1024)
	// log.Println(res.Status)
	return res.Status
}
