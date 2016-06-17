package survey

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/respondent"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	// "strings"
	"testing"
	"time"
)

const url = "tcp://127.0.0.1:600"

//Creates a new Survey Server and Tests a single Survey
func TestSingleSurvey(t *testing.T) {
	var s Server
	err := s.Listen(url, 2)
	if err != nil {
		t.Errorf("Error at survey.TestSingleSurvey:  %v", err.Error)
	}

	var sock mangos.Socket

	if sock, err = respondent.NewSocket(); err != nil {
		t.Errorf("Error creating new Socket at survey.TestSingleSurvey:  %v", err.Error)
	}

	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())

	if err = sock.Dial(url); err != nil {
		t.Errorf("Error Dialing at survey.TestSingleSurvey:  %v", err.Error)
		return
	}

	messages := make(chan string)

	go respondToSurvey(sock, t, messages, "TestSingle", "HelloWorld")

	time.Sleep(2 * time.Second)
	results := make(chan SurveyResults)
	go s.Send([]byte("TestSingle"), results)

	msg := <-messages
	t.Log(msg)

	sResults := <-results

	if sResults.Err != nil {
		t.Errorf("Error at survey.TestSingleSurvey:  %v", sResults.Err.Error)
		return
	}

	// go subscribeToAll(sock, t, messages)

	// s.Publish([]byte("TestSubscribeAll"))

}

//Responds to the Survey
func respondToSurvey(sock mangos.Socket, t *testing.T, messages chan string, surveyQuestion string, surveyResponse string) {
	var err error
	var msg []byte

	if msg, err = sock.Recv(); err != nil {
		t.Errorf("Error Receiving at survey.respondToSurvey:  %v", err.Error)
		messages <- "Test Failed"
		return
	}

	if string(msg) != surveyQuestion {
		t.Errorf("Failed to respond to survey question.")
		messages <- "Test Failed"
		return
	}

	if err = sock.Send([]byte(surveyResponse)); err != nil {
		t.Errorf("Error Sending Survey Response at survey.respondToSurvey:  %v", err.Error)
		messages <- "Test Failed"
		return
	}

	messages <- "Test Completed"
}
