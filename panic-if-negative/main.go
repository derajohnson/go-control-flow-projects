package main

import "fmt"

func main() {
	val := 0

	for {

		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		if val > 0 {
			fmt.Printf("You entered %v ", val)

		} else if val == 0 {
			fmt.Printf("0 is neither negative nor positive")

		} else {
			panic("Value is negative")
		}
	}

}
