package main

import (
	"bridge-pattern-go/pkg/device"
	"bridge-pattern-go/pkg/remote"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting Bridge Pattern Demo")

	tv := device.NewTV()
	radio := device.NewRadio()

	basicRemote := remote.NewBasicRemote(tv)
	advancedRemote := remote.NewAdvancedRemote(radio)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Turn TV On")
		fmt.Println("2. Turn TV Off")
		fmt.Println("3. Turn Radio On")
		fmt.Println("4. Turn Radio Off")
		fmt.Println("5. Set TV Channel")
		fmt.Println("6. Set Radio Favorite Channel")
		fmt.Println("7. Go to Radio Favorite Channel")
		fmt.Println("8. Exit")

		if !scanner.Scan() {
			break
		}
		action := scanner.Text()

		switch action {
		case "1":
			basicRemote.On()
		case "2":
			basicRemote.Off()
		case "3":
			advancedRemote.On()
		case "4":
			advancedRemote.Off()
		case "5":
			fmt.Print("Enter channel number for TV: ")
			if scanner.Scan() {
				channel, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
				if err == nil {
					basicRemote.SetChannel(channel)
				} else {
					fmt.Println("Invalid channel number")
				}
			}
		case "6":
			fmt.Print("Enter favorite channel for Radio: ")
			if scanner.Scan() {
				channel, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
				if err == nil {
					advancedRemote.SetFavoriteChannel(channel)
				} else {
					fmt.Println("Invalid channel number")
				}
			}
		case "7":
			advancedRemote.GoToFavoriteChannel()
		case "8":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid action")
		}
	}
}
