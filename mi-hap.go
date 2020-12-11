package main

import (
	"log"

	led_strip_light "github.com/AlexeySemigradsky/mi-hap/led-strip-light"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	info := accessory.Info{Name: "LED Strip Light 2077"}
	a, error := led_strip_light.NewAccessory(info, "192.168.1.211:5577")
	if error != nil {
		log.Panic(error)
	}

	config := hc.Config{Pin: "00002077"}
	t, error := hc.NewIPTransport(config, a.Accessory)
	if error != nil {
		log.Panic(error)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}
