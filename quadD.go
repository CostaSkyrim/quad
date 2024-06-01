package piscine

import "fmt"

func QuadD(x, y int) {
        if x > 0 && y > 0 {
                for row := 0; row < y; row++ {
                        for col := 0; col < x; col++ {
                                if (row == 0 && col == 0) || (col == 0 && row == y-1) {
                                                                                fmt.Print("A")
                                                                } else if (row == 0 && col == x-1) || (row == y-1 && col == x-1) {
                                                                                fmt.Print("C")
                                } else if row == 0 || row == y-1 || col == 0 || col == x-1 {
                                        // Print top, bottom, left and right edges
                                        fmt.Print("B")
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

