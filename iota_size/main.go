package main

import "fmt"

const (
	_ = iota
	KB
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d\t%b", KB, KB) // kb
	fmt.Printf("%d\t%b", MB, MB) // mb
	fmt.Printf("%d\t%b", GB, GB) // gb
	fmt.Printf("%d\t%b", TB, TB) // tb
	fmt.Printf("%d\t%b", PB, PB) // pb
	fmt.Printf("%d\t%b", EB, EB) // eb
}
