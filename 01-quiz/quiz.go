package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gameOver(points *int) {
	pointValue := *points
	fmt.Printf("\nGame Over | Total Score: %d\n", pointValue)
	os.Exit(0)
}

func gameTimer(points *int, seconds int) *time.Timer {
	timer := time.NewTimer(time.Second * time.Duration(seconds))

	go func(points *int) {
		<-timer.C
		gameOver(points)
	}(points)

	return timer
}

func main() {

	// read timer flag
	timer := flag.Int("limit", 20, "time limit (in seconds) - default 20")
	csvFile := flag.String("csv", "problems.csv", "CSV file path - default ./problems.csv")
	fileName := *csvFile
	// read in file data
	dat, err := ioutil.ReadFile(fileName)
	check(err)

	// convert data to csv records
	reader := csv.NewReader(strings.NewReader(string(dat)))
	records, err := reader.ReadAll()
	check(err)

	userInput := bufio.NewReader(os.Stdin)
	points := 0
	timeLeft := *timer
	fmt.Printf("You have %d seconds!\n", timeLeft)

	// start the timer
	ticker := gameTimer(&points, timeLeft)
	defer ticker.Stop()

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
	gameOver(&points)
}
