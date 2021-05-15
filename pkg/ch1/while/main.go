package main

import "fmt"

func main() {
	i := 7

	for {
		fmt.Println(i)
		i += 1
		if i == 17 {
			break
		}
	}
}
