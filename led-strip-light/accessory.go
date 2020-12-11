package led_strip_light

import (
	magic_home "github.com/AlexeySemigradsky/mi-hap/magic-home"
	"github.com/brutella/hc/accessory"
	"github.com/lucasb-eyer/go-colorful"
)

type RGBWLEDStripLightAccessory struct {
	*accessory.Accessory
	RGB        *RGBService
	W          *WService
	Controller *magic_home.Controller
}

func NewAccessory(info accessory.Info, address string) (*RGBWLEDStripLightAccessory, error) {
	c, error := magic_home.NewController("192.168.1.211:5577")
	if error != nil {
		return nil, error
	}

	isSwitchedOn, error := c.IsSwitchedOn()
	if error != nil {
		return nil, error
	}

	rgbw, error := c.GetRGBW()
	if error != nil {
		return nil, error
	}

	color := colorful.Color{
		R: float64(rgbw.Red / 255),
		G: float64(rgbw.Green / 255),
		B: float64(rgbw.Blue / 255),
	}
	h, s, b := color.Hsv()
	w := float64(rgbw.White) / 255

	a := RGBWLEDStripLightAccessory{}
	a.Accessory = accessory.New(info, accessory.TypeLightbulb)

	a.RGB = NewRGBService()
	a.W = NewWService()

	a.Controller = c

	a.AddService(a.RGB.Service)
	a.AddService(a.W.Service)

	a.RGB.On.SetValue(isSwitchedOn)
	a.RGB.Hue.SetValue(h)
	a.RGB.Saturation.SetValue(s * 100)
	a.RGB.Brightness.SetValue(int(b * 100))

	a.W.On.SetValue(isSwitchedOn)
	a.W.Brightness.SetValue(int(w * 100))

	a.RGB.On.OnValueRemoteUpdate(func(_ bool) {
		a.UpdateState()
	})
	a.RGB.Brightness.OnValueRemoteUpdate(func(brightness int) {
		a.UpdateState()
	})
	a.RGB.Saturation.OnValueRemoteUpdate(func(saturation float64) {
		a.UpdateState()
	})
	a.RGB.Hue.OnValueRemoteUpdate(func(hue float64) {
		a.UpdateState()
	})

	a.W.On.OnValueRemoteUpdate(func(_ bool) {
		a.UpdateState()
	})
	a.W.Brightness.OnValueRemoteUpdate(func(brightness int) {
		a.UpdateState()
	})

	return &a, nil
}

func (a *RGBWLEDStripLightAccessory) UpdateState() error {
	if !a.RGB.On.GetValue() && !a.W.On.GetValue() {
		error := a.Controller.SwitchOff()
		if error != nil {
			return error
		}
		return nil
	}

	error := a.Controller.SwitchOn()
	if error != nil {
		return error
	}

	color := colorful.Hsv(
		a.RGB.Hue.GetValue(),
		a.RGB.Saturation.GetValue()/100,
		float64(a.RGB.Brightness.GetValue())/100,
	)
	r, g, b := color.RGB255()
	w := uint8(a.W.Brightness.GetValue())
	rgbw := magic_home.RGBW{
		Red:   r,
		Blue:  b,
		Green: g,
		White: w,
	}
	a.Controller.SetRGBW(&rgbw)
	return nil
}
