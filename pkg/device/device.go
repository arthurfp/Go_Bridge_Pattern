package device

import (
	"errors"
	"fmt"
)

type Device interface {
	TurnOn()
	TurnOff()
	VolumeUp()
	VolumeDown()
	PowerStatus() string
	SetVolume(level int) error
	ToggleStandby()
	SetChannel(channel int) error
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
	if !tv.standby && tv.volume < 100 {
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

func (tv *TV) SetVolume(level int) error {
	if tv.standby {
		return errors.New("cannot change volume while in standby mode")
	}
	if level < 0 || level > 100 {
		return errors.New("volume must be between 0 and 100")
	}
	tv.volume = level
	fmt.Println("TV volume set to", tv.volume)
	return nil
}

func (tv *TV) ToggleStandby() {
	tv.standby = !tv.standby

	if tv.standby {
		fmt.Println("TV switched to standby mode")
	} else {
		fmt.Println("TV switched to active mode")
	}
}

func (tv *TV) SetChannel(channel int) error {
	if tv.standby {
		return errors.New("cannot change channel while in standby mode")
	}
	if channel < 1 || channel > 999 {
		return errors.New("channel must be between 1 and 999")
	}
	tv.channel = channel
	fmt.Println("TV channel set to", tv.channel)
	return nil
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
	if !radio.standby && radio.volume < 100 {
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

func (radio *Radio) SetVolume(level int) error {
	if radio.standby {
		return errors.New("cannot change volume while in standby mode")
	}
	if level < 0 || level > 100 {
		return errors.New("volume must be between 0 and 100")
	}
	radio.volume = level
	fmt.Println("Radio volume set to", radio.volume)
	return nil
}

func (radio *Radio) ToggleStandby() {
	radio.standby = !radio.standby

	if radio.standby {
		fmt.Println("Radio switched to standby mode")
	} else {
		fmt.Println("Radio switched to active mode")
	}
}

func (radio *Radio) SetChannel(channel int) error {
	if radio.standby {
		return errors.New("cannot change channel while in standby mode")
	}
	if channel < 1 || channel > 999 {
		return errors.New("channel must be between 1 and 999")
	}
	radio.channel = channel
	fmt.Println("Radio channel set to", radio.channel)
	return nil
}
