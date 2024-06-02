package piscine

import "fmt"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for row := 0; row < y; row++ {
			for col := 0; col < x; col++ {
				if row == 0 && col == 0 {
					fmt.Print("/")
					// Print the top left corner
				} else if (row == 0 && col == x-1) || (row == y-1 && col == 0) {
					fmt.Print("\\")
					// Print the top right and the bottom left corners
				} else if row == y-1 && col == x-1 {
					fmt.Print("/")
					// Print the bottom right corner
				} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
					fmt.Print("*")
					// Print the edges
				} else {
					fmt.Print(" ")
					// Print the inside of the rectangle
				}
			}
			fmt.Println()
			// Newline after each row
		}
	} else {
		return
	}
}
