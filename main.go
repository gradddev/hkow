package main

import (
	"log"

	"github.com/AlexeySemigradsky/hkmh"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	bridge := accessory.NewBridge(accessory.Info{Name: "Bridge", ID: 1})

	deskLight, err := hkmh.NewAccessory(
		accessory.Info{Name: "Desk Light", ID: 2},
		"192.168.1.211:5577",
	)

	if err != nil {
		log.Panicln(err)
	}

	bedLight, err := hkmh.NewAccessory(
		accessory.Info{Name: "Bed Light", ID: 3},
		"192.168.1.216:5577",
	)

	if err != nil {
		log.Panicln(err)
	}

	config := hc.Config{Pin: "00207700"}
	t, err := hc.NewIPTransport(
		config,
		bridge.Accessory,
		deskLight.Accessory,
		bedLight.Accessory,
	)
	if err != nil {
		log.Panicln(err)
	}
	hc.OnTermination(func() {
		<-t.Stop()
	})
	t.Start()
}
