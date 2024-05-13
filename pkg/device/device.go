package device

import "fmt"

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
	return &TV{}
}

func (tv *TV) TurnOn() {
	if !tv.standby {
		tv.powerOn = true
		fmt.Println("TV is now on")
	}
}

func (tv *TV) TurnOff() {
	tv.powerOn = false
	fmt.Println("TV is now off")
}

func (tv *TV) VolumeUp() {
	if !tv.standby {
		tv.volume++
		fmt.Println("TV volume increased to", tv.volume)
	}
}

func (tv *TV) VolumeDown() {
	if !tv.standby && tv.volume > 0 {
		tv.volume--
		fmt.Println("TV volume decreased to", tv.volume)
	}
}

func (tv *TV) PowerStatus() string {
	if tv.powerOn {
		return "on"
	}
	return "off"
}

func (tv *TV) SetVolume(level int) {
	if !tv.standby {
		tv.volume = level
		fmt.Println("TV volume set to", tv.volume)
	}
}

func (tv *TV) ToggleStandby() {
	tv.standby = !tv.standby
	if tv.standby {
		fmt.Println("TV is in standby mode")
	} else {
		fmt.Println("TV is active")
	}
}

func (tv *TV) SetChannel(channel int) {
	if !tv.standby {
		tv.channel = channel
		fmt.Println("TV channel set to", tv.channel)
	}
}

type Radio struct {
	powerOn bool
	volume  int
	standby bool
	channel int
}

func NewRadio() Device {
	return &Radio{}
}

func (radio *Radio) TurnOn() {
	if !radio.standby {
		radio.powerOn = true
		fmt.Println("Radio is now on")
	}
}

func (radio *Radio) TurnOff() {
	radio.powerOn = false
	fmt.Println("Radio is now off")
}

func (radio *Radio) VolumeUp() {
	if !radio.standby {
		radio.volume++
		fmt.Println("Radio volume increased to", radio.volume)
	}
}

func (radio *Radio) VolumeDown() {
	if !radio.standby && radio.volume > 0 {
		radio.volume--
		fmt.Println("Radio volume decreased to", radio.volume)
	}
}

func (radio *Radio) PowerStatus() string {
	if radio.powerOn {
		return "on"
	}
	return "off"
}

func (radio *Radio) SetVolume(level int) {
	if !radio.standby {
		radio.volume = level
		fmt.Println("Radio volume set to", radio.volume)
	}
}

func (radio *Radio) ToggleStandby() {
	radio.standby = !radio.standby
	if radio.standby {
		fmt.Println("Radio is in standby mode")
	} else {
		fmt.Println("Radio is active")
	}
}

func (radio *Radio) SetChannel(channel int) {
	if !radio.standby {
		radio.channel = channel
		fmt.Println("Radio channel set to", radio.channel)
	}
}
