package main

import "fmt"

func pattern8(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*n-(2*i+1); j++ {
			fmt.Print("*")
		}
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	pattern8(5)
}
