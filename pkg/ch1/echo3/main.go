package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))
	difference := time.Now().Sub(start)
	fmt.Println("This took", difference, "seconds")
}
