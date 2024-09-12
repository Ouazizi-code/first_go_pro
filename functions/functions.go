package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// this function about removing extra spaces
func Expand_Spaces(s string) string {
	// the field work with spaces or more spaces
	result := strings.Fields(s)
	valid_line := strings.TrimSpace(strings.Join(result, " "))

	return valid_line
}

// this function split line by newline
func Split_By_Newline(text string) string{
	
	test := ""
	
	
	for i := 0 ; i < len(text); i++ {
		
		char := text[i]
		test += string(char)
		if char == '\n' && i  > 1 {
			test += "~"
		}
		
		
	}
	test1 := strings.Fields(test)
	test2 := strings.Join(test1," ")
	/*fmt.Println(text)
	fmt.Println(test)
	fmt.Println(test1,len(test1))*/
	fmt.Println(test2)
	
	return test2
}

// this function split line from a )
func Split_line(s string) []string {
	result := []string{}
	startIndex := 0
	endIndex := 0
	delemeter := ')'

	for i, char := range s {
		if char == delemeter {
			endIndex = i
			result = append(result, s[startIndex:endIndex+1])
			startIndex = endIndex + 1
		}
	}
	// append the rest line in the end if exist
	result = append(result, Expand_Spaces(s[startIndex:]))

	return result
}

// this finction to search for keyword betwen braces
func Search_KeyWord(s string) (string, string, string, int) {
	flag := ""
	key_Word := ""
	Int_AsString := ""
	final_int := 0
	startIndex := 0
	end_index := 0
	// lets loop throught the stering in reverse to extract the full result
	for i := len(s) - 1; i >= 0; i-- {
		char := s[i]
		if char == '(' {
			startIndex = i
			break
		} else {
			startIndex = len(s)
		}

	}

	// this the fill line inside braces
	for_braces := s[startIndex:] // variable not edited for remove braces function
	flag = strings.Replace(s[startIndex:], " ", "", -1)
	//result = Expand_Spaces(result)
	startIndex = 0
	end_index = 0
	
	// lests extract just keyword
	for i := 0; i < len(for_braces); i++ {
		char := for_braces[i]
		if char >= 'a' && char <= 'z'  {
			startIndex = i
		}
		if char == ',' || char == ')'|| char >= '0' && char <= '9'{
			end_index = i-1
		}
	}
	key_Word = for_braces[startIndex:end_index]

	// lets extract thhe number inside the braces
	for i := 0; i < len(for_braces); i++ {
		char := for_braces[i]
		if char >= '0' && char <= '9' {
			Int_AsString += string(char)
		}
	}
	// convert the string into valid number
	final_int, _ = strconv.Atoi(Int_AsString)
	fmt.Println(for_braces,flag,key_Word,final_int)

	return for_braces, flag, key_Word, final_int
}

// this function check the valid keyword
func Is_Valid(flag, key_word string, number int) (bool, bool) {
	// convert the number to string
	num_as_string := strconv.Itoa(number)
	first_case := "(" + key_word + ")"
	second_case := "(" + key_word + "," + num_as_string + ")"

	// this bool s just for removing braces or not
	remove_braces_or_not := true

	// check now the full result if it a valid flag or not
	// all key words
	if key_word == "cap" || key_word == "up" || key_word == "low" {
		if flag == first_case || flag == second_case {
			return true, true
		}
	} else {
		remove_braces_or_not = false
	}

	// check the bin and hex alone
	if key_word == "hex" || key_word == "bin" {
		if flag == first_case {
			return true, true
		}
	} else {
		remove_braces_or_not = false
	}

	return false, remove_braces_or_not
}

