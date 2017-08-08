package btle

import (
	"github.com/currantlabs/ble"
	"log"
)

func NewNextPatternCharacteristic(nextPatternChannel chan<- string) *ble.Characteristic {
	c := ble.NewCharacteristic(ble.MustParse("720EEE67-CCE2-430D-BB40-4815F3B69295"))

	c.HandleWrite(ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
		command := string(req.Data())
		log.Printf("Received data: %s", command)

		if command == "next" {
			nextPatternChannel <- "next"
		}
	}))

	return c
}
