package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Append_Text appends text into result.txt
func Append_Text(text string) {
	resultFile := os.Args[2]
	file, err := os.OpenFile(resultFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	text = strings.TrimSpace(text)
	_, err = file.WriteString(text + "\n")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: not enough arguments")
		fmt.Println("Usage: go run . input.txt result.txt")
		return
	}

	inputFileName, outputFileName := os.Args[1], os.Args[2]
	if inputFileName != "semple.txt" || outputFileName != "result.txt" {
		fmt.Println("Error: invalid file naming")
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}

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
	lines := []string{}
	scanText := bufio.NewScanner(inputFile)
	for scanText.Scan() {
		line := scanText.Text()
		modifiedLine := DistributeLines(line)
		lines = append(lines, modifiedLine)
	}

	// Write lines to the output file
	for i, line := range lines {
		if i == len(lines)-1 {
			outputFile.WriteString(line)
		} else {
	
			outputFile.WriteString(line + "\n")
		}
	}
}

func DistributeLines(line string) string {
	line = VowelsManipulation(line)
	line = handlePunctuation(line)
	line = handleQuotes(line)
	line = SentenceManipulation(line)
	line = handlePunctuation(line)
	line = handleQuotes(line)
	line = Remove_Flags(line)

	return line
}

// SentenceManipulation applies transformations based on flags
func SentenceManipulation(line string) string {
	flags := regexp.MustCompile(`^(.+?)\s*\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+))?\s*\)`)
	for {
		matches := flags.FindStringSubmatchIndex(line)
		if matches == nil {
			break
		}

		sentence := line[matches[2]:matches[3]]
		sentence = strings.TrimSpace(sentence)

		flag := line[matches[4]:matches[5]]

		countStr := ""
		if matches[6] != -1 {
			countStr = line[matches[6]:matches[7]]
		}
		count := 1
		if countStr != "" {
			var err error
			count, err = strconv.Atoi(countStr)
			if err != nil {
				fmt.Println("Error: invalid number inside flag:", countStr)
				count = 0
			}
		}

		transformedText := ""
		switch flag {
		case "hex":
			words := strings.Fields(sentence)
			toSend := sentence
			if len(words) != 0 {
				toSend = words[len(words)-1]
			}

			if val, err := strconv.ParseInt(toSend, 16, 0); err == nil {
				numAsString := strconv.Itoa(int(val))
				words[len(words)-1] = numAsString
				transformedText = strings.Join(words, " ") + " "
			} else {
				fmt.Println("Error: invalid syntax for hex transformation:", "'", toSend, "'")
				transformedText = sentence
			}
		case "bin":
			words := strings.Fields(sentence)
			toSend := sentence
			if len(words) != 0 {
				toSend = words[len(words)-1]
			}

			if val, err := strconv.ParseInt(toSend, 2, 0); err == nil {
				numAsString := strconv.Itoa(int(val))
				words[len(words)-1] = numAsString
				transformedText = strings.Join(words, " ") + " "
			} else {
				fmt.Println("Error: invalid syntax for bin transformation:", "'", toSend, "'")
				transformedText = sentence
			}
		case "up":
			transformedText = transformWords(sentence, toUppercase, count)
		case "low":
			transformedText = transformWords(sentence, toLowercase, count)
		case "cap":
			transformedText = transformWords(sentence, capitalize, count)
		}

		line = line[:matches[0]] + transformedText + line[matches[1]:]
	}

	return line
}

// Remove_Flags removes flags from the line
func Remove_Flags(line string) string {
	re := regexp.MustCompile(`(?m)^\s*\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)\s*$`)
	line = re.ReplaceAllString(line, "")

	line = regexp.MustCompile(`(?m)^\s*\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)\s*`).ReplaceAllString(line, "")

	return line
}

func transformWords(text string, transform func(string) string, count int) string {
	words := strings.Fields(text)
	if count > len(words) {
		count = len(words)
	}
	n := len(words) - 1

	for i := n; i > n-count; i-- {
		words[i] = transform(words[i])
	}
	return strings.Join(words, " ")
}

func toUppercase(word string) string {
	return strings.ToUpper(word)
}

func toLowercase(word string) string {
	return strings.ToLower(word)
}

func capitalize(word string) string {
	word = strings.ToLower(word)
	return strings.Title(word)
}

func handlePunctuation(line string) string {
	re := regexp.MustCompile(`(\w)\s*([.,!?;:])\s*`)
	line = re.ReplaceAllString(line, "$1$2 ")

	re = regexp.MustCompile(`([.,!?;:])\s+([.,!?;:])`)
	line = re.ReplaceAllString(line, "$1$2")

	re = regexp.MustCompile(`\s+([.,!?;:])`)
	line = re.ReplaceAllString(line, "$1")

	re = regexp.MustCompile(`\s+(\.\.\.|\!\?)`)
	line = re.ReplaceAllString(line, "$1")

	return line
}

func handleQuotes(line string) string {
	statut := false
	var result []rune
	slice := []rune(line)
	for i := 0; i < len(slice); i++ {
		if slice[i] == '\'' && (i+1 < len(slice) && i-1 >= 0) && (unicode.IsLetter(slice[i+1]) && unicode.IsLetter(slice[i-1])) {
			result = append(result, rune(slice[i]))
			i++
		}
		if slice[i] == '\'' {
			if !statut {
				statut = true
				if i-1 >= 0 && slice[i-1] != ' ' && len(result) > 0 {
					result = append(result, ' ')
				}
				result = append(result, rune(slice[i]))
				if i+1 < len(slice) && slice[i+1] == ' ' {
					i++
				}
			} else {
				statut = false
				if len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, rune(slice[i]))
				if i+1 < len(slice) && slice[i+1] != ' ' {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, rune(slice[i]))
		}
	}
	return string(result)
}

func VowelsManipulation(line string) string {
	words := strings.Split(line, " ")
	vowels := "aeiouhAEIOUH"
	for i, word := range words {
		trimmedWord := strings.TrimSpace(word)
		if i+1 < len(words) && trimmedWord == "a" {
			nextWord := strings.TrimSpace(words[i+1])
			if len(nextWord) != 0 && strings.ContainsRune(vowels, rune(nextWord[0])) {
				words[i] = "an"
			}
		}
		if trimmedWord == "A" && i+1 < len(words) {
			nextWord := strings.TrimSpace(words[i+1])
			if len(nextWord) != 0 && strings.ContainsRune(vowels, rune(nextWord[0])) {
				words[i] = "An"
			}
		}

		if i+1 < len(words) && trimmedWord == "~a" {
			nextWord := strings.TrimSpace(words[i+1])
			if len(nextWord) != 0 && strings.ContainsRune(vowels, rune(nextWord[0])) {
				words[i] = "~an"
			}

		}
	}

	return strings.Join(words, " ")
}
