package option_test

import (
	"fmt"

	"atomicgo.dev/option"
)

func Example_demo() {
	fmt.Println(option.HelloWorld())
	// Output: Hello, World!
}

func ExampleHelloWorld() {
	fmt.Println(option.HelloWorld())
	// Output: Hello, World!
}
