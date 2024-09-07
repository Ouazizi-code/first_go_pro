package functions

import (
	"fmt"
	"strings"
)

func Test() {
	fmt.Println("hi your modules is working")
}

func Expand_Spaces(s string) string {
	result := strings.Fields(s)
	valid_Text := strings.Join(result, " ")

	return valid_Text
}
