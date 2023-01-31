package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type japaneseBot struct{}

func main() {
	eb := englishBot{}
	jb := japaneseBot{}

	printGreeting(eb)
	printGreeting(jb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello!!"
}

func (jb japaneseBot) getGreeting() string {
	return "Konnichiwa!!"
}
