package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", KB, KB) // kb
	fmt.Printf("%d\t\t\t%b\n", MB, MB) // mb
	fmt.Printf("%d\t\t%b\n", GB, GB)   // gb
	fmt.Printf("%d\t\t%b\n", TB, TB)   // tb
	fmt.Printf("%d\t%b\n", PB, PB)     // pb
	fmt.Printf("%d\t%b\n", EB, EB)     // eb
}
