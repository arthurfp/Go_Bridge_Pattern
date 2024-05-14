package device

import (
	"log"
)

type Device interface {
	TurnOn()
	TurnOff()
	VolumeUp()
	VolumeDown()
	PowerStatus() string
	SetVolume(level int)
	ToggleStandby()
	SetChannel(channel int)
}

type TV struct {
	powerOn bool
	volume  int
	standby bool
	channel int
}

func NewTV() Device {
	log.Println("Initializing new TV")
	return &TV{}
}

func (tv *TV) TurnOn() {
	if !tv.standby {
		tv.powerOn = true
		log.Println("TV turned on")
	}
}

func (tv *TV) TurnOff() {
	tv.powerOn = false
	log.Println("TV turned off")
}

func (tv *TV) VolumeUp() {
	if !tv.standby {
		tv.volume++
		log.Printf("TV volume increased to %d\n", tv.volume)
	}
}

func (tv *TV) VolumeDown() {
	if !tv.standby && tv.volume > 0 {
		tv.volume--
		log.Printf("TV volume decreased to %d\n", tv.volume)
	}
}

func (tv *TV) PowerStatus() string {
	status := "off"
	if tv.powerOn {
		status = "on"
	}
	log.Printf("TV power status: %s\n", status)
	return status
}

func (tv *TV) SetVolume(level int) {
	if !tv.standby {
		tv.volume = level
		log.Printf("TV volume set to %d\n", level)
	}
}

func (tv *TV) ToggleStandby() {
	tv.standby = !tv.standby
	state := "active"
	if tv.standby {
		state = "standby"
	}
	log.Printf("TV switched to %s mode\n", state)
}

func (tv *TV) SetChannel(channel int) {
	if !tv.standby {
		tv.channel = channel
		log.Printf("TV channel set to %d\n", channel)
	}
}

type Radio struct {
	powerOn bool
	volume  int
	standby bool
	channel int
}

func NewRadio() Device {
	log.Println("Initializing new Radio")
	return &Radio{}
}

func (radio *Radio) TurnOn() {
	if !radio.standby {
		radio.powerOn = true
		log.Println("Radio turned on")
	}
}

func (radio *Radio) TurnOff() {
	radio.powerOn = false
	log.Println("Radio turned off")
}

func (radio *Radio) VolumeUp() {
	if !radio.standby {
		radio.volume++
		log.Printf("Radio volume increased to %d\n", radio.volume)
	}
}

func (radio *Radio) VolumeDown() {
	if !radio.standby && radio.volume > 0 {
		radio.volume--
		log.Printf("Radio volume decreased to %d\n", radio.volume)
	}
}

func (radio *Radio) PowerStatus() string {
	status := "off"
	if radio.powerOn {
		status = "on"
	}
	log.Printf("Radio power status: %s\n", status)
	return status
}

func (radio *Radio) SetVolume(level int) {
	if !radio.standby {
		radio.volume = level
		log.Printf("Radio volume set to %d\n", level)
	}
}

func (radio *Radio) ToggleStandby() {
	radio.standby = !radio.standby
	state := "active"
	if radio.standby {
		state = "standby"
	}
	log.Printf("Radio switched to %s mode\n", state)
}

func (radio *Radio) SetChannel(channel int) {
	if !radio.standby {
		radio.channel = channel
		log.Printf("Radio channel set to %d\n", channel)
	}
}
