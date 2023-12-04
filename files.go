package main

import "os"

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
