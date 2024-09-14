package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	outputPointer := flag.String("output", "", "specify the output file")

	flag.Parse()

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
		return
	}
	// Extract text and banner style from command line arguments
	text := strings.Join(os.Args[2:len(os.Args)-1], " ")
	banner := os.Args[len(os.Args)-1]

	// Open the file corresponding to the specified banner style
	file, err := os.Open(fmt.Sprintf("%s.txt", banner))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Store lines from the file in a slice
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Map to store ASCII characters for each letter
	asciiChrs := make(map[int][]string)
	dec := 31

	// Populate the map with ASCII characters from the file
	for _, line := range lines {
		if line == "" {
			dec++
		} else {
			asciiChrs[dec] = append(asciiChrs[dec], line)
		}
	}

	// Get the result by writing the text

	result := WriteText(text, asciiChrs)

	// Check if the --output flag is provided
	if *outputPointer != "" {
		// Write result to the specified output file
		outputFile, err := os.Create(*outputPointer)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()

		// Write the result to the output file
		fmt.Fprint(outputFile, result)
		fmt.Println("Result written to", *outputPointer)
	} else {
		// Print the result to the console if --output flag is not provided
		fmt.Print(result)
	}
}

// NewLineScanner creates a scanner that splits by lines
func NewLineScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

// Function to write the text
func WriteText(n string, y map[int][]string) string {
	var result strings.Builder
	//for j := 0; j < len(y[32]); j++ {
	n = strings.ReplaceAll(n, "\\n\\n", "\\n")
	if strings.Contains(n, "\\n") {
		parts := strings.Split(n, "\\n")
		for _, part := range parts {
			for j := 0; j < len(y[32]); j++ {
				for _, letter := range part {
					result.WriteString(y[int(letter)][j])
				}

				result.WriteByte('\n')
			}
		}
	} else {
		for j := 0; j < len(y[32]); j++ {
			for _, letter := range n {
				// Append each character's ASCII representation to the result
				result.WriteString(y[int(letter)][j])
			}
			// Add a newline character to separate lines
			result.WriteByte('\n')
		}
	}
	//}
	// Return the final result as a string
	return result.String()
}
