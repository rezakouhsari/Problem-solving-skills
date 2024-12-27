package main

import "fmt"

func main() {
	sum := 0
	for counter := 1; counter < 100; counter++ {
		if counter%2 != 0 {
			sum += counter
		}
	}
	fmt.Println(sum)
}
