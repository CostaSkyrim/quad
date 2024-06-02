package piscine

import "fmt"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for row := 0; row < y; row++ {
			for col := 0; col < x; col++ {
				if (row == 0 && col == 0) || (row == 0 && col == x-1) || (row == y-1 && col == 0) || (row == y-1 && col == x-1) {
					fmt.Print("o")
					// Print the corners
				} else if row == 0 || row == y-1 {
					fmt.Print("-")
					// Print the top and the bottom edges
				} else if col == 0 || col == x-1 {
					fmt.Print("|")
					// Print the left and the right edges
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
