package main

import "fmt"

func main() {
	for i := 1; i <= 100; i += 1 {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
