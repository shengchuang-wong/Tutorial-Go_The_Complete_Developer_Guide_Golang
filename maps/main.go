package main

import "fmt"

func main() {

	// Declare map
	// Method 1 - key = string, value = string
	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00FF00",
	// 	"blue":  "#0000FF",
	// }

	// Method 2
	// var colors map[string]string

	// Method 3
	// colors := make(map[string]string)
	// colors["white"] = "#FFF"
	// colors["black"] = "#000"
	// delete(colors, "black")

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
