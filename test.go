package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
)

func main() {

	inputFile, inputError := os.Open("./data/GeoLite2-City-Locations-zh-CN.csv")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
		    "Does the file exist?\n" +
		    "Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
		    return
		}
	}
}
