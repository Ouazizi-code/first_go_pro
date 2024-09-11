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

	// array := functions.Split_Text(text)

	// this for just  to append line by line
	modifid_text := "" // this contain all text
	for i, line := range lines {
		// send this line to remove extra spaces
		//line = functions.Expand_Spaces(line)
		line = modifid_text + line
		modifid_text = ""

		// now send this line to manipulation zithout punctuations traitement
		modifed_line := functions.Destribute_Sentences(line)

		// now send the modifed_line to punctuations traitment
		final_line := functions.Punctuations(modifed_line)
		// now append the modifed line
		if i == len(lines)-1 {
			modifid_text += functions.Expand_Spaces(final_line)
		} else {
			modifid_text += functions.Expand_Spaces(final_line) + "\n"
		}

	}

	Append_Text(modifid_text)
}

// add comment
