package main

import "fmt"

func star_pattern(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			if i == 0 || j == 0 || i == j || i == n-1 || j == n-1 || i+j == n-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")

			}
		}
		fmt.Println()
	}
}
func main() {
	var n int
	fmt.Println("enter number:")
	fmt.Scan(&n)
	star_pattern(n)
}
