// prints the boiling point of water
package main

import "fmt"

const boilingF = 212.0

func main(){
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g0F or %g0c\n", f, c)
	//Output:
	//boiling point = 2120f or 1000c
}
