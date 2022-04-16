package main

import "fmt"

func main() {
	a := "Hello World!"
	he := Haha(a)
	fmt.Println(he)
}

func Haha(a string) string {
	return a
}
