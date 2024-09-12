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

	text := Extract_Text()
	//fmt.Println(text)
	// split the full text by a newline
	line := functions.Split_By_Newline(text)
	//fmt.Println(lines,len(lines))
	
	// lets loop throught all lines and send the to manipulation
	//for _, line := range lines {

		// up dtae our line with previous text
		//line := modifid_text + " " + line

		// clear the modefied text
		//fmt.Println(modifid_text)
		

		// now send this line to manipulation without punctuations traitement
		vowled_line := functions.Vowles_manioulation(line)
		//fmt.Println(vowled_line)
		punctuationed_line := functions.Punctuations(vowled_line)
		//fmt.Println(punctuationed_line)

		// now send the modifed_line to punctuations traitment
		final_line := functions.Destribute_Sentences(punctuationed_line)
		// now lets put a delimeter "~" to know the end of line
		//modifid_text += functions.Expand_Spaces(final_line) + "~"

	//}

	// now we have the modefid text
	// lets send it to single cote manipulation
	//modifid_text = functions.Real_punctions(modifid_text)

	single_quoteed_line := functions.Single_Quote(final_line)
	fmt.Println(single_quoteed_line)

	// now we have this modified text
	// send this modefied text to add newlines for a valid  format
	//final_text := functions.Append_New_Line(Single_Quoteed_text)

	// now simply add our final text to result.txt
	//Append_Text(final_text)
}

////////////////////////*** finaly  the project is done ****/////////////////////////////
///////////////////////************* great work ***********/////////////////////////////
