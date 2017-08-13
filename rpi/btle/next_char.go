package btle

import (
	"github.com/currantlabs/ble"
	"log"
	"github.com/Lobaro/slip"
	"bytes"
)

func NewNextPatternCharacteristic(nextPatternChannel chan<- string) *ble.Characteristic {
	c := ble.NewCharacteristic(ble.MustParse("720EEE67-CCE2-430D-BB40-4815F3B69295"))

	c.HandleWrite(ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
		reader := slip.NewReader(bytes.NewReader(req.Data()))

		packet, isPrefix, err := reader.ReadPacket()

		log.Printf("Received data: %s, %s, %s", packet, isPrefix, err)
	}))

	return c
}
