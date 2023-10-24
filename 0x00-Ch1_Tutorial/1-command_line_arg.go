/*
A go program that implement the command line argument
*/
package main

import (
	"fmt"
	"os"
)

func main(){
	//The main func
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
