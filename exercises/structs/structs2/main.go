// structs2
// Make me compile!
//
package main

import "fmt"

type Person struct {
	// don't just create the phone field here. embed a new struct
	name string
	age int
}

// I sure as hell hope this is what they mean by embed a new structure
type ContactDetails struct {
	Person
	phone string
}

func main() {
	// contactDetails := ContactDetails{}
	var person = ContactDetails{Person: Person {"John", 10}}
	person.phone = "11111111"
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
