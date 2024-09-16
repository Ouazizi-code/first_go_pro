package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: textmod <input file> <output file>")
		os.Exit(1)
	}

	// extract the files names
	inputFileName, outputFileName := os.Args[1], os.Args[2]

	// open the file
	inputFile, err := os.Open(inputFileName)
	// handle erors
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// if the file not exist create it
	outputFile, err := os.Create(outputFileName)
	// handle erors
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// use bufio to scan our text
	scan_text := bufio.NewScanner(inputFile)
	write_text := bufio.NewWriter(outputFile)
    test := ""
	for scan_text.Scan() {
		line := scan_text.Text()
		modifiedline := processLine(line)
		modifiedline = strings.TrimSpace(modifiedline)
		test += modifiedline + "\n"
		write_text.WriteString(modifiedline + "\n")
	}
	test = strings.TrimSpace(test)
	fmt.Println(test,"52")
	// handle erore for sccanner
	if err := scan_text.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}

	if err := write_text.Flush(); err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
	}
}

func processLine(line string) string {
	line = Vowles_manioulation(line)
	line = handlePunctuation(line)
	test , _ := Single_Quote(line)
	line = test
	line = applyModifiers(line)
	line = handlePunctuation(line)
	
	return line
}

func applyModifiers(line string) string {
	// Regex to find modifiers with optional counts
	//(\S+)\s+((?:\S+\s+)*)(\((hex|bin|up|low|cap)(?:,\s*(\d+))?\))
	// (\S+)\s+\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)
	//^(.+?)\s+\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)
	// ^(.+?)\s*\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+))?\s*\)
	// ^(.+?)\s*\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+))?\s*\)
	// ^(.+?)\s*\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+))?\s*\)$
	//(?i)\b(\S*)\s*\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)
	// Remove patterns with no preceding text
	line = removeEmptyModifiers(line)
	
	modifierRegex := regexp.MustCompile(`(?i)\b(\S*)\s*\(\s+(hex|bin|up|low|cap)(?:,\s*(\d+))?\s+\)`)
	for {
		matches := modifierRegex.FindStringSubmatchIndex(line)
		//fmt.Println(matches,"87")
		if matches == nil {
			
			break
		}

		
		
		word := line[matches[2]:matches[3]]
		word = strings.TrimSpace(word)
		//fmt.Println(word,"78")
		
		modifier := line[matches[4]:matches[5]]
		//fmt.Println(modifier,"80")
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
		//fmt.Println(count,"93")
		transformedText := ""
		switch modifier {
		case "hex":
			words := strings.Fields(word)
			to_send := words[len(words)-1]
			if val, err := strconv.ParseInt(to_send, 16, 0); err == nil {
				num_as_string := strconv.Itoa(int(val))
				words[len(words)-1] = num_as_string
				valid := strings.Join(words, " ")
				transformedText = valid+" "
			} else {
				fmt.Println("invalid syntacs to parsint in hex transformation : ", "'", to_send, "'")
				transformedText = word
			}
		case "bin":
			words := strings.Fields(word)
			to_send := words[len(words)-1]

			if val, err := strconv.ParseInt(to_send, 2, 0); err == nil {
				num_as_string := strconv.Itoa(int(val))
				words[len(words)-1] = num_as_string
				valid := strings.Join(words, " ")
				transformedText = valid+" "
			} else {
				fmt.Println("invalid syntacs to parsint in bin transformation : ", "'", to_send, "'")
				transformedText = word
			}
		case "up":
			transformedText = transformWords(word, toUppercase, count)
		case "low":
			transformedText = transformWords(word, toLowercase, count)
		case "cap":
			transformedText = transformWords(word, capitalize, count)
		/*default:
			transformedText = word*/
		}

		line = line[:matches[0]] + transformedText + line[matches[1]:]
	}
	

	return line
}

// Helper function to remove patterns with no preceding text
func removeEmptyModifiers(line string) string {
		// Regex to remove (modifier) if it's the only content or if preceded only by whitespace/newlines
	re := regexp.MustCompile(`(?m)^\s*\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)\s*$`)
	line = re.ReplaceAllString(line, "")
	
	// Also handle patterns with preceding newlines or whitespace but no actual text
	line = regexp.MustCompile(`(?m)^\s*\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)\s*`).ReplaceAllString(line, "")
	
	return line
}

