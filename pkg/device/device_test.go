package device

import (
	"testing"
)

func TestTVTurnOnOff(t *testing.T) {
	tv := NewTV()
	tv.TurnOn()
	if tv.PowerStatus() != "on" {
		t.Errorf("TV should be on")
	}
	tv.TurnOff()
	if tv.PowerStatus() != "off" {
		t.Errorf("TV should be off")
	}
}

func TestTVVolumeControl(t *testing.T) {
	tv := NewTV().(*TV)
	tv.TurnOn()
	initialVolume := tv.volume
	_ = tv.SetVolume(initialVolume+10, "default")
	if tv.volume != initialVolume+10 {
		t.Errorf("Expected volume %d, got %d", initialVolume+10, tv.volume)
	}
}

func TestTVStandbyMode(t *testing.T) {
	tv := NewTV().(*TV)
	tv.TurnOn()
	tv.ToggleStandby()
	if !tv.standby {
		t.Errorf("TV should be in standby mode")
	}
	_ = tv.SetVolume(0, "default") // This should fail due to standby mode
	if tv.volume != 0 {
		t.Errorf("Volume changes should not be allowed in standby mode")
	}
}
