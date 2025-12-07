package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1. Read input from the pipe
	reader := bufio.NewReader(os.Stdin)
	inputBytes, _ := io.ReadAll(reader)
	input := string(inputBytes)

	if input == "" {
		fmt.Println("Not a quad function")
		return
	}

	// 2. Determine dimensions
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	y := len(lines)
	x := 0
	if y > 0 {
		x = len(lines[0])
	}

	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// 3. Compare input against our internal string generators
	var matches []string

	if input == QuadA(x, y) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if input == QuadB(x, y) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if input == QuadC(x, y) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if input == QuadD(x, y) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if input == QuadE(x, y) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	// 4. Print results
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}


func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
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
				res.WriteString("o")
			} else if isTopEdge || isBottomEdge {
				res.WriteString("-")
			} else if isLeftEdge || isRightEdge {
				res.WriteString("|")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func QuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
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
				res.WriteString("/")
			} else if isTopRight {
				res.WriteString("\\")
			} else if isBottomLeft {
				res.WriteString("\\")
			} else if isBottomRight {
				res.WriteString("/")
			} else if isTopEdge || isBottomEdge || isLeftEdge || isRightEdge {
				res.WriteString("*")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func QuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
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

			// CORRECTED LOGIC: Check explicit corners in order
			if isTopLeft {
				res.WriteString("A")
			} else if isTopRight {
				res.WriteString("A")
			} else if isBottomLeft {
				res.WriteString("C")
			} else if isBottomRight {
				res.WriteString("C")
			} else if isTopEdge || isBottomEdge || isLeftEdge || isRightEdge {
				res.WriteString("B")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func QuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
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
				res.WriteString("A")
			} else if isTopRight {
				res.WriteString("C")
			} else if isBottomLeft {
				res.WriteString("A")
			} else if isBottomRight {
				res.WriteString("C")
			} else if isTopEdge || isBottomEdge || isLeftEdge || isRightEdge {
				res.WriteString("B")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func QuadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
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
				res.WriteString("A")
			} else if isTopRight {
				res.WriteString("C")
			} else if isBottomLeft {
				res.WriteString("C")
			} else if isBottomRight {
				res.WriteString("A")
			} else if isTopEdge || isBottomEdge || isLeftEdge || isRightEdge {
				res.WriteString("B")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

// --- Main Execution ---
