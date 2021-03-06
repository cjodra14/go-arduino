package main

import (
	"machine"

	"tinygo.org/x/drivers/hd44780i2c"
)

const carriageReturn = 0x0D

var (
	uart = machine.UART0
)

func main() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	lcd := hd44780i2c.New(i2c, 0x27)
	if err := lcd.Configure(hd44780i2c.Config{
		Width:  16,
		Height: 2,
	}); err != nil {
		println("err failed to configure display")
	}

	lcd.Print([]byte("Type to print:"))
	hadInput := false

	for {
		if uart.Buffered() == 0{
			continue
		}

		if !hadInput{
			hadInput = true
			lcd.ClearDisplay()
		}

		data, err := uart.ReadByte()
		if err != nil{
			println(err.Error())
		}

		if data == carriageReturn{
			lcd.Print([]byte("\n"))
			uart.Write([]byte("\r\n"))
			continue
		}

		lcd.Print([]byte{data})
		uart.WriteByte(data)
	}
}
