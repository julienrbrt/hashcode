package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pizzacut(path string) {

	var r string
	var c string
	var l string
	var h string
	var pizza []string

	// read file
	file, err := os.Open(path) // read file
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close() // close file

	bf := bufio.NewScanner(file)

	bf.Scan()              // read 1st line file
	firstline := bf.Text() // get first line of file

	for bf.Scan() {
		pizza = append(pizza, bf.Text()) // get pizza
	}

	arguments := strings.Split(firstline, " ") // split first line

	// assign lines to argument
	r = arguments[0]
	c = arguments[1]
	l = arguments[2]
	h = arguments[3]

	fmt.Println(r + c + l + h)
	fmt.Println(pizza)
}

func main() {
	pizzacut("b_small.in")
}
