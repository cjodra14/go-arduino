package main

import (
	"machine"

	"tinygo.org/x/drivers/hd44780i2c"
)

const (
	carriageReturn = 0x0D
	homeCommand    = "#home"
	clearCommand   = "#clear"
)

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
		Width:       16,
		Height:      2,
		CursorOn:    false,
		CursorBlink: false,
	}); err != nil {
		println("err failed to configure display")
	}

	homeScreen(lcd)

	var commandBuffer string
	var commandIndex uint8

	commandStart := false

	hadInput := false

	for {
		if uart.Buffered() == 0 {
			continue
		}

		if !hadInput {
			hadInput = true
			clearDisplay(lcd)
		}

		data, err := uart.ReadByte()
		if err != nil {
			println(err.Error())
		}

		if string(data) == "#" {
			commandStart = true
			uart.Write([]byte("\ncommand started\n"))
		}

		if commandStart {
			commandBuffer += string(data)
			commandIndex++
		}

		switch commandBuffer {
		case homeCommand:
			uart.WriteByte(data)
			homeScreen(lcd)
			commandStart = false
			commandIndex = 0
			commandBuffer = ""

			continue
		case clearCommand:
			uart.WriteByte(data)
			clearDisplay(lcd)
			commandStart = false
			commandIndex = 0
			commandBuffer = ""

			continue
		}

		if commandIndex > 5 {
			commandStart = false
			commandIndex = 0
			commandBuffer = ""
			uart.Write([]byte("\nresetting command state\n"))
		}

		if data == carriageReturn {
			lcd.Print([]byte("\n"))
			uart.Write([]byte("\r\n"))

			continue
		}

		lcd.Print([]byte{data})
		uart.WriteByte(data)
	}
}

func homeScreen(lcd hd44780i2c.Device) {
	println("\nexecuting command home screen\n")
	clearDisplay(lcd)
	lcd.Print([]byte("TinyGo UART \n CLI"))
}

func clearDisplay(lcd hd44780i2c.Device) {
	println("\nexecuting command clear display\n")
	lcd.ClearDisplay()
}
