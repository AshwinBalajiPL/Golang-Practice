package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	var url string
	fmt.Print("Enter url : ")
	fmt.Scanln(&url)
	res, err := http.Get("https://" + url)

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	fmt.Println()

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	// fmt.Println("The size of the data : ", len(bs))

	return len(bs), nil
}
