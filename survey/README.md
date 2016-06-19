# mangosServer Survey

Example Code to start a survey server, send a survey, and receive a survey response.

	package main

	import(
		"github.com/DanielRenne/mangosServer/survey"
		"log"
	)

	const url = "tcp://127.0.0.1:600"

	func main(){

		err := s.Listen(url, 500)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}

		err = s.Send([]byte("TestSurvey"), handleSurveyResponse)
		if err != nil {
			log.Printf("Error sending survey message:  %v", err.Error)
			return
		}
	}

	func handleSurveyResponse(msg []byte){
		//Process Survey Results.
		log.Printf(string(msg))
	}
	