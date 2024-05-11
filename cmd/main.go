package main

import (
	"bridge-pattern-go/pkg/device"
	"bridge-pattern-go/pkg/remote"
	"fmt"
)

func main() {
	fmt.Println("Bridge Pattern in Go")

	tv := device.NewTV()
	radio := device.NewRadio()

	// Initialize remotes with their respective devices
	basicRemoteTV := remote.NewBasicRemote(tv)
	advancedRemoteRadio := remote.NewAdvancedRemote(radio)

	// Demonstrate basic remote functions with TV
	fmt.Println("\nTesting Basic Remote with TV")
	basicRemoteTV.On()
	basicRemoteTV.VolumeUp()
	basicRemoteTV.VolumeDown()
	basicRemoteTV.Off()

	// Demonstrate advanced remote functions with Radio
	fmt.Println("\nTesting Advanced Remote with Radio")
	advancedRemoteRadio.On()
	advancedRemoteRadio.VolumeUp()
	advancedRemoteRadio.Mute() // Ensure this function is available in your AdvancedRemote type
	advancedRemoteRadio.Off()

	// Checking the power status using the advanced remote
	status := advancedRemoteRadio.CheckPowerStatus()
	fmt.Println("Radio power status:", status)
}
