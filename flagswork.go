package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processInput(scanner *bufio.Scanner, flags Flags) ([]LineInfo, []OriginalLine) {
	var lines []LineInfo
	var originalLines []OriginalLine

	for scanner.Scan() {
		line := scanner.Text()
		originalLine := line

		if flags.i {
			line = strings.ToLower(line)
		}

		fields := strings.Fields(line)
		if flags.f > 0 {
			if flags.f < len(fields) {
				line = strings.Join(fields[flags.f:], " ")
			} else {
				line = strings.Join(fields[:], " ")
			}
		}
		if flags.s > 0 {
			if flags.s < len(line) {
				line = line[flags.s:]
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

func outProcess(output *os.File, flags Flags, lines []LineInfo, originalLines []OriginalLine) {
	for i, line := range lines {
		switch {
		case flags.c:
			if output == os.Stdout {
				fmt.Println(line.Count, originalLines[i].OriginalLine)
			} else {
				_, err := output.WriteString(fmt.Sprintf("%d ", line.Count) + originalLines[i].OriginalLine + "\n")
				if err != nil {
					return
				}
			}
		case flags.d:
			if line.Count > 1 {
				if output == os.Stdout {
					fmt.Println(originalLines[i].OriginalLine)
				} else {
					_, err := output.WriteString(originalLines[i].OriginalLine + "\n")
					if err != nil {
						return
					}
				}
			}
		case flags.u:
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
