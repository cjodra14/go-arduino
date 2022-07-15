package main

import (
	"machine"
	"time"
)

func main() {
	led1()
	led2()
	led3()
}

func led1() {
	led := machine.D9
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(500)
		led.Low()
		delay(500)
	}
}

func led2() {
	led := machine.D10
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(400)
		led.Low()
		delay(400)
	}
}

func led3() {
	led := machine.D11
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		delay(600)
		led.Low()
		delay(600)
	}
}

func delay(t int64) {
	time.Sleep(time.Duration(1000000 * t))
}
