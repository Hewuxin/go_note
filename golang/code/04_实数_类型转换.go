package main

import (
	"fmt"
)

func main() {
	message := "shalom"
	for count := 0; count < len(message); count++ {
		fmt.Printf("%c\n", message[count])
	}
}
