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

// Gets a letter from channel and puts into final string
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

	initialString = "Today the most civilized countries of the world spend a maximum of their income on war and a minimum on education. The twenty-first century will reverse this order. It will be more glorious to fight against ignorance than to die on the field of battle. The discovery of a new scientific truth will be more important than the squabbles of diplomats. Even the newspapers of our own day are beginning to treat scientific discoveries and the creation of fresh philosophical concepts as news. The newspapers of the twenty-first century will give a mere 'stick' in the back pages to accounts of crime or political controversies, but will headline on the front pages the proclamation of a new scientific hypothesis.Progress along such lines will be impossible while nations persist in the savage practice of killing each other off. I inherited from my father, an erudite man who labored hard for peace, an ineradicable hatred of war."
	initialBytes := []byte(initialString)

	// 2)
	var letterChannel chan string = make(chan string)

	stringLength = len(initialBytes)

	for i := 0; i < stringLength; i++ {
		wg.Add(2)

		// 3)
		go capitalize(letterChannel, string(initialBytes[i]), &wg)
		// 4)
		go addToFinalStack(letterChannel, &wg)

		wg.Wait()
	}

	fmt.Println(finalString)

}
