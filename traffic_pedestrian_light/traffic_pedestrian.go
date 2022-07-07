package main

import (
	"machine"
	"time"
)

var stopTraffic bool
var lights [5]machine.Pin
var buttonInput machine.Pin

var outputConfig machine.PinConfig
var inputConfig machine.PinConfig

func main() {
	stopTraffic = false

	outputConfig = machine.PinConfig{Mode: machine.PinOutput}
	inputConfig = machine.PinConfig{Mode: machine.PinInput}

	buttonInput = machine.D2
	buttonInput.Configure(inputConfig)

	lights[0] = machine.D13
	lights[1] = machine.D12
	lights[2] = machine.D11
	lights[3] = machine.D5
	lights[4] = machine.D4

	for i := range lights {
		lights[i].Configure(outputConfig)
	}

	lights[3].High()
	lights[4].Low()

	for {
		if buttonInput.Get() {
			stopTraffic = true
		}
		lightChange(lights[:])

		time.Sleep(time.Second * 3)
	}
}

func lightChange(pins []machine.Pin) {

	if stopTraffic {
		pins[0].High()
		pins[1].Low()
		pins[2].Low()

		pins[4].High()
		pins[3].Low()
		time.Sleep(time.Second * 20)

		pins[4].Low()
		pins[3].High()

		stopTraffic = false
	} else {
		pins[3].High()
		pins[4].Low()
	}

	pins[0].Low()
	pins[1].Low()
	pins[2].High()
	time.Sleep(time.Second * 40)

	pins[2].Low()
	pins[1].High()
	time.Sleep(time.Second * 3)

	pins[1].Low()
	pins[3].Low()
	pins[4].High()
	pins[0].High()
	time.Sleep(time.Second * 20)

	pins[1].High()
	time.Sleep(time.Second * 3)
}
