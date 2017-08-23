package btle

import (
	"github.com/currantlabs/ble"
)

func NewNextPatternCharacteristic(nextPatternChannel chan<- []byte) *ble.Characteristic {
	c := ble.NewCharacteristic(ble.MustParse("720EEE67-CCE2-430D-BB40-4815F3B69295"))

	c.HandleWrite(ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
		packet := req.Data()
		nextPatternChannel <- packet
	}))

	return c
}
