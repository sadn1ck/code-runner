package judge

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/raydwaipayan/test-runner/runner/run"
)

// Submission is the incoming request format
type Submission struct {
	UserID     string
	Code       string
	QuestionID uint64
}

type checkStatusRequest struct {
	SubmissionID string
}

type submissionStats struct {
	SubmissionID string
	Status       string
}

var questionStatus []submissionStats

// SubmitHandler initializes the server
func SubmitHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parsedBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var currentSubmission Submission
	err = json.Unmarshal(parsedBody, &currentSubmission)
	if err != nil {
		panic(err)
	}
	var response string = "{\"message\": \"submitted\", \"submissionID\": \"" + getSubmissionID(currentSubmission) + "\"}"

	switch req.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Invalid request"}`))
	}
}

func getSubmissionID(currentSubmission Submission) string {
	// fetch these from question storage with appropriate calls (more endpoints!)
	in := []string{"2", "4"}
	out := []string{"4", "16"}
	genuuid, _ := uuid.NewUUID()
	generatedSubmissionID := strings.Replace(genuuid.String(), "-", "", -1)

	questionStatus = append(questionStatus, submissionStats{generatedSubmissionID, "Running"})
	res := run.Evaluate(currentSubmission.Code, "cpp", "a.cpp", in, out, 2, 1, 500*1024*1024)

	var statusOfCurrent string

	results := res.Result
	for index := range results {
		if results[index] != "ACCEPTED" {
			statusOfCurrent = results[index]
			break
		}
		statusOfCurrent = "ACCEPTED"
	}
	for index := range questionStatus {
		if questionStatus[index].SubmissionID == generatedSubmissionID {
			questionStatus[index].Status = statusOfCurrent
		}
	}
	log.Println(res.Result)
	return generatedSubmissionID
}

// StatusHandler handles incoming POST request to produce a response code
func StatusHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parsedBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var statusToCheck checkStatusRequest
	err = json.Unmarshal(parsedBody, &statusToCheck)
	if err != nil {
		panic(err)
	}

	var status string
	for index := range questionStatus {
		if questionStatus[index].SubmissionID == statusToCheck.SubmissionID {
			status = questionStatus[index].Status
		}
	}
	response := "{\"status\": \"" + status + "\"}"
	switch req.Method {
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Invalid request"}`))
	}

}

// @TODO: This should be a GET request in the main backend
