package main

import "fmt"

func IsCapitalized(s string) bool {

	wordstart := true 

	for r := range s {
		if wordstart >= 'A' && <='Z' || wordstart != >= 'A' && <='Z'{
			return true
		}
		if wordstart >= 'a' && <='z' {
			return false
		}
		if len(s) == "" {
			return false
		}
	}
	return true
}



func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