// this function remove only braces and return valid string
func Rmove_braces(sentenc, delimiter string, remove_braces_or_not bool) (string, string) {
	// in this case our delimiter contains any thing inside braces
	// exemple delimiter = (cap,5) passed as params
	status := false // this condition for checking if ) exist or not
	index := 0
	for i := 0; i < len(sentenc); i++ {
		char := sentenc[i]
		if char == '(' {
			index = i
		}
	}
	for i := index; i < len(sentenc); i++ {
		char := sentenc[i]
		if char == ')' {
			status = true
		}
	}

	bin_or_hex := ""

	result := ""

	result = strings.Replace(sentenc, delimiter, "", 1)
	result = strings.TrimSpace(result)

	// so let do some work for bin and hex
	// extract the string befor braces asn well as is a binary or hex
	arr := strings.Split(result, " ")
	//fmt.Println(arr)
	// after removing braces the last index is the bin_or_hex
	bin_or_hex = arr[len(arr)-1]
	if bin_or_hex == "" {
		bin_or_hex = arr[len(arr)-1]
	}

	// now depend on remove_braces_or_not we can proced
	if status {
		if remove_braces_or_not {
			return result, bin_or_hex
		} else {
			return sentenc, bin_or_hex
		}
	} else {
		return sentenc, ""
	}
}

// this function edit the sentence depend the keyword and number
// this function contains swith cases
func Edit_Sentece(sentenc, key_word, bin_or_hex string, number int) string {
	result := ""
	// start our switch
	switch key_word {
	case "cap":
		result = Capitalize(sentenc, number)
	case "up":
		result = To_Upper(sentenc, number)
	case "low":
		result = To_Lower(sentenc, number)
	case "bin":
		// convert the bin string to dicimal
		dicimal := To_Dicimal(bin_or_hex, key_word)
		// let convert now the number to string
		num_as_string := strconv.Itoa(dicimal)
		result = Raplace_Dicimal(sentenc, bin_or_hex, num_as_string)
	case "hex":
		dicimal := To_Dicimal(bin_or_hex, key_word)
		// let convert now the number to string
		num_as_string := strconv.Itoa(dicimal)
		result = Raplace_Dicimal(sentenc, bin_or_hex, num_as_string)
	}
	return result
}

// this the sentence manipulation function
func Sentenc_Mainpulation(valid_sentence, full_result, key_word, bin_or_hex string, number int, status bool) string {
	result := ""
	// check the keyword if it  valid
	if status {
		result = Edit_Sentece(valid_sentence, key_word, bin_or_hex, number)
	} else {
		result = valid_sentence
	}
	return result
}

// this function destribute each sentence to manipulate
func Destribute_Sentences(line string) string {
	// lets splite our line into small sentences from our delimiter
	array_of_sentences := Split_line(line)
	n := len(array_of_sentences)

	result := ""

	// now lets destribute our sentences to manipulate with a for loop
	for i := 0; i < n; i++ {

		sentenc := array_of_sentences[i]
		// send this sentenc to search_keyword to find the keyword
		for_braces, flag, key_word, number := Search_KeyWord(sentenc)
		// lets check if the keyword is valid
		status, remove_braces_or_not := Is_Valid(flag, key_word, number)
		// lets modifid the sentence and remove the braces if exist
		valid_sentence, bin_or_hex := Rmove_braces(sentenc, for_braces, remove_braces_or_not)
		// lets send this valid sentence to manipulation depend on the keyword and status
		manipulated_sentenc := Sentenc_Mainpulation(result+valid_sentence, flag, key_word, bin_or_hex, number, status)
		// refresh the result and concat it with  the valid sentence
		result = ""

		result += manipulated_sentenc + " "

	}

	return result
}

// this function convert hex and binary  to dicimal

func To_Dicimal(bin_or_hex, key_Word string) int {
	var result int64
	switch key_Word {
	case "bin":
		result, err := strconv.ParseInt(bin_or_hex, 2, 64)
		// handle erors
		if err != nil {
			fmt.Println("erour", err)
			return 0
		}
		return int(result)
	case "hex":

		result, err := strconv.ParseInt(bin_or_hex, 16, 64)
		// handle erors
		if err != nil {
			fmt.Println("erour", err)
			return 0
		}
		return int(result)
	}

	return int(result)
}

