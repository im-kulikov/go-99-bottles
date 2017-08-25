package main

import (
	"fmt"
)

func main() {
	for num := 99; num >= 0; num-- {
		if num > 0 {
			if num > 1 {
				fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", num, num)
				fmt.Printf("Take one down and pass it around, %d bottle of beer on the wall.\n", num-1)
			} else {
				fmt.Printf("%d bottle of beer on the wall, %d bottle of beer.\n", num, num)
				fmt.Println("Take one down and pass it around, no more bottles of beer on the wall.")
			}
			println()
		} else {
			fmt.Println("No more bottles of beer on the wall, no more bottles of beer. ")
			fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
		}
	}
}
