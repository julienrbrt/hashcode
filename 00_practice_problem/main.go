package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 84 is T
// 77 is M

var row int
var col int
var minIngredient int // min ingredient
var maxArea int
var pizza []string
var results [][]int

func pizzacut(path string) ([]string, int, int, int, int) {

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
	// _ as always sure to get good value
	row, _ = strconv.Atoi(arguments[0])
	col, _ = strconv.Atoi(arguments[1])
	minIngredient, _ = strconv.Atoi(arguments[2])
	maxArea, _ = strconv.Atoi(arguments[3])

	return pizza, row, col, minIngredient, maxArea
}

func main() {
	pizza, row, col, minIngredient, maxArea = pizzacut("a_example.in")

	for i := range pizza {
		beg := 0
		end := 0
		mushroom := 0
		tomato := 0

		for end < col { // for
			if pizza[i][end] == 77 {
				mushroom++
			} else if pizza[i][end] == 84 {
				tomato++
			}
			end++
		}

		if end-beg > maxArea { // if slide too big, remove one ingredient
			if pizza[i][beg] == 77 {
				mushroom--
			} else if pizza[i][beg] == 84 {
				tomato--
			}
			beg++
		}

		if end-beg <= maxArea && mushroom >= minIngredient && tomato >= minIngredient {
			results = append(results, []int{i, beg, i, end - 1}) // TODO: fix the one column only
			beg = end
			tomato = 0
			mushroom = 0
		}
	}
	fmt.Println(results)
}
