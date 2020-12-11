package led_strip_light

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type WService struct {
	*service.Service

	On         *characteristic.On
	Brightness *characteristic.Brightness
}

func NewWService() *WService {
	s := WService{}
	s.Service = service.New(service.TypeLightbulb)

	s.On = characteristic.NewOn()
	s.AddCharacteristic(s.On.Characteristic)

	s.Brightness = characteristic.NewBrightness()
	s.AddCharacteristic(s.Brightness.Characteristic)

	return &s
}
