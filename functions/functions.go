package functions

import (
	"fmt"
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
