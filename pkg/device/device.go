package device

import "fmt"

// Device interface
type Device interface {
	TurnOn()
	TurnOff()
}

// TV struct implementing Device
type TV struct{}

func NewTV() Device {
	return &TV{}
}

func (tv *TV) TurnOn() {
	fmt.Println("TV is on")
}

func (tv *TV) TurnOff() {
	fmt.Println("TV is off")
}

// Radio struct implementing Device
type Radio struct{}

func NewRadio() Device {
	return &Radio{}
}

func (radio *Radio) TurnOn() {
	fmt.Println("Radio is on")
}

func (radio *Radio) TurnOff() {
	fmt.Println("Radio is off")
}
