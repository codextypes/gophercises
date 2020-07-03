package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// read in file data
	dat, err := ioutil.ReadFile("problems.csv")
	check(err)

	// convert data to csv records
	reader := csv.NewReader(strings.NewReader(string(dat)))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// get ready to read from stdin
	userInput := bufio.NewReader(os.Stdin)
	points := 0
	// go over each record
	for i := 0; i < len(records); i++ {
		// left item is the question
		fmt.Print(records[i][0])
		fmt.Print("\t")

		// await user input
		text, _ := userInput.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		// right item is the answer
		if text == records[i][1] {
			points++
			fmt.Printf("Correct! You now have %d points\n", points)
		} else {
			fmt.Println("Nope, sorry.")
		}
		fmt.Println()
	}
	fmt.Printf("Game Over | Total Score: %d\n", points)

}
