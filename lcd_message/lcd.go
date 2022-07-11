package main

import (
	"machine"

	"tinygo.org/x/drivers/hd44780i2c"
)

var data = []byte("Hello world! \nLCD 16x02")

func main() {

	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(i2c, 0x27)
	lcd.Configure(hd44780i2c.Config{Width: 16, Height: 2})
	// lcd.BacklightOn(false)
	lcd.Print(data)
}
