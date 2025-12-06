package main

import (
	"fmt"
	"os"
	"strconv"
)

func QuadB(x, y int) {
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

			if isTopLeft {
				fmt.Print("/")
			} else if isTopRight {
				fmt.Print("\\")
			} else if isBottomLeft {
				fmt.Print("\\")
			} else if isBottomRight {
				fmt.Print("/")
			} else if isTopEdge || isBottomEdge {
				fmt.Print("*")
			} else if isLeftEdge || isRightEdge {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	if err1 == nil && err2 == nil {
		QuadB(x, y)
	}
}
