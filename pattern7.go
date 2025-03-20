package main

import "fmt"

func pattern7(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()

	}
}

func main() {
	pattern7(5)
}
