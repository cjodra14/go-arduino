package main

import (
	"machine"
	"strings"
	"time"
)

func main() {
	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led.Set(false)
	pinSleep(3)

	LedMorse("Hello world", led)

	// Message finished flag
	for i := 0; i < 5; i++ {
		led.Set(true)
		time.Sleep(time.Millisecond * 100)
		led.Set(false)
		time.Sleep(time.Millisecond * 100)
	}
}

func Dot(pin machine.Pin) {
	pin.Set(true)
	pinSleep(1)
	pin.Set(false)
	pinSleep(1)
}

func dashFinal(pin machine.Pin) {
	pin.Set(true)
	pinSleep(3)
}

func dotFinal(pin machine.Pin) {
	pin.Set(true)
	pinSleep(1)
}

func Dash(pin machine.Pin) {
	pin.Set(true)
	pinSleep(3)
	pin.Set(false)
	pinSleep(1)
}

func letterSpace(pin machine.Pin) {
	pin.Set(false)
	pinSleep(3)
}

func pinSleep(sleeptime int) {
	time.Sleep(time.Second * time.Duration(sleeptime))
}

func wordSpace(pin machine.Pin) {
	pin.Set(false)
	pinSleep(7)
}

func LedMorse(text string, pin machine.Pin) {
	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890.,? ")
	word := strings.ToUpper(text)
	for _, char := range word {
		switch char {
		case alphabet[0]:
			Dot(pin)
			dashFinal(pin)
		case alphabet[1]:
			Dash(pin)
			Dot(pin)
			Dot(pin)
			dotFinal(pin)
		case alphabet[2]:
			Dash(pin)
			Dot(pin)
			Dash(pin)
			dotFinal(pin)
		case alphabet[3]:
			Dash(pin)
			Dot(pin)
			dotFinal(pin)
		case alphabet[4]:
			dotFinal(pin)
		case alphabet[5]:
			Dot(pin)
			Dot(pin)
			Dash(pin)
			dotFinal(pin)
		case alphabet[6]:
			Dash(pin)
			Dash(pin)
			dotFinal(pin)
		case alphabet[7]:
			Dot(pin)
			Dot(pin)
			Dot(pin)
			dotFinal(pin)
		case alphabet[8]:
			Dot(pin)
			dotFinal(pin)
		case alphabet[9]:
			Dot(pin)
			Dash(pin)
			Dash(pin)
			dashFinal(pin)
		case alphabet[10]:
			Dash(pin)
			Dot(pin)
			dashFinal(pin)
		case alphabet[11]:
			Dot(pin)
			Dash(pin)
			Dot(pin)
			dotFinal(pin)
		case alphabet[12]:
			Dash(pin)
			dashFinal(pin)
		case alphabet[13]:
			Dash(pin)
			dotFinal(pin)
		case alphabet[14]:
			Dash(pin)
			Dash(pin)
			dashFinal(pin)
		case alphabet[15]:
			Dot(pin)
			Dash(pin)
			Dash(pin)
			dotFinal(pin)
		case alphabet[16]:
			Dash(pin)
			Dash(pin)
			Dot(pin)
			dashFinal(pin)
		case alphabet[17]:
			Dot(pin)
			Dash(pin)
			dotFinal(pin)
		case alphabet[18]:
			Dot(pin)
			Dot(pin)
			dotFinal(pin)
		case alphabet[19]:
			dashFinal(pin)
		case alphabet[20]:
			Dot(pin)
			Dot(pin)
			dashFinal(pin)
		case alphabet[21]:
			Dot(pin)
			Dot(pin)
			Dot(pin)
			dashFinal(pin)
		case alphabet[22]:
			Dot(pin)
			Dash(pin)
			dashFinal(pin)
		case alphabet[23]:
			Dash(pin)
			Dot(pin)
			Dot(pin)
			dashFinal(pin)
		case alphabet[24]:
			Dash(pin)
			Dot(pin)
			Dash(pin)
			dashFinal(pin)
		case alphabet[25]:
			Dash(pin)
			Dash(pin)
			Dot(pin)
			dotFinal(pin)
		// TODO numbers dot, comma and signs
		case alphabet[39]:
			wordSpace(pin)
		default:
			for i := 0; i < 5; i++ {
				pin.Set(true)
				time.Sleep(time.Millisecond * 100)
				pin.Set(false)
				time.Sleep(time.Millisecond * 100)
			}
		}
		if alphabet[39] != char {
			letterSpace(pin)
		}
	}
}
