package main

import "fmt"

func main() {
	n := 7 // Must be odd
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && (j == 0 || j >= n/2)) || // Top row
				(i < n/2 && (j == 0 || j == n/2)) || // Upper left vertical
				(i == n/2) || // Middle row
				(i > n/2 && (j == n/2 || j == n-1)) || // Lower right vertical
				(i == n-1 && (j <= n/2 || j == n-1)) { // Bottom row
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
