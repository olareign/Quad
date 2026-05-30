package main

import "fmt"

func main() {
	QuadA(10,6)
}

// START
//   ↓
// Are x and y both positive?
//   NO → print nothing, stop
//   YES → continue
//   ↓
// Print TOP line
//   ↓
// Is y greater than 1?
//   NO → stop here (only 1 row needed)
//   YES → continue
//   ↓
// Print (y - 2) MIDDLE rows
//   ↓
// Print BOTTOM line (same as top)
// DONE

func QuadA(x, y int) {
    // Step 1: guard — do nothing if invalid
    if x <= 0 || y <= 0 {
        return
    }

    // Step 2: build the top/bottom line as a string
    topLine := "o"
    if x > 1 {
        for i := 0; i < x-2; i++ {
            topLine += "-"
        }
        topLine += "o"
    }

    // Step 3: build the middle row as a string
    middleRow := "|"
    if x > 1 {
        for i := 0; i < x-2; i++ {
            middleRow += " "
        }
        middleRow += "|"
    }

    // Step 4: print top line
    fmt.Println(topLine)

    // Step 5: if height > 1, print middles then bottom
    if y > 1 {
        for i := 0; i < y-2; i++ {
            fmt.Println(middleRow)
        }
        fmt.Println(topLine) // bottom = same as top
    }
}