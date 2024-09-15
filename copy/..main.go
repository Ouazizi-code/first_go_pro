package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: textmod <input file> <output file>")
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := processLine(line)
		writer.WriteString(modifiedLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}

	if err := writer.Flush(); err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
	}
}

func processLine(line string) string {
	line = applyModifiers(line)
	line = handlePunctuation(line)
	line = handleQuotes(line)
	line = handleA(line)
	return line
}

func applyModifiers(line string) string {
	// Regex to find modifiers with optional counts
	modifierRegex := regexp.MustCompile(`(\S+)\s+\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)`)
	for {
		matches := modifierRegex.FindStringSubmatchIndex(line)
		if matches == nil {
			break
		}

		// Extract the components
		precedingText := line[:matches[0]]
		modifierWord := line[matches[2]:matches[3]]
		modifier := line[matches[4]:matches[5]]
		countStr := ""
		if matches[6] != -1 {
			countStr = line[matches[6]:matches[7]]
		}
		count := 1
		if countStr != "" {
			var err error
			count, err = strconv.Atoi(countStr)
			if err != nil {
				count = 1
			}
		}

		// Find preceding words
		precedingWords := strings.Fields(precedingText)
		if len(precedingWords) == 0 {
			line = line[matches[1]:]
			continue
		}
		
		if len(precedingWords) < count {
			count = len(precedingWords)
		}

		// Apply transformations
		switch modifier {
		case "hex":
			if len(precedingWords) > 0 {
				precedingWords[len(precedingWords)-1] = applyHex(modifierWord)
			}
		case "bin":
			if len(precedingWords) > 0 {
				precedingWords[len(precedingWords)-1] = applyBin(modifierWord)
			}
		case "up":
			precedingWords = transformWords(precedingWords, toUppercase, count)
		case "low":
			precedingWords = transformWords(precedingWords, toLowercase, count)
		case "cap":
			precedingWords = transformWords(precedingWords, capitalize, count)
		}

		// Construct the new line with transformed text
		newLine := strings.Join(precedingWords, " ") + line[matches[1]:]
		line = newLine
	}

	return line
}

func applyHex(word string) string {
	if val, err := strconv.ParseInt(word, 16, 0); err == nil {
		return fmt.Sprintf("%d", val)
	}
	return word
}

func applyBin(word string) string {
	if val, err := strconv.ParseInt(word, 2, 0); err == nil {
		return fmt.Sprintf("%d", val)
	}
	return word
}

func transformWords(words []string, transform func(string) string, count int) []string {
	if count > len(words) {
		count = len(words)
	}
	for i := 0; i < count; i++ {
		words[i] = transform(words[i])
	}
	return words
}

func toUppercase(word string) string {
	return strings.ToUpper(word)
}

func toLowercase(word string) string {
	return strings.ToLower(word)
}

func capitalize(word string) string {
	return strings.Title(word)
}

func handlePunctuation(line string) string {
	// Remove space before punctuation and handle special cases
	re := regexp.MustCompile(`(\w)\s*([.,!?;:])\s*`)
	line = re.ReplaceAllString(line, "$1$2 ")

	// Handle cases with multiple punctuation
	re = regexp.MustCompile(`([.,!?;:])\s+([.,!?;:])`)
	line = re.ReplaceAllString(line, "$1$2")

	// Remove spaces before single punctuation
	re = regexp.MustCompile(`\s+([.,!?;:])`)
	line = re.ReplaceAllString(line, "$1")

	// Handle cases with ellipses and exclamations
	re = regexp.MustCompile(`\s+(\.\.\.|\!\?)`)
	line = re.ReplaceAllString(line, "$1")

	return line
}

func handleQuotes(line string) string {
	// Trim spaces around single quotes
	re := regexp.MustCompile(`'\s+([^']+?)\s+'`)
	line = re.ReplaceAllString(line, "'$1'")

	return line
}

func handleA(line string) string {
	// Change 'a' to 'an' before words starting with vowel or 'h'
	re := regexp.MustCompile(`\b(a)\s+([aeiouhH]\w*)`)
	return re.ReplaceAllString(line, "an $2")
}
