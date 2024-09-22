// structs3
// Make me compile!
//
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}


// https://go101.org/article/type-embedding.html
func (p *Person) FullName() string {
	return p.firstName + " " + p.lastName
}



func main() {
	person := Person{firstName: "Maurício", lastName: "Antunes"}
	fmt.Printf("Person full name is: %s\n", person.FullName()) // here it must output Person full name is: Maurício Antunes
}
