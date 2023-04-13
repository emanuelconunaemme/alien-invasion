package main

import (
	"flag"
	"fmt"
	"os"
	"saga/aliens/pkg/play"
)

func main() {
	// Command line args
	var aliensNumber int
	var inputFile string
	var showHelp bool
	var verbose bool

	flag.IntVar(&aliensNumber, "aliensNumber", 2, "Number of aliens playing the game")
	flag.StringVar(&inputFile, "inputFile", "data/input.txt", "Input file with formal `Foo north=Bar west=Baz south=Qu-ux`")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.BoolVar(&verbose, "verbose", false, "Print the state at each iteration")

	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	err := play.Play(inputFile, aliensNumber, verbose)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] %s\n", err.Error())
	}
}

func printHelp() {
	fmt.Println("Usage: aliens [options]")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
}
