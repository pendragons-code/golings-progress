// range1
// Make me compile!
//
package main

import "fmt"

func main() {
	evenNumbers := []int{2, 4, 6, 8, 10}

	for v := 0; v < len(evenNumbers); v++ {
		fmt.Printf("%d is even\n", v)
	}
}