func transformWords(text string, transform func(string) string, count int) string {
	fmt.Println(text,"125")
	//fmt.Println(count,126)
	words := strings.Fields(text)
	if count > len(words) {
		count = len(words)
	}
	n := len(words) - 1
	for i := n; i > n-count; i-- {
		words[i] = transform(words[i])
	}
	return strings.Join(words, " ") + " "
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

func Vowles_manioulation(line string) string {
	result := ""
	// this function Correct Indefinite article
	words := strings.Split(line, " ")
	vowels := "aeiouhAEIOUH"
	for i, word := range words {
		trimmedWord := strings.TrimSpace(word)
		if trimmedWord == "a" && i+1 < len(words) {
			nextWord := strings.TrimSpace(words[i+1])
			if strings.ContainsRune(vowels, rune(nextWord[0])) {
				words[i] = "an"
			}
		}
		if trimmedWord == "A" && i+1 < len(words) {
			nextWord := strings.TrimSpace(words[i+1])
			if strings.ContainsRune(vowels, rune(nextWord[0])) {
				words[i] = "An"
			}
		}
	}
	result = strings.Join(words, " ")
	return result
}

func ReplaceInDescendingOrder(s, old, new string) string {
	// Start with the original string
	result := s
	// Find all occurrences of old substring
	for {
		// Find the last occurrence of the substring to replace
		index := strings.LastIndex(result, old)
		if index == -1 {
			break
		}
		// Replace the found substring with the new substring
		result = result[:index] + new + result[index+len(old):]
	}
	return result
}

func IS_Punctuation(char string) bool {
	puncts := ".,!?;:"

	return strings.Contains(puncts, (char))
}

// this function it really manipulate punctuations
func Punctuations(text string) string {
	// Remove spaces before punctuation marks

	// Define a punctuation mark
	puncts := regexp.MustCompile(`\s([.,;:!?])\s`)

	// Replace  punctuation mark with a single instance
	text = puncts.ReplaceAllString(text, "$1")

	// Handle spacing around punctuation marks
	for i := range text {
		if IS_Punctuation(string(text[i]))  && i < len(text)-1 && (unicode.IsLetter(rune(text[i+1])) || unicode.IsDigit(rune(text[i+1]))) {
			text = text[:i+1] + " " + text[i+1:]
		}
	}

	// this just to smplify the code
	to_replace := strings.NewReplacer(
		" .", ".",
		" ,", ",",
		" ;", ";",
		" :", ":",
		" !", "!",
		" ?", "?",
		"~", "~ ",
		//"  ~.", "~ .",
	)

	text = to_replace.Replace(text)

	return text
}

func between(before rune, after rune) bool {
	return isALphaNumeric(before) && isALphaNumeric(after)
}

func isPair(number int) bool {
	return number%2 == 0
}

func Single_Quote(sentence string) (string, error) {
	words := []rune(sentence)
	result := ""
	is_open := false
	var start, end int
	var word string
	count_quotes := 0
	is_closed := true
	for i := 0; i < len(words); i++ {
		if words[i] == '\'' && i+1 < len(words) && i-1 >= 0 && between(words[i-1], words[i+1]) {
			continue
		} else {
			if words[i] == '\'' && !is_open {
				if i-1 >= 0 {
					result += " "
				}

				is_open = true
				start = i
				is_closed = false
				count_quotes++
			} else if words[i] == '\'' && is_open {
				end = i
				word = string(words[start+1 : end])
				is_open = false

				result += "'" + strings.TrimSpace(word) + "'" + " "

				word = ""
				is_closed = true
				count_quotes++
			} else if is_closed {
				result += string(words[i])
			}
		}
	}

	if !isPair(count_quotes) {
		return "", errors.New("there is no ending quote")
	} else {
		test := strings.Fields(result)
		return strings.Join(test, " "), nil
	}
}
func isALphaNumeric(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}
