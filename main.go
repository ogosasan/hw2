package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type LineInfo struct {
	Count int
	Line  string
}

type OriginalLine struct {
	Line         string
	OriginalLine string
}

type Flags struct {
	c bool
	d bool
	u bool
	i bool
	f int
	s int
}

func main() {
	flags := processInputArguments()
	if flags.c && flags.d || flags.c && flags.u || flags.u && flags.d {
		fmt.Println("Error flags")
		os.Exit(1)
	}
	input, err := openInputFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	output, err := createOutputFile(flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	lineCount, originalLines := processInput(scanner, flags)

	outProcess(output, flags, lineCount, originalLines)
}
