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
	i, err := file.WriteString(text)
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

	text := Extract_Text()
	// split the full text by a newline
	lines := functions.Split_By_Newline(text)

	modifid_text := ""
	// lets loop throught all lines and send the to manipulation
	for _, line := range lines {

		// up dtae our line with previous text
		line = modifid_text + " " + line

		// clear the modefied text
		modifid_text = ""

		// now send this line to manipulation without punctuations traitement
		modifed_line := functions.Destribute_Sentences(line)

		// now send the modifed_line to punctuations traitment
		final_line := functions.Punctuations(modifed_line)
		// now lets put a delimeter "~" to know the end of line
		modifid_text += functions.Expand_Spaces(final_line) + "~"

	}

	// now we have this modified text 
	// send this modefied text to add newlines for a valid  format
	final_text := functions.Append_New_Line(modifid_text)

	// now simply add our final text to result.txt
	Append_Text(final_text)
}

////////////////////////***finaly  the project is done**** /////////////////////////////
///////////////////////               great work          /////////////////////////////
