package remote

import (
	"bridge-pattern-go/pkg/device"
	"testing"
)

func TestBasicRemote(t *testing.T) {
	tv := device.NewTV()
	remote := NewBasicRemote(tv)

	remote.On()
	if status := tv.PowerStatus(); status != "The TV is on" {
		t.Errorf("Expected TV to be on, got %s", status)
	}

	remote.VolumeUp()
	remote.ToggleStandby()
	remote.Off()
	if status := tv.PowerStatus(); status != "The TV is off" {
		t.Errorf("Expected TV to be off, got %s", status)
	}
}

func TestAdvancedRemote(t *testing.T) {
	radio := device.NewRadio()
	remote := NewAdvancedRemote(radio)

	remote.On()
	remote.Mute()
	if status := radio.PowerStatus(); status != "The Radio is on" {
		t.Errorf("Expected Radio to be on, got %s", status)
	}

	remote.VolumeDown() // Should remain muted (volume at 0)
	remote.Off()
	if status := radio.PowerStatus(); status != "The Radio is off" {
		t.Errorf("Expected Radio to be off, got %s", status)
	}
}
