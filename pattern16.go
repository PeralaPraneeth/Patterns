package main

import "fmt"

func pattern16(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", 'A'+i)
		}
		fmt.Println()
	}
}

func main() {
	pattern16(5)
}
