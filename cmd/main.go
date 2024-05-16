package main

import (
	"bridge-pattern-go/pkg/device"
	"bridge-pattern-go/pkg/remote"
	"bridge-pattern-go/pkg/user"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tv := device.NewTV()
	radio := device.NewRadio()

	basicRemote := remote.NewBasicRemote(tv)
	advancedRemote := remote.NewAdvancedRemote(radio)

	userManager := user.NewUserManager()
	// Example users
	userManager.AddUser("John")
	userManager.AddUser("Jane")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nEnter command (on, off, volume, channel, profile, exit):")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		args := strings.Split(input, " ")
		command := args[0]

		switch command {
		case "on":
			basicRemote.On()
		case "off":
			basicRemote.Off()
		case "volume":
			if len(args) < 3 {
				fmt.Println("Usage: volume up/down user")
				continue
			}
			volCommand := args[1]
			targetUser := args[2]
			if profile, exists := userManager.GetUser(targetUser); exists {
				if volCommand == "up" {
					advancedRemote.SetVolume(profile.Volume+1, targetUser)
				} else if volCommand == "down" {
					advancedRemote.SetVolume(profile.Volume-1, targetUser)
				}
			}
		case "channel":
			if len(args) < 3 {
				fmt.Println("Usage: channel <number> <user>")
				continue
			}
			channel, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Invalid channel number:", args[1])
				continue
			}
			targetUser := args[2]
			advancedRemote.SetChannel(channel, targetUser)
		case "profile":
			if len(args) < 2 {
				fmt.Println("Usage: profile <user>")
				continue
			}
			targetUser := args[1]
			if profile, exists := userManager.GetUser(targetUser); exists {
				fmt.Printf("User: %s, Volume: %d, Channel: %d\n", profile.Name, profile.Volume, profile.Channel)
			} else {
				fmt.Println("User profile not found:", targetUser)
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}
