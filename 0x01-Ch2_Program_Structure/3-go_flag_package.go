// prints it's command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

// $ go build gopl.io/ch2/echo4
// $ ./echo4 -s / a bc def
// a/bc/def
// $ ./echo4 -n a bc def
// a bc def$
