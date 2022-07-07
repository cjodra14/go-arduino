package main

import (
	"machine"
	"time"
)

var dataFlag bool = false

var dataSlice []byte

func main() {
	for {
		//read bytes from RX buffer
		readByte, err := machine.UART0.ReadByte()
		if readByte != 0 {
			dataSlice = append(dataSlice, readByte)
		}
		if err != nil {
			dataFlag = true
		}

		time.Sleep(time.Millisecond * 350)

		if dataFlag {
			if len(dataSlice) > 0 {
				print("Your data is : ")
			}
			//Write bytes to UART
			machine.UART0.Write(dataSlice)
			dataFlag = false
			dataSlice = nil
		}

	}
}
