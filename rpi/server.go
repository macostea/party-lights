package main

import (
	"github.com/macostea/party-lights/rpi/btle"
	"log"
	"github.com/macostea/party-lights/rpi/arduino"
)

func main() {
	arduinoConnection := arduino.NewConnection("/dev/cu.usbmodem1421", 9600)

	nextPatternChannel := make(chan string)
	btle.SetupBTLE(nextPatternChannel)

	for range nextPatternChannel {
		log.Print("Switch to next pattern")

		arduinoConnection.WriteMessage("next")
	}
}
