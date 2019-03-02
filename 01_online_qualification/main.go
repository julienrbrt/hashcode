package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var filename = "b_lovely_landscapes"
var photos = make(map[int][]string)

var slideshow []int

//readFile and source file
func readFile(path string) (int, map[int][]string) {

	// read file
	source, err := os.Open(path + ".txt") // read file
	if err != nil {
		fmt.Print(err)
	}
	defer source.Close() // close file

	// import source file
	bfSource := bufio.NewScanner(source)

	bfSource.Scan() // read 1st line file
	numPhoto, _ := strconv.Atoi(bfSource.Text())

	indexPhoto := 0
	for bfSource.Scan() {
		photos[indexPhoto] = strings.Split(bfSource.Text(), " ") // split photo
		indexPhoto++
	}

	// read file
	output, err := os.Open(path + ".out") // read file
	if err == nil {
		// start from saved results -- usefull only when bruteforce
		bfOutput := bufio.NewScanner(output)
		for bfOutput.Scan() {
			i, _ := strconv.Atoi(bfOutput.Text())
			_, currExist := photos[i]

			if currExist {
				fmt.Println(i, " removed.")
				delete(photos, i)
			}
		}
	}
	defer output.Close() // close file
	return numPhoto, photos
}

func printResult(slideshow []int) {
	// create file
	file, err := os.OpenFile(filename+".out", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		file, err := os.Create(filename + ".out")
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	// for i := range slideshow {
	// 	if photos[i][0] == "V" {

	// 	}
	// }

	output := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slideshow)), "\n"), "[]")
	fmt.Fprintln(file, output)
}

func main() {
	fmt.Println("Google Hashcode 2019")
	_, photos := readFile(filename)
	defer printResult(slideshow)

	// print result if exited
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		printResult(slideshow)
		os.Exit(1)
	}()

	for i := range photos {
		for tags := 2; tags < len(photos[i]); tags++ { // tags := 2 as tags starts as index too. The number of tags is defined by there index
			for j := range photos {
				for nextPhotoTag := 2; nextPhotoTag < len(photos[j]); nextPhotoTag++ {
					_, currExist := photos[i]
					_, nextExist := photos[j]
					if currExist && nextExist && photos[i][tags] == photos[j][nextPhotoTag] && i != j {
						slideshow = append(slideshow, i)
						slideshow = append(slideshow, j)
						delete(photos, i)
						delete(photos, j)
						// search relation with next photo
						i = j
						break
					}
				}
			}
		}
	}
	fmt.Println("No more photos availables.")
}
