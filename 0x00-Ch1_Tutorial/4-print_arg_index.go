//Print the index and arg
package main

import (
	"fmt"
	"os"
)

func main(){
	for i, arg := range os.Args[1:]{
		fmt.Printf("index: %v, arg: %v \n", i, arg)
	}
}
