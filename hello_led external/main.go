package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		for i := 0; i < 20; i++ {
			led.High()
			time.Sleep(time.Millisecond * 100)
			led.Low()
			time.Sleep(time.Millisecond * 100)
		}

		led.High()
		time.Sleep(time.Millisecond * 1000)
		led.Low()
		time.Sleep(time.Millisecond * 1000)
	}
}
