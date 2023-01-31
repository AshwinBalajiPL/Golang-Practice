package main

import "fmt"

func main() {
	colors := map[string]string{
		"Red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	for _, hex := range colors {
		fmt.Print(hex)
	}
}
