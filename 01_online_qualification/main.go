package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filename = "b_lovely_landscapes"
var photos = make(map[int][]string)

// var slideshow [][]int
var slideshow []int

func readFile(path string) (int, map[int][]string) {

	// read file
	file, err := os.Open(path + ".txt") // read file
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close() // close file

	bf := bufio.NewScanner(file)

	bf.Scan() // read 1st line file
	numPhoto, _ := strconv.Atoi(bf.Text())

	indexPhoto := 0
	for bf.Scan() {
		photos[indexPhoto] = strings.Split(bf.Text(), " ") // split photo
		indexPhoto++
	}

	return numPhoto, photos
}

func printResult(slideshow []int) {
	// create file
	file, err := os.Create(filename + ".out")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	output := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slideshow)), "\n"), "[]")
	fmt.Fprintln(file, output)
}

func main() {
	fmt.Println("Google Hashcode 2019")
	_, photos := readFile(filename)

	for i := 0; i < len(photos); i++ {
		for tags := 2; tags < len(photos[i]); tags++ { // tags := 2 as tags starts as index too. The number of tags is defined by there index
			for j := 0; j < len(photos); j++ {
				for nextPhotoTag := 2; nextPhotoTag < len(photos[j]); nextPhotoTag++ {
					_, currExist := photos[i]
					_, nextExist := photos[j]
					if currExist && nextExist && photos[i][tags] == photos[j][nextPhotoTag] && i != j {
						slideshow = append(slideshow, i)
						slideshow = append(slideshow, j)
						delete(photos, i)
						delete(photos, j)
						if i > 5000 {
							fmt.Println("Saved.")
							printResult(slideshow)
						}
						break
					}
				}
			}
		}
	}
}
