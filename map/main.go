package main

import (
	"fmt"
)

func main() {
	// var colors map[int]string

	// colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hexcode for", color, "is", hex)
	}
}
