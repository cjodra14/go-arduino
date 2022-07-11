package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/hd44780i2c"
)

var data = []byte("Hello World\n LCD 16x02")

func main() {

	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(i2c, 0x27)
	lcd.Configure(hd44780i2c.Config{Width: 16, Height: 2})

	lcd.Print(data)
	time.Sleep(time.Second * 5)
	lcd.Print([]byte("We just print more text, to see what happens, when we overflow the 16x2 character limit"))

	time.Sleep(time.Second * 5)
	animation(lcd)
}

func animation(lcd hd44780i2c.Device) {
	text := []byte("Hello World \n Sent by \n Arduino Mega \n 2560 \n powered by TinyGo")
	lcd.ClearDisplay()

	for {
		for i := range text {
			lcd.Print([]byte(string(text[i])))
			time.Sleep(time.Millisecond * 150)
		}

		time.Sleep(time.Second * 2)
		lcd.ClearDisplay()
	}
}
