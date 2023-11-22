package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	cPtr := flag.Bool("c", false, "a bool")
	dPtr := flag.Bool("d", false, "a bool")
	uPtr := flag.Bool("u", false, "a bool")
	iPtr := flag.Bool("i", false, "a bool")
	fPtr := flag.Int("f", 0, "an int")
	sPtr := flag.Int("s", 0, "an int")

	flag.Parse()
	if *cPtr && *dPtr || *cPtr && *uPtr || *uPtr && *dPtr {
		fmt.Println("Error flags")
		os.Exit(1)
	}

	var input *os.File
	var output *os.File

	if len(flag.Args()) > 0 {
		inputFile, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer inputFile.Close()
		input = inputFile
	} else {
		input = os.Stdin
	}
	if len(flag.Args()) > 1 {
		outputFile, err := os.Create(flag.Arg(1))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outputFile.Close()
		output = outputFile
	} else {
		output = os.Stdout
	}

	lineCount := make(map[string]int)
	originalLines := make(map[string]string)
	scanner := bufio.NewScanner(input)
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
				line = ""
			}

		}
		if *sPtr > 0 {
			if *sPtr < len(line) {
				line = line[*sPtr:]
			} else {
				line = ""
			}
		}
		lineCount[line]++
		if _, found := originalLines[line]; !found {
			originalLines[line] = originalLine
		}
	}
	for line, count := range lineCount {
		switch {
		case *cPtr:
			if output == os.Stdout {
				fmt.Println(count, originalLines[line])
			} else {
				output.WriteString(fmt.Sprintf("%d ", count) + originalLines[line] + "\n")
			}
		case *dPtr:
			if count > 1 {
				if output == os.Stdout {
					fmt.Println(originalLines[line])
				} else {
					output.WriteString(originalLines[line] + "\n")
				}
			}
		case *uPtr:
			if count == 1 {
				if output == os.Stdout {
					fmt.Println(originalLines[line])
				} else {
					output.WriteString(originalLines[line] + "\n")
				}
			}
		default:
			if output == os.Stdout {
				fmt.Println(originalLines[line])
			} else {
				output.WriteString(originalLines[line] + "\n")
			}

		}
	}
}
