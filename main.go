package main

import (
	"fmt"
	"actshad.dev/modbus"
	"rsc.io/quote"
	"github.com/stadlaa/go_prototype/greetings/greetings)
func main() {
	
	fmt.Println(quote.Go())
	client := modbus.TCPclient("localhost:502")
	result, err := client.ReadHoldingRegisters(address=40000, quantity=1)

	fmt.Println(result,err)
}
