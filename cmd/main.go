package main

import (
	"bridge-pattern-go/pkg/device"
	"bridge-pattern-go/pkg/remote"
	"fmt"
)

func main() {
	fmt.Println("Starting Bridge Pattern Demo")

	// Initialize devices
	tv := device.NewTV()
	radio := device.NewRadio()

	// Initialize remotes for each device
	basicRemoteTV := remote.NewBasicRemote(tv)
	advancedRemoteRadio := remote.NewAdvancedRemote(radio)

	// Test TV controls
	fmt.Println("\nTesting TV with Basic Remote")
	basicRemoteTV.On()
	basicRemoteTV.VolumeUp()
	basicRemoteTV.VolumeDown()
	basicRemoteTV.Off()

	// Test Radio controls
	fmt.Println("\nTesting Radio with Advanced Remote")
	advancedRemoteRadio.On()
	advancedRemoteRadio.VolumeUp()
	advancedRemoteRadio.ToggleStandby()
	advancedRemoteRadio.Off()
}
