package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var choice string
var height int

func main() {

	fmt.Println("Do You want a star?")
	fmt.Println("Yes / No")
	scanner.Scan()
	choice = string(scanner.Text())

	fmt.Println("Tell me how big should the tree be?")
	fmt.Scan(&height)

	starOption()

	treeGenerator()

}

func starOption() {

	if choice == "Yes" {
		for i := 1; i < (height * 2); i++ {
			if i == height {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

}

func treeGenerator() {

	for y := height; y > 0; y-- {
		for x := 1; x < (height * 2); x++ {
			radius := height - y
			if x == height {
				fmt.Print("X")
			} else if (x >= height-radius) && (x < height) {
				fmt.Print("x")
			} else if (x > height) && (x <= height+radius) {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for j := 1; j < (height * 2); j++ {
		if j == height {
			fmt.Print("I")
		} else {
			fmt.Print(" ")
		}

	}

	fmt.Println()

}
