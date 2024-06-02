package piscine

import "fmt"

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for row := 0; row < y; row++ {
			for col := 0; col < x; col++ {
				if (row == 0 && col == 0) || (row == 0 && col == x-1) {
					fmt.Print("A")
					// Print the top corners
				} else if row == y-1 && col == 0 || row == y-1 && col == x-1 {
					fmt.Print("C")
					// Print the bottom corners
				} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
					fmt.Print("B")
					// Print the edges
				} else {
					fmt.Print(" ")
					// Print inside of the rectangle
				}
			}
			// Newline after each row
			fmt.Println()
		}
	} else {
		return
	}
}
