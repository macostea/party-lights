package main

import (
	"github.com/macostea/party-lights/rpi/btle"
	"log"
	"github.com/macostea/party-lights/rpi/arduino"
	"github.com/hypebeast/go-osc/osc"
	"time"
)

func main() {
	arduinoConnection := arduino.NewConnection("/dev/cu.usbmodem1421", 9600)

	nextPatternChannel := make(chan string)
	btle.SetupBTLE(nextPatternChannel)


	for range nextPatternChannel {
		log.Print("Switch to next pattern")

		bundle := osc.NewBundle(time.Now().Add(time.Second * 2))
		msg := osc.NewMessage("/next")
		msg.Append(int32(1))

		bundle.Append(msg)

		if binaryMsg, err := bundle.MarshalBinary(); err != nil {
			log.Fatal("Failed to convert OSC message to binary")
		} else {
			arduinoConnection.WriteMessage(binaryMsg)
		}
	}
}
