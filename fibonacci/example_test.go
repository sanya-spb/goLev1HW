package fibonacchi_test

import (
	"fmt"

	fibonacchi "github.com/sanya-spb/goLev1HW/fibonacci"
)

func ExampleFibonacciR() {
	var inputR, inputM uint8 = 11, 10

	FibR := fibonacchi.FibonacciR(inputR)
	FibM := fibonacchi.FibonacciM(inputM)
	fmt.Println(FibR - FibM)

	// Output:
	// 34
}

func ExampleFibonacciM() {
	fmt.Println(fibonacchi.FibonacciM(uint8(10)))

	// Output:
	// 55
}
