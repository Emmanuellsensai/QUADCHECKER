package main

import (
	"fmt"
	"os"
	"strconv"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			isFirstRow := row == 1
			isLastRow := row == y
			isFirstCol := col == 1
			isLastCol := col == x

			isTopLeft := isFirstRow && isFirstCol
			isTopRight := isFirstRow && isLastCol
			isBottomLeft := isLastRow && isFirstCol
			isBottomRight := isLastRow && isLastCol

			isTopEdge := isFirstRow && col > 1 && col < x
			isBottomEdge := isLastRow && col > 1 && col < x
			isLeftEdge := isFirstCol && row > 1 && row < y
			isRightEdge := isLastCol && row > 1 && row < y

			if isTopLeft || isTopRight || isBottomLeft || isBottomRight {
				fmt.Print("o")
			} else if isTopEdge || isBottomEdge {
				fmt.Print("-")
			} else if isLeftEdge || isRightEdge {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// 1. Check for arguments
	if len(os.Args) != 3 {
		return
	}

	// 2. Parse arguments
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	// 3. Call the function if args are valid numbers
	if err1 == nil && err2 == nil {
		QuadA(x, y)
	}
}
