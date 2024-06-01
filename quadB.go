package piscine

import "fmt"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for row := 0; row < y; row++ {
			for col := 0; col < x; col++ {
				if row == 0 && col == 0 {
					// Print the left corners
					fmt.Print("/")
				} else if (row == 0 && col == x-1) || (row == y-1 && col == 0) {
					// Print the right corners
					fmt.Print("\\")
				} else if row == y-1 && col == x-1 {
					// Print the right corners
					fmt.Print("/")
				} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
					// Print top, bottom, left and right edges
					fmt.Print("*")
				} else {
					// Print inside of the rectangle
					fmt.Print(" ")
				}
			}
			// Newline after each row
			fmt.Println()
		}
	} else {
		return
	}
}
