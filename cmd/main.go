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

	basicRemote := remote.NewBasicRemote(tv)
	advancedRemote := remote.NewAdvancedRemote(radio)

	fmt.Println("Testing Basic Remote with TV")
	basicRemote.On()
	basicRemote.Off()

	fmt.Println("Testing Advanced Remote with Radio")
	advancedRemote.On()
	advancedRemote.Mute()
	advancedRemote.Off()
}
