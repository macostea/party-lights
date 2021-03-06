package arduino

import (
	"github.com/tarm/serial"
	"github.com/mgutz/logxi/v1"
	"github.com/Lobaro/slip"
)

type S struct {
	port *serial.Port
}

func NewConnection(name string, baud int) *S {
	s := &S{}

	c := &serial.Config{}
	c.Name = name
	c.Baud = baud

	p, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal("Error opening serial connection", err)
	}

	s.port = p

	return s
}

func (s *S) WriteMessage(message []byte) {
	writer := slip.NewWriter(s.port)

	err := writer.WritePacket(message)

	if err != nil {
		log.Fatal("Failed to write message to serial connection", err)
	}
}
