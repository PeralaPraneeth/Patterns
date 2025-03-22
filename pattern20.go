package main

import "fmt"

func pattern20(n int) {
	spaces := 2 * (n - 1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}

		for j := i; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()

		spaces -= 2
	}

	space := 2
	for i := 1; i < n; i++ {
		for j := n; j >= i+1; j-- {
			fmt.Print("*")
		}

		for j := 1; j <= space; j++ {
			fmt.Print(" ")
		}

		for j := n; j >= i+1; j-- {
			fmt.Print("*")
		}
		fmt.Println()

		space += 2
	}

}

func main() {
	pattern20(5)
}
