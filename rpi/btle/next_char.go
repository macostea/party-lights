package btle

import (
	"github.com/currantlabs/ble"
	"github.com/Lobaro/slip"
	"bytes"
	"io"
	"log"
)

func NewNextPatternCharacteristic(nextPatternChannel chan<- []byte) *ble.Characteristic {
	c := ble.NewCharacteristic(ble.MustParse("720EEE67-CCE2-430D-BB40-4815F3B69295"))

	c.HandleWrite(ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
		reader := slip.NewReader(bytes.NewReader(req.Data()))

		packet, _, err := reader.ReadPacket()

		if err != io.EOF {
			log.Printf("Failed to receive packet, %s", err)
			return
		}

		nextPatternChannel <- packet
	}))

	return c
}
