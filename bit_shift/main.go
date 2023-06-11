package main

import "fmt"

func main() {
	value := 1
	fmt.Println("Loop\tDecimal\tHex\tBinary")
	for i := 0; i < 62; i++ {
		fmt.Printf("%d\t%d\t%x\t%b\n", i, value, value, value)
		value = value << 1
	}
}
