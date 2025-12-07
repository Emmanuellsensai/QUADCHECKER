func CheckQuadE(x, y int) string {
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
