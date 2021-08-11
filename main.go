package main

import "fmt"

func main () {
	//var colors map[string]string // initialize with zero value (empty map)
	//colors := make(map[string]string) // create a map with no values inside where all keys and values are string;

	colors := map[string]string{ // [keys]values
		"red":"#ff0000",
		"green":"#4bf745",
		"white":"#ffffff",
	}
	//colors["white"] = "#ffffff" // add to map
	//delete(colors, "white") // delete from map

	printMap(colors)
}

func printMap (c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}