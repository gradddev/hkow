package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/AlexeySemigradsky/hkmh"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	bridge := accessory.NewBridge(accessory.Info{Name: "Bridge", ID: 1})

	deskLight, err := hkmh.NewAccessory(
		accessory.Info{Name: "Desk Light", ID: 2},
		"192.168.1.211:5577",
		3*time.Second,
	)
	if err != nil {
		log.Panicln(err)
	}

	bedLight, err := hkmh.NewAccessory(
		accessory.Info{Name: "Bed Light", ID: 3},
		"192.168.1.216:5577",
		3*time.Second,
	)
	if err != nil {
		log.Panicln(err)
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicln(err)
	}

	storagePath := path.Join(userHomeDir, ".hkow")
	log.Println(storagePath)
	config := hc.Config{
		StoragePath: storagePath,
		Pin:         "00207700",
	}
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
