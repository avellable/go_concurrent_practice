package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var initialString string
var finalString string

var stringLength int

// Gets a letter from channel and puts into fila string
func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup) {
	// get a letter from channel
	letter := <-letterChannel
	// add it to the finalString
	finalString += letter
	// signal done to waitgroup
	wg.Done()
}

// Capitalize a letter and add to channel
func capitalize(letterChannel chan string, currentLetter string, wg *sync.WaitGroup) {

	// capitalize the letter
	thisLetter := strings.ToUpper(currentLetter)
	// signal done to waitgroup
	wg.Done()
	// insert the capitalized letter to channel
	letterChannel <- thisLetter
}

func main() {

	/*
	   This function simply describes the way an application flow for
	   channel based go application works.
	   The flow:
	   1) Start Application
	   2) Start Channel
	   3) Capitalize letter and put it to channel
	   4) Take out capitalized letter from channel and put it to finaString
	*/
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	initialString = "Four score and seven years ago our fathers brought forth on this continent, a new nation, conceived in Liberty, and dedicated to the proposition that all men are created equal."
	initialBytes := []byte(initialString)

	var letterChannel chan string = make(chan string)

	stringLength = len(initialBytes)

	for i := 0; i < stringLength; i++ {
		wg.Add(2)

		go capitalize(letterChannel, string(initialBytes[i]), &wg)
		go addToFinalStack(letterChannel, &wg)

		wg.Wait()
	}

	fmt.Println(finalString)

}
