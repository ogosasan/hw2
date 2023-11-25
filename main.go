package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type LineInfo struct {
	Count int
	Line  string
}

type OriginalLine struct {
	Line         string
	OriginalLine string
}

func processInput(scanner *bufio.Scanner, iPtr *bool, fPtr, sPtr *int) ([]LineInfo, []OriginalLine) {
	var lines []LineInfo
	var originalLines []OriginalLine

	for scanner.Scan() {
		line := scanner.Text()
		originalLine := line

		if *iPtr {
			line = strings.ToLower(line)
		}

		fields := strings.Fields(line)
		if *fPtr > 0 {
			if *fPtr < len(fields) {
				line = strings.Join(fields[*fPtr:], " ")
			} else {
				line = strings.Join(fields[:], " ")
			}
		}
		if *sPtr > 0 {
			if *sPtr < len(line) {
				line = line[*sPtr:]
			} else {
				line = ""
			}
		}

		found := false
		for i, l := range lines {
			if l.Line == line {
				lines[i].Count++
				found = true
				break
			}
		}
		if !found {
			lines = append(lines, LineInfo{Count: 1, Line: line})
			originalLines = append(originalLines, OriginalLine{Line: line, OriginalLine: originalLine})
		}
	}
	return lines, originalLines
}

func outProcess(output *os.File, cPtr, uPtr, dPtr *bool, lines []LineInfo, originalLines []OriginalLine) {
	for i, line := range lines {
		switch {
		case *cPtr:
			if output == os.Stdout {
				fmt.Println(line.Count, originalLines[i].OriginalLine)
			} else {
				_, err := output.WriteString(fmt.Sprintf("%d ", line.Count) + originalLines[i].OriginalLine + "\n")
				if err != nil {
					return
				}
			}
		case *dPtr:
			if line.Count > 1 {
				if output == os.Stdout {
					fmt.Println(originalLines[i].OriginalLine)
				} else {
					_, err := output.WriteString(originalLines[i].OriginalLine)
					if err != nil {
						return
					}
				}
			}
		case *uPtr:
			if line.Count == 1 {
				if output == os.Stdout {
					fmt.Println(originalLines[i].OriginalLine)
				} else {
					_, err := output.WriteString(originalLines[i].OriginalLine + "\n")
					if err != nil {
						return
					}
				}
			}
		default:
			if output == os.Stdout {
				fmt.Println(originalLines[i].OriginalLine)
			} else {
				_, err := output.WriteString(originalLines[i].OriginalLine + "\n")
				if err != nil {
					return
				}
			}
		}
	}
}

func processInputArguments() (*bool, *bool, *bool, *bool, *int, *int) {
	cPtr := flag.Bool("c", false, "a bool")
	dPtr := flag.Bool("d", false, "a bool")
	uPtr := flag.Bool("u", false, "a bool")
	iPtr := flag.Bool("i", false, "a bool")
	fPtr := flag.Int("f", 0, "an int")
	sPtr := flag.Int("s", 0, "an int")

	flag.Parse()

	return cPtr, dPtr, uPtr, iPtr, fPtr, sPtr
}

func openInputFile(fileName string) (*os.File, error) {
	if fileName != "" {
		inputFile, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		return inputFile, nil
	}
	return os.Stdin, nil
}

func createOutputFile(fileName string) (*os.File, error) {
	if fileName != "" {
		outputFile, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		return outputFile, nil
	}
	return os.Stdout, nil
}

func main() {
	cPtr, dPtr, uPtr, iPtr, fPtr, sPtr := processInputArguments()

	if *cPtr && *dPtr || *cPtr && *uPtr || *uPtr && *dPtr {
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
	lineCount, originalLines := processInput(scanner, iPtr, fPtr, sPtr)

	outProcess(output, cPtr, uPtr, dPtr, lineCount, originalLines)
}
