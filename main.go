package main

import (
	"fmt"
	"os"
	"strings"

	"go-reloaded/functions"
)

// this function to extract text from file and append it to another file
func Extract_Text() string {
	// lets read the content of the file
	content, err := os.ReadFile("semple.txt")
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

	// remove etra spaces and new lines
	text = strings.TrimSpace(text)
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
	if len(os.Args) != 3 {
		fmt.Println("Error: not enough arguments")
		fmt.Println("Usage: go run . semple.txt result.txt")
		return
	}

	// should respect file naming
	// let whrite a condition

	sourceFile, resultFile := os.Args[1], os.Args[2]
	if sourceFile != "semple.txt" || resultFile != "result.txt" {
		fmt.Println("Error: invalid file naming")
		fmt.Println("Usage: go run . semple.txt result.txt")
		return
	}

	// lets extract our text
	text := Extract_Text()
	//fmt.Println(text,"line 67")
	// split the full text by a newline
	line := functions.Split_By_Newline(text)
	fmt.Println(line," line 70")
	
	
	
	//return
	// now send this line to vowel manipulation
	vowled_line := functions.Vowles_manioulation(line)

	// now send this voweled line to punctuations manipulation
	punctuationed_line := functions.Punctuations(vowled_line)

	// now send the punctuationed_line to ingle_Quote traitment
	single_quoteed_line := functions.Single_Quote(punctuationed_line)
	single_quoteed_line = functions.Expand_Spaces(single_quoteed_line)

	// now send the single_quoteed_line to manipulation traitment
	manipulated_line := functions.Destribute_Sentences(single_quoteed_line)
	fmt.Println(manipulated_line,"line 88")
	// now we have this manipulated_line
	// send this manipulated_line  to add newlines for a valid  format
	final_text := functions.Append_New_Line(manipulated_line)

	// now simply add our final text to result.txt
	Append_Text(final_text)
}

////////////////////////*** finaly  the project is done ****/////////////////////////////
///////////////////////************* great work ***********/////////////////////////////
