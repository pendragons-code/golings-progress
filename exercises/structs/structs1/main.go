// structs1
// Make me compile!
//
package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	var person Person
	person.age = 10
	person.name = "pentose"
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
