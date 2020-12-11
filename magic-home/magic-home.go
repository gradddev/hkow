package magic_home

import (
	"net"
)

type Controller struct {
	Connection *net.TCPConn
}

func NewController(address string) (*Controller, error) {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	if error != nil {
		return nil, error
	}
	connection, error := net.DialTCP("tcp", nil, tcpAddr)
	if error != nil {
		return nil, error
	}
	controller := Controller{}
	controller.Connection = connection
	return &controller, nil
}

type Request []byte
type Response []byte

func (c *Controller) SendRequest(request Request, waitResponse bool) (Response, error) {
	checksum := CalculateChecksum(request)
	request = append(request, checksum)
	_, error := c.Connection.Write(request)
	if error != nil {
		return nil, error
	}
	if waitResponse {
		response := make([]byte, 32)
		_, error = c.Connection.Read(response)
		if error != nil {
			return nil, error
		}
		return response, nil
	}
	return nil, nil
}

func CalculateChecksum(data []byte) byte {
	checksum := uint8(0)
	for _, b := range data {
		checksum += uint8(b)
	}
	checksum &= 0xFF
	return checksum
}

func (c *Controller) IsSwitchedOn() (bool, error) {
	request := []byte{0x81, 0x8a, 0x8b}
	response, error := c.SendRequest(request, true)
	if error != nil {
		return false, error
	}
	return response[2] == 0x23, nil
}

func (c *Controller) SwitchOn() error {
	request := []byte{0x71, 0x23, 0x0f}
	_, error := c.SendRequest(request, false)
	if error != nil {
		return error
	}
	return nil
}

func (c *Controller) SwitchOff() error {
	request := []byte{0x71, 0x24, 0x0f}
	_, error := c.SendRequest(request, false)
	if error != nil {
		return error
	}
	return nil
}

type RGBW struct {
	Red   uint8
	Green uint8
	Blue  uint8
	White uint8
}

func (c *Controller) GetRGBW() (*RGBW, error) {
	request := []byte{0x81, 0x8a, 0x8b}
	response, error := c.SendRequest(request, true)
	if error != nil {
		return nil, error
	}
	return &RGBW{
		Red:   response[6],
		Green: response[7],
		Blue:  response[8],
		White: response[9],
	}, nil
}

func (c *Controller) SetRGBW(rgbw *RGBW) error {
	request := []byte{
		0x31,
		rgbw.Red,
		rgbw.Green,
		rgbw.Blue,
		rgbw.White,
		0,
		0x0f,
	}
	_, error := c.SendRequest(request, false)
	if error != nil {
		return error
	}
	return nil
}
