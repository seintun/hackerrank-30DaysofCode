package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.

	// Read and save an integer, double, and String to your variables.

	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	scanner.Scan()
	j, _ := strconv.ParseUint(scanner.Text(), 0, 64)
	scanner.Scan()
	e, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	var t string = scanner.Text()

	fmt.Println(i + j)
	fmt.Printf("%.1f", d+e)
	fmt.Printf("\n%s%s", s, t)
}
