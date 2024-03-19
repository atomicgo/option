package option_test

import (
	"fmt"

	"atomicgo.dev/option"
)

type Person struct {
  Name option.StringOption
  Age  option.IntOption
}

func Example_demo() {
  
}

func ExampleHelloWorld() {
	fmt.Println(option.HelloWorld())
	// Output: Hello, World!
}
