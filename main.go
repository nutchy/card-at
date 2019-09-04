package main

import "fmt"

func main() {
	fmt.Printf("Result is : %s \n", cardAt(51))
}

// cardAt -
func cardAt(n int) string {

	if n < 0 || n > 51 {
		return `Incorrect input number.`
	}

	// Declaration rank and suit of face card
	rank := []string{"2", "3", "4", "5", "6", "7", "8", "9", "0", "J", "Q", "K", "A"}

	// Club (♧), Diamond (♢), Heart (♡), Spade (♤)
	suit := []string{"C", "D", "H", "S"}

	// Use modulo and division to finding an index of rank & suit
	return fmt.Sprintf("%s%s", rank[n%13], suit[n/13])
}
