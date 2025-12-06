# 🧩 Project: `quadchecker` — Beginner Guide

-----

## 🧠 Step 1: What It Does (In Plain English)

The **`quadchecker`** is a detective program.
It reads a **shape** (text) that is sent to it.
It **analyzes** that shape and says:

> "Hey, this shape looks exactly like what quadA makes when width is 3 and height is 3."

If the shape matches multiple quads (like `quadC`, `quadD`, and `quadE` often do for `1x1` sizes), it lists **all of them** alphabetically.
If it looks like nothing it knows, it says **"Not a quad function"**.

-----

## 🧱 Step 2: How Data Moves (Visual Diagram)

Let’s visualize the "Pipe" `|` command you use to run this.

### 🧩 1️⃣ The Pipe Connection

You type: `./quadA 3 3 | go run main.go`

```
+---------------+        (The Pipe | )       +-------------------+
|  ./quadA 3 3  |  ----------------------->  |   go run main.go  |
| (Makes Shape) |   Sends shape text to...   | (The Investigator)|
+---------------+                            +-------------------+
```

### 🧩 2️⃣ `quadchecker` Reads the Input

It receives this text:

```
o-o
| |
o-o
```

### 🧩 3️⃣ It Measures the Shape

It looks at the text and calculates:

```
Width (x)  = 3 characters
Height (y) = 3 lines
```

### 🧩 4️⃣ It Re-Creates and Compares

It internally runs its own version of every Quad using those measurements (`3` and `3`) to see which one matches the input.

```
Input Shape   vs   CheckQuadA(3,3)?  --> MATCH! ✅
Input Shape   vs   CheckQuadB(3,3)?  --> No.
Input Shape   vs   CheckQuadC(3,3)?  --> No.
Input Shape   vs   CheckQuadD(3,3)?  --> No.
Input Shape   vs   CheckQuadE(3,3)?  --> No.
```

### 🧩 5️⃣ It Prints the Result

```
[quadA] [3] [3]
```

-----

## 🧠 Step 3: Project Setup

To avoid the "redeclared" errors, we put the logic inside `main.go`.

### 1\. Folder Structure

```
quadchecker/
 ├── go.mod
 └── main.go  <-- All logic is here
```

### 2\. Initialize Module

Open terminal in the folder:

```bash
go mod init quadchecker
```

### 3\. The Code

Paste the code below into `main.go`. This version is **Self-Contained**. It doesn't need external files, so it won't crash with "redeclared" errors.

-----

## ✅ Final Code (`main.go`)

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// --- 1. The Generator Functions ---
// These create the shapes strictly for comparison strings.

func CheckQuadA(x, y int) string {
	if x <= 0 || y <= 0 { return "" } // Safety check
	var res strings.Builder           // Efficient string builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			// Logic for Quad A
			if (row == 1 || row == y) && (col == 1 || col == x) {
				res.WriteString("o")
			} else if row == 1 || row == y {
				res.WriteString("-")
			} else if col == 1 || col == x {
				res.WriteString("|")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n") // Newline at end of row
	}
	return res.String()
}

func CheckQuadB(x, y int) string {
	if x <= 0 || y <= 0 { return "" }
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			// Logic for Quad B
			if row == 1 && col == 1 {
				res.WriteString("/")
			} else if row == 1 && col == x {
				res.WriteString("\\")
			} else if row == y && col == 1 {
				res.WriteString("\\")
			} else if row == y && col == x {
				res.WriteString("/")
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteString("*")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func CheckQuadC(x, y int) string {
	if x <= 0 || y <= 0 { return "" }
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				res.WriteString("A")
			} else if row == y && (col == 1 || col == x) {
				res.WriteString("C")
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteString("B")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func CheckQuadD(x, y int) string {
	if x <= 0 || y <= 0 { return "" }
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 || row == y) && col == 1 {
				res.WriteString("A")
			} else if (row == 1 || row == y) && col == x {
				res.WriteString("C")
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteString("B")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

func CheckQuadE(x, y int) string {
	if x <= 0 || y <= 0 { return "" }
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == y && col == x) {
				res.WriteString("A")
			} else if (row == 1 && col == x) || (row == y && col == 1) {
				res.WriteString("C")
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteString("B")
			} else {
				res.WriteString(" ")
			}
		}
		res.WriteString("\n")
	}
	return res.String()
}

// --- 2. The Main Investigator ---

func main() {
	// A. Read everything piped into the program
	reader := bufio.NewReader(os.Stdin)
	inputBytes, _ := io.ReadAll(reader)
	input := string(inputBytes)

	// B. If empty, fail immediately
	if input == "" {
		fmt.Println("Not a quad function")
		return
	}

	// C. Calculate Dimensions
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	y := len(lines)
	x := 0
	if y > 0 {
		x = len(lines[0])
	}

	// Safety check for bad dimensions
	if x == 0 || y == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// D. The Comparison Loop
	var matches []string

	// Check against every quad function
	if input == CheckQuadA(x, y) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if input == CheckQuadB(x, y) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if input == CheckQuadC(x, y) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if input == CheckQuadD(x, y) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if input == CheckQuadE(x, y) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	// E. Print Results
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		// Join multiple matches with " || "
		fmt.Println(strings.Join(matches, " || "))
	}
}
```

-----

# ✅ MEMORIZE THIS 2-MINUTE EXPLANATION

## 🔹 What the program does (1 Sentence)

> "The program reads a shape from standard input, calculates its width and height, generates all 5 possible quads with those dimensions, and prints the names of the ones that match."

## 🔹 How the logic works (3 Steps)

1.  **Read & Measure:** I read the input text and count lines (height) and characters (width).
2.  **Generate:** I use internal functions (`CheckQuadA`, etc.) to build strings for all 5 quads using those measurements.
3.  **Compare:** I check `if input == generatedString` for each quad. If it matches, I print it.

-----

# ✅ AUDIT QUICK ANSWERS

### **Q: Why did you use `strings.Builder`?**

**A:** "Because in Go, strings are immutable (cannot change). If I used `+` to add characters in a loop, it would be very slow and use too much memory. `strings.Builder` is the efficient way to build text."

### **Q: Why `os.Stdin`?**

**A:** "Because this program is designed to work with a **pipe** (`|`). The output of `./quadA` becomes the input of my checker through Stdin (Standard Input)."

### **Q: Why do you handle `matches` as a slice (array)?**

**A:** "Because sometimes a shape matches multiple Quads (like `1x1`). Using a slice allows me to store all of them and then print them joined by `||` at the end."

### **Q: How do you handle the error where `main` is redeclared?**

**A:** "I moved my checking logic into `main.go` functions named `CheckQuadA`, `CheckQuadB`, etc. This way, `main.go` is independent and doesn't conflict with the other files."

-----

# ✅ CRISIS MODE (If you panic, say this)

> "It's a reverse-engineering tool. It takes the output (the drawing), measures it, and checks it against all known Quad functions to see which one created it."
