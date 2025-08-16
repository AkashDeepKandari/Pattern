package main

import "fmt"

func printButterfly(n int) {
	// Upper half
	for i := 1; i <= n; i++ {
		// Left stars
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		// Spaces
		for j := 1; j <= 2*(n-i); j++ {
			fmt.Print(" ")
		}
		// Right stars
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Lower half
	for i := n; i >= 1; i-- {
		// Left stars
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		// Spaces
		for j := 1; j <= 2*(n-i); j++ {
			fmt.Print(" ")
		}
		// Right stars
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Enter N: ")
	fmt.Scan(&n)
	printButterfly(n)
}
