// just a simple starter program
package main

import "fmt"

func main() {
	const name, age = "Todd", 45

	who := "GOphers"

	emoji := "😜😎"

	fmt.Printf("Hello %s%s!\n", who, emoji)

	fmt.Printf("%s is %d\n", name, age)
}
