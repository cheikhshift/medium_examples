package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
}

// MarshalText implements the encoding.TextMarshaler interface.
func (p Person) MarshalText() ([]byte, error) {

	fullName := fmt.Sprintf("%s %s", p.FirstName, p.LastName)
	return []byte(fullName), nil

}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (p *Person) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}
	s := string(text)
	parts := strings.Split(s, " ")

	// must have atleast has 2 names
	if len(parts) < 2 {
		return errors.New("Please specify a first and last name")
	}

	*p = Person{
		FirstName: parts[0],
		LastName:  parts[1],
	}
	return nil
}

func main() {
	var p Person
	defaultValue := Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	flag.TextVar(&p, "person", defaultValue, "Enter the first and last name of person")

	flag.Parse()

	fmt.Println(p)
}
