package main

import (
	"fmt"     //formatting text & read formatted text
	"os"      //platform-independent OS var & functions
	"strings" //functions for string manipulations
)

func main() {
	//No need to specify type. Go can deduce & prevent assignment of other type
	who := "World!"       //short var declaration: both declare & initialize
	if len(os.Args) > 1 { // os.Args[0] is "hello" or "hello.exe"
		who = strings.Join(os.Args[1:], " ") //= is assignment operator
	}
	fmt.Println("Hello", who)
}
