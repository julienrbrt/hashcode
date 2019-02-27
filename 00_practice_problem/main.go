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

var resultsCol [][]int
var resultsRow [][]int
var filename = "b_small"

func makePizza(path string) ([]string, int, int, int, int) {

	// read file
	file, err := os.Open(path + ".in") // read file
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close() // close file

	bf := bufio.NewScanner(file)

	bf.Scan()              // read 1st line file
	firstline := bf.Text() // get first line of file

	for bf.Scan() {
		pizza = append(pizza, bf.Text()) // get pizza, by default by rows
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

// calculate columns
func calculateColumns(pizza []string, row, col, minIngredient, maxArea int) [][]int {
	var results [][]int

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
	}

	return results
}

// create Columns
func createColumns(pizza []string, col int) []string {
	var newPizza []string

	for i := 0; i < col; i++ {
		var columnsPizza []byte
		for j := range pizza {
			columnsPizza = append(columnsPizza, pizza[j][i])
		}
		newPizza = append(newPizza, string(columnsPizza))
	}
	return newPizza
}

// calculate slices (pizza must be in line)
func calculateSlices(pizza []string, row, col, minIngredient, maxArea int, byRow bool) [][]int {
	var results [][]int

	for i := range pizza {
		beg := 0
		end := 0
		mushroom := 0
		tomato := 0

		var last int
		if byRow {
			last = col
		} else {
			last = row
		}

		for end < last { // for
			if pizza[i][end] == 77 {
				mushroom++
			} else if pizza[i][end] == 84 {
				tomato++
			}
			end++

			if end-beg > maxArea { // if slide too big, remove one ingredient
				if pizza[i][beg] == 77 {
					mushroom--
				} else if pizza[i][beg] == 84 {
					tomato--
				}
				beg++
			}

			if end-beg <= maxArea && mushroom >= minIngredient && tomato >= minIngredient {
				if byRow {
					results = append(results, []int{i, beg, i, end - 1})
				} else {
					results = append(results, []int{beg, i, end - 1, i})
				}
				beg = end
				tomato = 0
				mushroom = 0
			}
		}
	}

	return results
}

// print and write results
func printResult(results [][]int) {
	// create file
	file, err := os.Create(filename + ".out")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	delim := " "
	// print score
	output := len(results)
	fmt.Fprintln(file, output)
	fmt.Println(output)
	// print coordinates
	for i := range results {
		output := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(results[i])), delim), "[]")
		// print to file
		fmt.Fprintln(file, output)
		// print to terminal
		fmt.Println(output)
	}

}

func main() {
	pizza, row, col, minIngredient, maxArea = makePizza(filename)
	colPizza := createColumns(pizza, col)

	// rowResults := calculateSlices(pizza, row, col, minIngredient, maxArea, true)
	colResults := calculateSlices(colPizza, row, col, minIngredient, maxArea, false)

	printResult(colResults)
}
