package main

import (
	"./histlib"
	"fmt"
)

//This is just a stub to show how to use the histlib
func main() {
	fmt.Println(histlib.Rot("CAESAR", histlib.PRINTABLE_ASCII, 3))
}
