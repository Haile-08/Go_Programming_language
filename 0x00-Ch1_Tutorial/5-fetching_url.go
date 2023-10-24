// Fetch prints the content fount at a url
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:]{
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", res)
	}
}
