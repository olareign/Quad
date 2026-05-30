// START
//   ↓
// Are x and y both positive?
//   NO  → print nothing, stop
//   YES → continue
//   ↓
// Print TOP line  ( ABBBC )
//   ↓
// Is y greater than 1?
//   NO  → stop (only 1 row)
//   YES → continue
//   ↓
// Print (y - 2) MIDDLE rows  ( B   B )
//   ↓
// Print BOTTOM line  ( CBBBA )  ← mirrored top
// DONE

package main

import "fmt"

func main() {
	QuadE(5,3)
}

func QuadE(x, y int) {
    // Step 1: guard
    if x <= 0 || y <= 0 {
        return
    }

    // Step 2: build top line  ABBBC
    topLine := 'A'
    if x > 1 {
        for i := 0; i < x-2; i++ {
            topLine += 'B'
        }
        topLine += 'C'
    }

    // Step 3: build middle row  B   B
    middleRow := 'B'
    if x > 1 {
        for i := 0; i < x-2; i++ {
            middleRow += ' '
        }
        middleRow += 'B'
    }

    // Step 4: build bottom line  CBBBA  (mirror of top)
    bottomLine := 'C'
    if x > 1 {
        for i := 0; i < x-2; i++ {
            bottomLine += 'B'
        }
        bottomLine += 'A'
    }

    // Step 5: print top
    z01.PrintRune(topLine)

    // Step 6: if height > 1, print middles then bottom
    if y > 1 {
        for i := 0; i < y-2; i++ {
            z01.PrintRune(middleRow)
        }
        z01.PrintRune(bottomLine)
    }
}