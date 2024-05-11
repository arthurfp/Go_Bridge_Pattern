package device

import "fmt"

// Device interface includes the methods that any device should implement
type Device interface {
	TurnOn()
	TurnOff()
	VolumeUp()
	VolumeDown()
	PowerStatus() string
	SetVolume(level int)
}

// TV struct implementing the Device interface
type TV struct {
	powerOn bool
	volume  int
}

// NewTV creates a new TV object
func NewTV() Device {
	return &TV{}
}

// TurnOn turns the TV on
func (tv *TV) TurnOn() {
	tv.powerOn = true
	fmt.Println("TV is turned on")
}

// TurnOff turns the TV off
func (tv *TV) TurnOff() {
	tv.powerOn = false
	fmt.Println("TV is turned off")
}

// VolumeUp increases the volume of the TV
func (tv *TV) VolumeUp() {
	tv.volume++
	fmt.Println("TV volume increased to", tv.volume)
}

// VolumeDown decreases the volume of the TV
func (tv *TV) VolumeDown() {
	if tv.volume > 0 {
		tv.volume--
		fmt.Println("TV volume decreased to", tv.volume)
	}
}

// PowerStatus returns the current power status of the TV
func (tv *TV) PowerStatus() string {
	if tv.powerOn {
		return "on"
	}
	return "off"
}

// SetVolume sets the volume to a specific level
func (tv *TV) SetVolume(level int) {
	tv.volume = level
	fmt.Println("TV volume set to", tv.volume)
}
