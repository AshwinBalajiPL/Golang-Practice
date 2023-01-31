package main

import "fmt"

func main() {
	s := []int{}

	for i := 0; i <= 10; i++ {
		s = append(s, i)
		if s[i]%2 == 0 {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}
