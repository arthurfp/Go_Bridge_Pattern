package remote

import (
	"bridge-pattern-go/pkg/device"
	"fmt"
)

// Remote interface now includes volume control and power status check
type Remote interface {
	On()
	Off()
	VolumeUp()
	VolumeDown()
	CheckPowerStatus() string
}

// BasicRemote struct implementing the Remote interface
type BasicRemote struct {
	device device.Device
}

func NewBasicRemote(device device.Device) Remote {
	return &BasicRemote{device: device}
}

func (r *BasicRemote) On() {
	r.device.TurnOn()
	fmt.Println("Basic remote: device turned on")
}

func (r *BasicRemote) Off() {
	r.device.TurnOff()
	fmt.Println("Basic remote: device turned off")
}

func (r *BasicRemote) VolumeUp() {
	r.device.VolumeUp()
	fmt.Println("Basic remote: volume up")
}

func (r *BasicRemote) VolumeDown() {
	r.device.VolumeDown()
	fmt.Println("Basic remote: volume down")
}

func (r *BasicRemote) CheckPowerStatus() string {
	return r.device.PowerStatus()
}

// AdvancedRemote includes additional functionality
type AdvancedRemote struct {
	*BasicRemote
}

func NewAdvancedRemote(device device.Device) Remote {
	return &AdvancedRemote{BasicRemote: NewBasicRemote(device).(*BasicRemote)}
}

func (r *AdvancedRemote) Mute() {
	r.device.SetVolume(0)
	fmt.Println("Advanced remote: device muted")
}
