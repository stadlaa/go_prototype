package main

import (
	"fmt"
	"actshad.dev/modbus"
	"greetings"
	"rsc.io/quote")
func main() {
	
	fmt.Println(quote.Go())
	client := modbus.TCPclient("localhost:502")
	result, err := client.ReadHoldingRegisters(address=40000, quantity=1)

	fmt.Println(result,err)
}
