// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var string, seperator string

	for i := 0; i < len(os.Args); i += 1 {
		string += seperator + os.Args[i]
		seperator = " "
	}

	fmt.Println(string)
	difference := time.Now().Sub(start)
	fmt.Println("This took", difference, "seconds")
}
