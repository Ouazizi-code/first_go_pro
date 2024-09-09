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
		}
	}

	// this the fill text inside braces
	result = Expand_Spaces(s[startIndex:])

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
func Rmove_braces(sentenc, delimiter string) string {
	result := strings.Replace(sentenc, delimiter, "", 1)
	return result
}

// this the sentence manipulation function
func Sentenc_Mainpulation(valid_sentence, key_word string, number int, status bool) string {
	result := ""
	// check the keyword if valid
	if status {

		result = "good"

	} else {

		return valid_sentence

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
		// check if this sentence caontein  flag ()
		// send this sentenc to search_keyword to find the keyword
		full_result, key_word, number := Search_KeyWord(sentenc)
		fmt.Println(full_result, key_word, number)
		// lets check if the keyword is valid
		status := Is_Valid(full_result, key_word, number)
		// lets modifid the sentence and remove the braces if exist
		valid_sentence := Rmove_braces(sentenc, full_result)
		// lets send this valid sentence to manipulation depend on the keyword and status
		manipulated_sentenc := Sentenc_Mainpulation(valid_sentence, key_word, number, status)
		result += manipulated_sentenc
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
