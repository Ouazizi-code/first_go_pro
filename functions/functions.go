package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func Test() {
	fmt.Println("hi your modules is working")
}

// this function about removing extra spaces
func Expand_Spaces(s string) string {
	result := strings.Fields(s)
	valid_Text := strings.Join(result, " ")

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
func Search_KeyWord(s string) (string, string, int) {
	result := ""
	key_Word := ""
	Int_AsString := ""
	final_int := 0
	startIndex := 0

	for i, char := range s {
		if char == '(' {
			startIndex = i
			break
		} else {
			startIndex = len(s)
		}
	}

	// this the fill text inside braces
	test := strings.Replace(s[startIndex:], " ", "", -1)
	result = test

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

	return result, key_Word, final_int
}

// this function check the valid keyword
func Is_Valid(full_resul, key_word string, number int) bool {
	// convert the number to string
	num_as_string := strconv.Itoa(number)
	first_case := "(" + key_word + ")"
	second_case := "(" + key_word + "," + num_as_string + ")"

	// check now
	if full_resul == first_case || full_resul == second_case {
		return true
	}

	return false
}

// this function remove only braces and return valid string
func Rmove_braces(sentenc, delimiter string) (string, string) {
	// in this case our delimiter contains any thing inside braces
	// exemple delimiter = (cap,5) passed as params
	result := strings.Replace(sentenc, delimiter, "", 1)

	// so let do some work for bin and hex
	// extract the string befor braces asn well as is a binary or hex
	arr := strings.Split(result, " ")
	// after removing braces the last index is the bin_or_hex
	bin_or_hex := arr[len(arr)-1]

	return result, bin_or_hex
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
func Sentenc_Mainpulation(valid_sentence, key_word, bin_or_hex string, number int, status bool) string {
	result := ""
	// check the keyword if valid
	if status {

		result = "good"
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
	fmt.Println(array_of_sentences, n)
	result := ""

	// now lets destribute our sentences to manipulate with a for loop
	for i := 0; i < n; i++ {
		sentenc := array_of_sentences[i]
		// send this sentenc to search_keyword to find the keyword
		full_result, key_word, number := Search_KeyWord(sentenc)
		fmt.Println(full_result, key_word, number)
		// lets check if the keyword is valid
		status := Is_Valid(full_result, key_word, number)
		// lets modifid the sentence and remove the braces if exist
		valid_sentence, bin_or_hex := Rmove_braces(sentenc, full_result)
		valid_sentence = Expand_Spaces(valid_sentence)
		// lets send this valid sentence to manipulation depend on the keyword and status
		manipulated_sentenc := Sentenc_Mainpulation(result+valid_sentence, key_word, bin_or_hex, number, status)
		// refresh the result and concat it with  the valid sentence
		result = ""
		result += manipulated_sentenc + " "
	}

	return result
}

// this function convert hex and binary  to dicimal

func To_Dicimal(s, key_Word string) int {
	var result int64
	switch key_Word {
	case "bin":
		result, err := strconv.ParseInt(s, 2, 64)
		// handle erors
		if err != nil {
			fmt.Println("erour", err)
			return 0
		}
		return int(result)
	case "hex":
		result, err := strconv.ParseInt(s, 16, 64)
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
	result := strings.Replace(sentenc, bin_or_hex, num_as_string, 1)
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
