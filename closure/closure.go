package closure

import (
	"fmt"
	"time"
)

// Example01 Several anonymous functions as go routines
func Example01() {
	fmt.Println("Closure - Example 01: several anonymous functions as go routines")

	go func(s string) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d - Hello %s, you are inside loop!\n", i, s)
		}
	}("James")

	go func(s string) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d - Hello %s, you are inside loop!\n", i, s)
		}
	}("Moneypenny")

	go func(s string) {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("%d - Hello %s, you are inside loop!\n", i, s)
		}
	}("M")

	time.Sleep(5 * time.Second)
}
