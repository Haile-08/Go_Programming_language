// pointer function example
package main

import "fmt"

func main(){
	fmt.Println(f())
	fmt.Println(f() == f()) // call of f returns a distinct value
}

func f() *int{
	v := 1
	return &v
}
