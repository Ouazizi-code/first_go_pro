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

// this function split text from a delemeter

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
	result = s[startIndex:]

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

	// this part for (hex) and (bin)

	bin_or_hex := ""

	for i := startIndex - 2; i >= 0; i-- {
		char := s[i]
		if char == ' ' {
			break
		}
		bin_or_hex += string(char)
	}
	to_reverse := bin_or_hex
	bin_or_hex = ""
	// reverse yhe string by this method
	for _, v := range to_reverse {
		bin_or_hex = string(v) + bin_or_hex
	}
	fmt.Println(bin_or_hex)

	// send this bin_or_hex to To_Dicimal function
	dicimal := To_Dicimal(bin_or_hex, key_Word)
	fmt.Println(dicimal)
	return result, key_Word, final_int
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
