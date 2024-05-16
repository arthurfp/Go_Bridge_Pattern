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
	SetVolume(level int, user string) error
	ToggleStandby()
	SetChannel(channel int, user string) error
	ResetProfiles()
}

type UserProfile struct {
	Name    string
	Volume  int
	Channel int
}

type TV struct {
	powerOn  bool
	volume   int
	standby  bool
	channel  int
	profiles map[string]UserProfile
}

func NewTV() Device {
	return &TV{profiles: make(map[string]UserProfile)}
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

func (tv *TV) SetVolume(level int, user string) error {
	if tv.standby {
		return errors.New("cannot change volume while in standby mode")
	}
	if level < 0 || level > 100 {
		return errors.New("volume must be between 0 and 100")
	}
	tv.volume = level
	tv.profiles[user] = UserProfile{Name: user, Volume: level, Channel: tv.channel}
	fmt.Printf("TV volume set to %d for user %s\n", level, user)
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

func (tv *TV) SetChannel(channel int, user string) error {
	if tv.standby {
		return errors.New("cannot change channel while in standby mode")
	}
	if channel < 1 || channel > 999 {
		return errors.New("channel must be between 1 and 999")
	}
	tv.channel = channel
	tv.profiles[user] = UserProfile{Name: user, Volume: tv.volume, Channel: channel}
	fmt.Printf("TV channel set to %d for user %s\n", channel, user)
	return nil
}

func (tv *TV) ResetProfiles() {
	tv.profiles = make(map[string]UserProfile)
	fmt.Println("All TV user profiles have been reset.")
}

type Radio struct {
	powerOn  bool
	volume   int
	standby  bool
	channel  int
	profiles map[string]UserProfile
}

func NewRadio() Device {
	return &Radio{profiles: make(map[string]UserProfile)}
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

func (radio *Radio) SetVolume(level int, user string) error {
	if radio.standby {
		return errors.New("cannot change volume while in standby mode")
	}
	if level < 0 || level > 100 {
		return errors.New("volume must be between 0 and 100")
	}
	radio.volume = level
	radio.profiles[user] = UserProfile{Name: user, Volume: level, Channel: radio.channel}
	fmt.Printf("Radio volume set to %d for user %s\n", level, user)
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

func (radio *Radio) SetChannel(channel int, user string) error {
	if radio.standby {
		return errors.New("cannot change channel while in standby mode")
	}
	if channel < 1 || channel > 999 {
		return errors.New("channel must be between 1 and 999")
	}
	radio.channel = channel
	radio.profiles[user] = UserProfile{Name: user, Volume: radio.volume, Channel: channel}
	fmt.Printf("Radio channel set to %d for user %s\n", channel, user)
	return nil
}

func (radio *Radio) ResetProfiles() {
	radio.profiles = make(map[string]UserProfile)
	fmt.Println("All Radio user profiles have been reset.")
}
