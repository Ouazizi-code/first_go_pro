package main

import (
	"fmt"
	"os"

	"first_go_pro/functions"
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
	file, err := os.OpenFile(result_file, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
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

// add some
func main() {
	// check the args length
	if len(os.Args) < 2 {
		fmt.Println("usage : go run . text.txt result.txt")
		return
	}

	text := Extract_Text()
	Append_Text(text)
	functions.Test()

	a := "hassan is (a, 6)    good    person     (in) his zone"
	t := functions.Split_Text(a)
	for _, i := range t {
		fmt.Println(i)
	}

	tes := t[0]
	r, k, i := functions.Search_KeyWord(tes)
	fmt.Println(tes)
	fmt.Println(r)
	fmt.Println(k)
	fmt.Println(i)
}

// add comment
