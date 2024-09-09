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
	return text
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
	i, err := file.WriteString(text + "\n")
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
	// split the full text by a newline
	lines := functions.Split_By_Newline(text)

	//array := functions.Split_Text(text)

	// this for just  to append line by line
	modifid_text := "" // this contain all text
	for _, line := range lines {
		// send this line to remove extra spaces
		line = functions.Expand_Spaces(line)
		//fmt.Println(line)
		modifed_line := functions.Destribute_Sentences(line)
		//fmt.Println(modifed_line)
		// now append the modifed line

		modifid_text += modifed_line + "\n"
	}

	Append_Text(modifid_text)

}

// add comment
