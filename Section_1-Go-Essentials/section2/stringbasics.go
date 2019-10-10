package main

import "fmt"

func main() {

	var subject string = "Gopher"

	fmt.Println("First element of 'Yuhu' string: ", string("Yuhu"[0]))
	fmt.Println("String 'subject': " + subject)
	fmt.Printf("The first value(%v) of the subject string: %v\n", 0, string(subject[0]))
	fmt.Printf("The last value(%v) of the subject string: %v\n", len(subject)-1, string(subject[len(subject)-1]))

}
