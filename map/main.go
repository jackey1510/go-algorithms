package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#97cc92",
	}
	colors["blue"] = "#298abe"

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, code := range c {
		c[color] = ""
		fmt.Printf("%s : %s\n", color, code)
	}
}
