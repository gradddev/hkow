package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AlexeySemigradsky/hkmh"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: hkow /path/to/storage/")
		return
	}

	storagePath := os.Args[1]
	stat, err := os.Stat(storagePath)
	if err != nil || !stat.Mode().IsDir() {
		fmt.Println("usage: hkow /path/to/storage/")
		return
	}

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
