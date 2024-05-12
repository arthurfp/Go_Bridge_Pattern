package device

import "fmt"

// Device interface defines common operations for all devices.
type Device interface {
	TurnOn()
	TurnOff()
	VolumeUp()
	VolumeDown()
	PowerStatus() string
	SetVolume(level int)
	ToggleStandby()
}

// TV struct represents a television with basic functionalities.
type TV struct {
	powerOn bool
	volume  int
	standby bool
}

// NewTV initializes a new TV instance.
func NewTV() Device {
	return &TV{}
}

// TurnOn powers on the TV unless it's in standby mode.
func (tv *TV) TurnOn() {
	if !tv.standby {
		tv.powerOn = true
		fmt.Println("TV is now on")
	}
}

// TurnOff powers off the TV.
func (tv *TV) TurnOff() {
	tv.powerOn = false
	fmt.Println("TV is now off")
}

// VolumeUp increases the volume by one unit.
func (tv *TV) VolumeUp() {
	tv.volume++
	fmt.Println("Increased TV volume to", tv.volume)
}

// VolumeDown decreases the volume by one unit.
func (tv *TV) VolumeDown() {
	if tv.volume > 0 {
		tv.volume--
		fmt.Println("Decreased TV volume to", tv.volume)
	}
}

// PowerStatus checks if the TV is on or off.
func (tv *TV) PowerStatus() string {
	if tv.powerOn {
		return "The TV is on"
	}
	return "The TV is off"
}

// SetVolume sets the volume to a specific level.
func (tv *TV) SetVolume(level int) {
	tv.volume = level
	fmt.Println("Set TV volume to", tv.volume)
}

// ToggleStandby toggles the TV's standby status.
func (tv *TV) ToggleStandby() {
	tv.standby = !tv.standby
	state := "off"
	if tv.standby {
		state = "on"
	}
	fmt.Println("TV standby mode is", state)
}

// Radio struct represents a radio with basic functionalities.
type Radio struct {
	powerOn bool
	volume  int
	standby bool
}

// NewRadio initializes a new Radio instance.
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
	radio.volume++
	fmt.Println("Increased Radio volume to", radio.volume)
}

func (radio *Radio) VolumeDown() {
	if radio.volume > 0 {
		radio.volume--
		fmt.Println("Decreased Radio volume to", radio.volume)
	}
}

func (radio *Radio) PowerStatus() string {
	if radio.powerOn {
		return "The Radio is on"
	}
	return "The Radio is off"
}

func (radio *Radio) SetVolume(level int) {
	radio.volume = level
	fmt.Println("Set Radio volume to", radio.volume)
}

func (radio *Radio) ToggleStandby() {
	radio.standby = !radio.standby
	state := "off"
	if radio.standby {
		state = "on"
	}
	fmt.Println("Radio standby mode is", state)
}
