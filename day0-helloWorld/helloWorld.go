// Input Format
// A single line of text denoting  (the variable whose contents must be printed).

// Output Format
// Print Hello, World. on the first line, and the contents of  on the second line.

// Sample Input
// Welcome to 30 Days of Code!

// Sample Output
// Hello, World.
// Welcome to 30 Days of Code!

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	var message string
	message, _ = reader.ReadString('\n')

	fmt.Printf("Hello, World.\n%s", message)
}
