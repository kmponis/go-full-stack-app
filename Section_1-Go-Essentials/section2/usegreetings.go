// In this example we use the greetingspackage
package main

import (
	"fmt"
	// Original import: "github.com/EngineerKamesh/gofullstack/volume1/section2/greetingspackage"
	"go-full-stack-app/Section_1-Go-Essentials/section2/greetingspackage"
)

func main() {
	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()
	fmt.Println("The value of the Magic Number is:", greetingspackage.MagicNumber)
}
