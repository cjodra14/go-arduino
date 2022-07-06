package main

import(
	"machine"
)

func main(){
	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInput})
	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		if button.Get(){
			led.Set(true)
			continue
		}
		led.Set(false)
	}
}