// this function for the bin and hex
// it replace the bin_or_hex string by the num_as_atring in the sentenc
func Raplace_Dicimal(sentenc, bin_or_hex, num_as_string string) string {
	arr := strings.Split(sentenc, " ")
	arr[len(arr)-1] = num_as_string

	// sory for this bin_or hex we dont need it
	result := strings.Join(arr, " ")
	return result
}

// this a capitalise function
func Capitalize(sentenc string, number int) string {
	array_of_words := strings.Fields(sentenc)
	// loop throught the array
	n := len(array_of_words) - 1

	// update the number if it is grather the len(array_of_words) and if = 0
	if number == 0 || number == 1 {
		number = 1
	}
	if number > n {
		number = n+1
	}
	for i := n; i > n-number; i-- {
		array_of_words[i] = strings.Title(array_of_words[i])
	}

	result := strings.Join(array_of_words, " ")
	return result
}

// this function to uper the words
func To_Upper(sentenc string, number int) string {
	// lets remove extra spaces

	// array_of_words := strings.Split(sentenc, " ")
	array_of_words := strings.Fields(sentenc)
	// loop throught the array
	n := len(array_of_words) - 1
	// update the number if it is grather the len(array_of_words) and if = 0
	if number == 0 || number == 1 {
		number = 1
	}
	if number > n {
		number = n+1
	}
	for i := n; i > n-number; i-- {
		array_of_words[i] = strings.ToUpper(array_of_words[i])
	}

	result := strings.Join(array_of_words, " ")
	return result
}

// this function to lower the words
func To_Lower(sentenc string, number int) string {
	array_of_words := strings.Fields(sentenc)
	// loop throught the array
	n := len(array_of_words) - 1
	// update the number if it is grather the len(array_of_words) and if = 0
	if number == 0 || number == 1 {
		number = 1
	}
	if number > n {
		number = n+1
	}
	for i := n; i > n-number; i-- {
		array_of_words[i] = strings.ToLower(array_of_words[i])
	}

	result := strings.Join(array_of_words, " ")
	return result
}

// this function just for punctuations traitment

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

// this function handle the single quote only
func Single_Quote(line string) string {
	// now lets handle the single quote
	// single_quotes := regexp.MustCompile(`\s'([^']*)'\s`) // this exprestion handle all line inside
	// result := single_quotes.ReplaceAllStringFunc(line, func(match string) string {
	// 	// match is the line matched  including single quotes
	// 	// Extract the line between the quotes, trim spaces, and return the result
	// 	edited_match := strings.Replace(match," ","",-1)
	// 	//edited_match := strings.Join(test, "")
	// 	return " " + edited_match + " "
	// })

	// line = result

	// return line

	singleQuotesRegex := regexp.MustCompile(`\s'[^']*'`)

	// Replace the single-quoted text with trimmed content
	line = singleQuotesRegex.ReplaceAllStringFunc(line, func(match string) string {
		// Remove surrounding spaces and single quotes
		content := strings.TrimSpace(match)
		trimmedContent := strings.ReplaceAll(content, " ", "")
		return " " + trimmedContent + " "
	})

	// Regular expression to match single quoted text that are not surrounded by spaces
	line = regexp.MustCompile(`'\S+[^']*'`).ReplaceAllStringFunc(line, func(match string) string {
		// Do not modify if it contains contractions or possessives
		if strings.Contains(match, "'") && !strings.Contains(match, " ") {
			return match
		}
		return match
	})

	return line
}

// is punctuation
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
		if IS_Punctuation(string(text[i])) && i < len(text)-1 && (unicode.IsLetter(rune(text[i+1])) || unicode.IsDigit(rune(text[i+1]))) {
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
	)
	text = to_replace.Replace(text)
	

	return text
}

// this the final function that append the newlines

func Append_New_Line(line string) string {
	split_line := strings.Split(line, "~")
	arr := []string{}
	for _, word := range split_line {
		word = strings.TrimSpace(word)
		arr = append(arr, word)
	}
	splited_line := strings.Join(arr, " "+"\n")

	return splited_line
}
