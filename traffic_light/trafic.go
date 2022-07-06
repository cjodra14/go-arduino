package main

import (
	"machine"
	"time"
)

func main() {
	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	var pins [3]machine.Pin
	pins[0] = machine.D11
	pins[1] = machine.D12
	pins[2] = machine.D13
	for i := range pins {
		pins[i].Configure(outputConfig)
	}
	lightChange(pins[:])

}

func lightChange(pins []machine.Pin) {
	for {
		pins[0].High()
		time.Sleep(time.Second)
		pins[1].High()
		time.Sleep(time.Second)
		pins[0].Low()
		pins[1].Low()
		pins[2].High()
		time.Sleep(time.Second)
		pins[2].Low()
		pins[1].High()
		time.Sleep(time.Second)
		pins[1].Low()
	}

}
