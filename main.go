package main

import (
	"first_go_pro/functions"
	"fmt"
	"os"
)

// this function to extract text from file and append it to another file
func Extract_Text() string {

	// lets read the content of the file
	content, err := os.ReadFile("test.txt")
	// handle the errors
	if err != nil {
		fmt.Println("eror :", err)
	}

	// convert the content into string
	text := string(content)
	valid_Text := functions.Expand_Spaces(text)
	return valid_Text
}

// this function append text into result.txt
func Append_Text(text string) {
	// lets append the text int result.txt
	result_file := os.Args[2]
	file, err := os.OpenFile(result_file, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// handle the errors
	if err != nil {
		fmt.Println("eror :", err)
	}

	// lets write the text
	i, err := file.WriteString(text)
	i++
	// handle the errors
	if err != nil {
		fmt.Println("eror :", err)
	}
}

func main() {
	// check the args length
	if len(os.Args) < 2 {
		fmt.Println("usage : go run . text.txt result.txt")
		return
	}

	text := Extract_Text()
	Append_Text(text)
	functions.Test()
}

// add comment
