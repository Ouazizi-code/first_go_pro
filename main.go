package main

import (
	"fmt"
	"os"

	"go-reloaded/functions"
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

// this function distribute each sentence from the  splited array
func Destribut_Sentenes(array, full_result, key_Word string, num int) {
	n := len(array)

	// loop throught the array and send it to manipulation
	for i := 0; i < n; i++ {
		sentence := string(array[i])
		functions.Manipulate_sentenc(sentence, full_result, key_Word, num)
	}
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

func main() {
	// check the args length
	if len(os.Args) < 2 {
		fmt.Println("usage : go run . text.txt result.txt")
		return
	}

	text := Extract_Text()
	Append_Text(text)

	a := "hassan is (a, 6)good    person(in) his zone"
	array := functions.Split_Text(a)
	for _, i := range array {
		fmt.Println(i)
	}

	tes := array[2]
	full_result, key_Word, number := functions.Search_KeyWord(tes)

	fmt.Println(full_result, len(full_result))
	fmt.Println(key_Word)
	fmt.Println(number)
	status := functions.Is_Valid(tes, full_result, key_Word, number)
	if status {
		fmt.Println("the sentense ready to be manipulated")
	} else {
		fmt.Println("the sentense not valid")
	}
}

// add comment
