package main
import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

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
func processInput(scanner *bufio.Scanner, iPtr *bool, fPtr, sPtr *int) (map[string]int, map[string]string) {
	lineCount := make(map[string]int)
	originalLines := make(map[string]string)

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
	return lineCount, originalLines
}

func outProcess(output *os.File, cPtr, uPtr, dPtr *bool, lineCount map[string]int, originalLines map[string]string) {
	for line, count := range lineCount {
		switch {
		case *cPtr:
			if output == os.Stdout {
				fmt.Println(count, originalLines[line])
			} else {
				_, err := output.WriteString(fmt.Sprintf("%d ", count) + originalLines[line] + "\n")
				if err != nil {
					return
				}
			}
		case *dPtr:
			if count > 1 {
				if output == os.Stdout {
					fmt.Println(originalLines[line])
				} else {
					_, err := output.WriteString(originalLines[line])
					if err != nil {
						return
					}}}
		case *uPtr:
			if count == 1 {
				if output == os.Stdout {
					fmt.Println(originalLines[line])
				} else {
					_, err := output.WriteString(originalLines[line] + "\n")
					if err != nil {
						return
					}}}
		default:
			if output == os.Stdout {
				fmt.Println(originalLines[line])
			} else {
				_, err := output.WriteString(originalLines[line] + "\n")
				if err != nil {
					return
				}}}}

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
