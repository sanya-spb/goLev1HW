package config_test

import (
	"fmt"

	"github.com/sanya-spb/goLev1HW/utils/config"
)

func ExampleIsIPv4Net_true() {
	IP := "10.0.0.1"
	if ok := config.IsIPv4Net(IP); ok {
		fmt.Println("valid")
	} else {
		fmt.Println("no valid")
	}

	// Output:
	// valid
}

func ExampleIsIPv4Net_false() {
	IP := "a.b.c.d"
	if ok := config.IsIPv4Net(IP); ok {
		fmt.Println("valid")
	} else {
		fmt.Println("no valid")
	}

	// Output:
	// no valid

}

func ExampleIsURL_true() {
	url := "http://google.com/?go=lang#ok"
	if ok := config.IsURL(url); ok {
		fmt.Println("valid")
	} else {
		fmt.Println("no valid")
	}

	// Output:
	// valid
}

func ExampleIsURL_false() {
	url := "./google.txt"
	if ok := config.IsURL(url); ok {
		fmt.Println("valid")
	} else {
		fmt.Println("no valid")
	}

	// Output:
	// no valid
}
