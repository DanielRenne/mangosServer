//Package pubsub supports the implementation of a publishing server.
package pubsub

import (
	"fmt"
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pub"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

type Server struct {
	url  string
	sock mangos.Socket
}

func (self *Server) Listen(url string) error {

	self.url = url

	var err error
	if self.sock, err = pub.NewSocket(); err != nil {
		return err
	}

	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())
	if err = self.sock.Listen(url); err != nil {
		return err
	}

	return nil
}

func (self *Server) Publish(payload []byte) error {
	err := self.sock.Send(payload)
	return err
}

func (self *Server) PublishTopic(topic string, message string) error {
	err := self.sock.Send([]byte(fmt.Sprintf("%s|%s", topic, message)))
	return err
}
