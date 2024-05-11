package remote

import (
	"bridge-pattern-go/pkg/device"
	"fmt"
)

// Remote interface
type Remote interface {
	On()
	Off()
	VolumeUp()
	VolumeDown()
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

// AdvancedRemote struct with additional functionality
type AdvancedRemote struct {
	device device.Device
}

func NewAdvancedRemote(device device.Device) Remote {
	return &AdvancedRemote{device: device}
}

func (r *AdvancedRemote) On() {
	r.device.TurnOn()
	fmt.Println("Advanced remote: device turned on")
}

func (r *AdvancedRemote) Off() {
	r.device.TurnOff()
	fmt.Println("Advanced remote: device turned off")
}

func (r *AdvancedRemote) Mute() {
	fmt.Println("Advanced remote: device muted")
	// Assuming mute functionality is implemented in the device
}
