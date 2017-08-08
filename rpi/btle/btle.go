package btle

import (
	"github.com/currantlabs/ble/examples/lib/dev"
	"log"
	"github.com/currantlabs/ble"
	"context"
)

func SetupBTLE(nextPatternChannel chan<- string) {
	d, err := dev.NewDevice("")
	if err != nil {
		log.Fatalf("Can't create new device: %s", err)
	}

	ble.SetDefaultDevice(d)

	lightsService := ble.NewService(ble.MustParse("AA6E87FC-288B-40E8-942D-33409B435957"))
	lightsService.AddCharacteristic(NewNextPatternCharacteristic(nextPatternChannel))

	if err := ble.AddService(lightsService); err != nil {
		log.Fatalf("Can't add service: %s", err)
	}

	go ble.AdvertiseNameAndServices(context.Background(), "Party Lights", lightsService.UUID)
}