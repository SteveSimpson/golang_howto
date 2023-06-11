// just a simple starter program
package main

import "fmt"

func main() {
	const name, age = "Todd", 42

	who := "GOphers"

	emoji := "😜😎"

	fmt.Printf("Hello %s%s!\n", who, emoji)

	fmt.Printf("%s is %d\n", name, age)

	fmt.Printf("%v is %v\ttypes: %T - %T", name, age, name, age)

	fmt.Print(`
	There is no room on this planet for anything less than a miracle
We gather here today to revel in the rebellion of a silent tongue
Every day, we lean forward into the light of our brightest designs
       & cherish...
`)

	fmt.Printf("Decimal: %d, Binary: %b, hex: %x\n", age, age, age)
}
