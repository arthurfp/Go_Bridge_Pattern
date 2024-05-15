package remote

import (
	"bridge-pattern-go/pkg/device"
	"log"
)

type Remote interface {
	On()
	Off()
	VolumeUp()
	VolumeDown()
	ToggleStandby()
	SetChannel(channel int, user string)
	SetVolume(level int, user string)
}

type BasicRemote struct {
	device device.Device
}

func NewBasicRemote(device device.Device) Remote {
	log.Println("New Basic Remote created")
	return &BasicRemote{device: device}
}

func (r *BasicRemote) On() {
	log.Println("Turning device on with Basic Remote")
	r.device.TurnOn()
}

func (r *BasicRemote) Off() {
	log.Println("Turning device off with Basic Remote")
	r.device.TurnOff()
}

func (r *BasicRemote) VolumeUp() {
	log.Println("Increasing volume with Basic Remote")
	r.device.VolumeUp()
}

func (r *BasicRemote) VolumeDown() {
	log.Println("Decreasing volume with Basic Remote")
	r.device.VolumeDown()
}

func (r *BasicRemote) ToggleStandby() {
	log.Println("Toggling standby mode with Basic Remote")
	r.device.ToggleStandby()
}

func (r *BasicRemote) SetChannel(channel int, user string) {
	if err := r.device.SetChannel(channel, user); err != nil {
		log.Println("Error setting channel:", err)
	}
}

func (r *BasicRemote) SetVolume(level int, user string) {
	if err := r.device.SetVolume(level, user); err != nil {
		log.Println("Error setting volume:", err)
	}
}

type AdvancedRemote struct {
	*BasicRemote
	favoriteChannel int
}

func NewAdvancedRemote(device device.Device) *AdvancedRemote {
	log.Println("New Advanced Remote created")
	return &AdvancedRemote{BasicRemote: NewBasicRemote(device).(*BasicRemote)}
}

func (r *AdvancedRemote) Mute() {
	log.Println("Muting device with Advanced Remote")
	if err := r.device.SetVolume(0, "default"); err != nil {
		log.Println("Error muting volume:", err)
	}
}

func (r *AdvancedRemote) SetFavoriteChannel(channel int) {
	log.Printf("Setting favorite channel to %d with Advanced Remote\n", channel)
	r.favoriteChannel = channel
}

func (r *AdvancedRemote) GoToFavoriteChannel() {
	log.Printf("Going to favorite channel %d with Advanced Remote\n", r.favoriteChannel)
	if err := r.device.SetChannel(r.favoriteChannel, "default"); err != nil {
		log.Println("Error setting favorite channel:", err)
	}
}
