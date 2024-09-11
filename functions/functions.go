package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)



// this function about removing extra spaces
func Expand_Spaces(s string) string {
	// the field work with spaces or more spaces
	result := strings.Fields(s) 
	valid_Text := strings.TrimSpace(strings.Join(result, " "))

	return valid_Text
}

// this function split text by newline
func Split_By_Newline(text string) []string {
	array := strings.Split(text, "\n")
	return array
}

// this function split text from a )
func Split_Text(s string) []string {
	result := []string{}
	startIndex := 0
	endIndex := 0
	delemeter := ')'

	for i, char := range s {
		if char == delemeter {
			endIndex = i
			result = append(result, Expand_Spaces(s[startIndex:endIndex+1]))
			startIndex = endIndex + 1
		}
	}
	// append the rest text in the end if exist
	result = append(result, Expand_Spaces(s[startIndex:]))

	return result
}

// this finction to search for keyword betwen braces
func Search_KeyWord(s string) (string, string, string, int) {
	result := ""
	key_Word := ""
	Int_AsString := ""
	final_int := 0
	startIndex := 0

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

	// this the fill text inside braces
	for_braces := s[startIndex:] // variable not edited for remove braces function
	result= strings.Replace(s[startIndex:], " ", "", -1)
	result= Expand_Spaces(result)
	

	// lests extract just keyword
	for i := 0; i < len(result); i++ {
		char := result[i]
		if char >= 'a' && char <= 'z' {
			key_Word += string(char)
		}
	}

	// lets extract thhe number inside the braces
	for i := 0; i < len(result); i++ {
		char := result[i]
		if char >= '0' && char <= '9' {
			Int_AsString += string(char)
		}
	}
	// convert the string into valid number
	final_int, _ = strconv.Atoi(Int_AsString)

	return for_braces, result, key_Word, final_int
}

// this function check the valid keyword
func Is_Valid(full_resul, key_word string, number int) (bool, bool) {
	// convert the number to string
	num_as_string := strconv.Itoa(number)
	first_case := "(" + key_word + ")"
	second_case := "(" + key_word + "," + num_as_string + ")"

	// this bool s just for removing braces or not
	remove_braces_or_not := true

	// check now the full result if it a valid flag or not
	// all key words
	if key_word == "cap" || key_word == "up" || key_word == "low" || key_word == "bin" || key_word == "hex" {
		if full_resul == first_case || full_resul == second_case {
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
	result := ""
	bin_or_hex := ""

	result = strings.Replace(sentenc, delimiter, "", 1)
	result = Expand_Spaces(result)

	// so let do some work for bin and hex
	// extract the string befor braces asn well as is a binary or hex
	arr := strings.Split(result, " ")
	// after removing braces the last index is the bin_or_hex
	bin_or_hex = arr[len(arr)-1]

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
	array_of_sentences := Split_Text(line)
	n := len(array_of_sentences)
	
	result := ""

	// now lets destribute our sentences to manipulate with a for loop
	for i := 0; i < n; i++ {

		sentenc := array_of_sentences[i]
		// send this sentenc to search_keyword to find the keyword
		for_braces, full_result, key_word, number := Search_KeyWord(sentenc)
		// lets check if the keyword is valid
		status, remove_braces_or_not := Is_Valid(full_result, key_word, number)
		// lets modifid the sentence and remove the braces if exist
		valid_sentence, bin_or_hex := Rmove_braces(sentenc, for_braces, remove_braces_or_not)
		// valid_sentence = Expand_Spaces(valid_sentence)
		// lets send this valid sentence to manipulation depend on the keyword and status
		manipulated_sentenc := Sentenc_Mainpulation(result+valid_sentence, full_result, key_word, bin_or_hex, number, status)
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
	array_of_words := strings.Split(sentenc, " ")
	// loop throught the array
	n := len(array_of_words) - 1

	// update the number if it is grather the len(array_of_words) and if = 0
	if number == 0 || number == 1 {
		number = 1
	}
	if number > len(array_of_words) {
		number = len(array_of_words)
	}
	for i := n; i > n-number; i-- {
		if array_of_words[i] == "\n" {
			continue
		}

		array_of_words[i] = strings.Title(array_of_words[i])
	}

	result := strings.Join(array_of_words, " ")
	return result
}

// this function to uper the words
func To_Upper(sentenc string, number int) string {
	array_of_words := strings.Split(sentenc, " ")
	// loop throught the array
	n := len(array_of_words) - 1
	// update the number if it is grather the len(array_of_words) and if = 0
	if number == 0 || number == 1 {
		number = 1
	}
	if number > len(array_of_words) {
		number = len(array_of_words)
	}
	for i := n; i > n-number; i-- {
		array_of_words[i] = strings.ToUpper(array_of_words[i])
	}

	result := strings.Join(array_of_words, " ")
	return result
}

// this function to lower the words
func To_Lower(sentenc string, number int) string {
	array_of_words := strings.Split(sentenc, " ")
	// loop throught the array
	n := len(array_of_words) - 1
	// update the number if it is grather the len(array_of_words) and if = 0
	if number == 0 || number == 1 {
		number = 1
	}
	if number > len(array_of_words) {
		number = len(array_of_words)
	}
	for i := n; i > n-number; i-- {
		array_of_words[i] = strings.ToLower(array_of_words[i])
	}

	result := strings.Join(array_of_words, " ")
	return result
}

// this function just for punctuations traitment

func Punctuations(line string) string {
	// art just for punctuations
	punctuations := regexp.MustCompile(`\s([.,!?;:])`)
	line = punctuations.ReplaceAllString(line, `$1`)

	// now lets handle the single quote
	single_quotes := regexp.MustCompile(`'([^']*)'`) // this exprestion handle all text insode
	result := single_quotes.ReplaceAllStringFunc(line, func(match string) string {
		// match is the text matched  including single quotes
		// Extract the text between the quotes, trim spaces, and return the result
		edited_match := strings.TrimSpace(match[1 : len(match)-1])
		return " " + "'" + edited_match + "'" + " "
	})

	line = result

	// now lets handle the vowle case
	vowels := regexp.MustCompile(`([aeiouhAEIOUH])`)
	expretion := regexp.MustCompile(`\b([aA])\s+([aeiouhAEIOUH])`)
	line = expretion.ReplaceAllStringFunc(line, func(match string) string {
		// match is the text matched, including "a" and the following word
		// Extract the word after "a", trim spaces, and return the result
		words := strings.Split(match, " ")
		next_word := strings.ToLower(words[1])

		// check if the a is upper or lower
		if vowels.MatchString(next_word) {
			if strings.HasPrefix(match, "A ") {
				return "An " + words[1]
			} else {
				return "an " + words[1]
			}
		}
		return match
	})

	return line
}

// this the final function that append the newlines

func Append_New_Line(text string)(string){
	split_text := strings.Split(text,"~")
	arr := []string {}
	for _, word := range split_text {
		word = strings.TrimSpace(word)
		arr = append(arr, word)
	}
	splited_text := strings.Join(arr," "+"\n")

	return splited_text

}