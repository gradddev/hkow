package main

import (
	"log"
	"net"
	"time"

	"github.com/gradddev/hkmh"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	bridge := accessory.NewBridge(accessory.Info{Name: "Mi Router 3G", ID: 1})

	deskLight := hkmh.NewAccessory(hkmh.Config{
		ID:      2,
		IP:      net.ParseIP("192.168.1.211"),
		Name:    "Desk Light",
		Timeout: 3 * time.Second,
	})

	bedLight := hkmh.NewAccessory(hkmh.Config{
		ID:      3,
		IP:      net.ParseIP("192.168.1.215"),
		Name:    "Bed Light",
		Timeout: 3 * time.Second,
	})

	storagePath := "/etc/.hkow"
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
	handleError(err)

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}

func handleError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
