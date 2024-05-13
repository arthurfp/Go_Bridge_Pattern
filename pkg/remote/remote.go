package remote

import (
	"bridge-pattern-go/pkg/device"
	"fmt"
)

type Remote interface {
	On()
	Off()
	VolumeUp()
	VolumeDown()
	ToggleStandby()
	SetChannel(channel int)
}

type BasicRemote struct {
	device device.Device
}

func NewBasicRemote(device device.Device) Remote {
	return &BasicRemote{device: device}
}

func (r *BasicRemote) On() {
	r.device.TurnOn()
	fmt.Println("Device turned on")
}

func (r *BasicRemote) Off() {
	r.device.TurnOff()
	fmt.Println("Device turned off")
}

func (r *BasicRemote) VolumeUp() {
	r.device.VolumeUp()
	fmt.Println("Volume increased")
}

func (r *BasicRemote) VolumeDown() {
	r.device.VolumeDown()
	fmt.Println("Volume decreased")
}

func (r *BasicRemote) ToggleStandby() {
	r.device.ToggleStandby()
	fmt.Println("Standby mode toggled")
}

func (r *BasicRemote) SetChannel(channel int) {
	r.device.SetChannel(channel)
	fmt.Println("Channel set to", channel)
}

type AdvancedRemote struct {
	*BasicRemote
	favoriteChannel int
}

func NewAdvancedRemote(device device.Device) *AdvancedRemote {
	return &AdvancedRemote{BasicRemote: NewBasicRemote(device).(*BasicRemote)}
}

func (r *AdvancedRemote) Mute() {
	r.device.SetVolume(0)
	fmt.Println("Device muted")
}

func (r *AdvancedRemote) SetFavoriteChannel(channel int) {
	r.favoriteChannel = channel
	fmt.Println("Favorite channel set to", channel)
}

func (r *AdvancedRemote) GoToFavoriteChannel() {
	r.device.SetChannel(r.favoriteChannel)
	fmt.Println("Switched to favorite channel", r.favoriteChannel)
}
