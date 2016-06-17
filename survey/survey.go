//Package survey supports the implementation of a survey server.
package survey

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/surveyor"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"time"
)

type Server struct {
	url  string
	sock mangos.Socket
}

type SurveyResponse []byte

type SurveyResults struct {
	Responses []SurveyResponse
	Err       error
}

//Starts a Survey on the specified url for a specif.
func (self *Server) Listen(url string, seconds time.Duration) error {

	self.url = url
	var err error

	if self.sock, err = surveyor.NewSocket(); err != nil {
		return err
	}
	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())

	if err = self.sock.Listen(url); err != nil {
		return err
	}
	err = self.sock.SetOption(mangos.OptionSurveyTime, time.Second*seconds)
	if err != nil {
		return err
	}

	return nil

}

//Send the survey question to clients and set a channel of slice messages to process.
func (self *Server) Send(payload []byte, values chan SurveyResults) {

	var err error
	var results SurveyResults

	if err = self.sock.Send(payload); err != nil {
		results.Err = err
		values <- results
		return
	}

	var responses []SurveyResponse

	for {
		var msg SurveyResponse
		if msg, err = self.sock.Recv(); err != nil {
			break
		}
		responses = append(responses, msg)
	}

	results.Responses = responses

	values <- results
}

// //Sends the survey for a go routine.
// func sendSurvey(self *Server, payload []byte, values chan SurveyResults) {

// }
