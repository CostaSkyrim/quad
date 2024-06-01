package piscine

import "fmt"

func QuadA(x,y int) {
        if x > 0 && y > 0 {
        for row := 0; row < y; row++ {
                for col := 0; col < x; col++ {
                        if (row == 0 && col == 0) || (row == 0 && col == x-1) || (row == y-1 && col == 0) || (row == y-1 && col == x-1) {
                                // Print corners
                                fmt.Print("o")
                        } else if row == 0 || row == y-1 {
                                // Print top or bottom edges
                                fmt.Print("-")
                        } else if col == 0 || col == x-1 {
                                // Print left or right edges
                                fmt.Print("|")
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

