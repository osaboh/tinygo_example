package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter
var advState = false

func main() {

	ser := machine.UART0
	ser.Configure(machine.UARTConfig{
		BaudRate: 115200,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	})

	time.Sleep(3 * time.Second)

	left := machine.BUTTONA
	left.Configure(machine.PinConfig{Mode: machine.PinInput})

	right := machine.BUTTONB
	right.Configure(machine.PinConfig{Mode: machine.PinInput})

	ledrow := machine.LED_ROW_1
	ledrow.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledcol := machine.LED_COL_1
	ledcol.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledcol.Low()

	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "Go Bluetooth",
	}))

	for {
		if !advState && !left.Get() {
			must("start adv", adv.Start())
			fmt.Println("start advertising...")
			address, _ := adapter.Address()
			fmt.Println("Go Bluetooth /", address.MAC.String())

			ledrow.High()
			advState = true
		} else if advState && !right.Get() {
			fmt.Println("stop advertising...")
			must("stop adv", adv.Stop())

			ledrow.Low()
			advState = false
		} else {
			if advState {
				fmt.Println("Push right to stop BLE advertise")
			} else {
				fmt.Println("Push left to start BLE advertise")
			}
		}

		time.Sleep(time.Second)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
