package remote

import (
	"bridge-pattern-go/pkg/device"
	"fmt"
)

// Remote defines operations available on all remotes.
type Remote interface {
	On()
	Off()
	VolumeUp()
	VolumeDown()
	ToggleStandby()
}

// BasicRemote provides basic remote control functions.
type BasicRemote struct {
	device device.Device
}

// NewBasicRemote initializes a remote for any device.
func NewBasicRemote(device device.Device) Remote {
	return &BasicRemote{device: device}
}

// On powers the device on.
func (r *BasicRemote) On() {
	r.device.TurnOn()
	fmt.Println("Device turned on")
}

// Off powers the device off.
func (r *BasicRemote) Off() {
	r.device.TurnOff()
	fmt.Println("Device turned off")
}

// VolumeUp increases the volume by one level.
func (r *BasicRemote) VolumeUp() {
	r.device.VolumeUp()
	fmt.Println("Volume increased")
}

// VolumeDown decreases the volume by one level.
func (r *BasicRemote) VolumeDown() {
	r.device.VolumeDown()
	fmt.Println("Volume decreased")
}

// ToggleStandby toggles the standby mode of the device.
func (r *BasicRemote) ToggleStandby() {
	r.device.ToggleStandby()
	fmt.Println("Standby mode toggled")
}

// AdvancedRemote provides advanced features on top of basic remote functionalities.
type AdvancedRemote struct {
	*BasicRemote
}

// NewAdvancedRemote initializes an advanced remote with additional features.
func NewAdvancedRemote(device device.Device) *AdvancedRemote {
	return &AdvancedRemote{BasicRemote: NewBasicRemote(device).(*BasicRemote)}
}

// Mute completely silences the device.
func (r *AdvancedRemote) Mute() {
	r.device.SetVolume(0)
	fmt.Println("Device muted")
}
