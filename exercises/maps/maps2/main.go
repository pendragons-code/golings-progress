// maps2
// Make me compile!
//
package main

import "fmt"

func main() {
	m := map[string]int{
		"John": 10,
		"Ana": 11, // it seems that it expects you to end with a comma for further manips
	}
	fmt.Printf("John is %d and Ana is %d", m["John"], m["Ana"])
}
