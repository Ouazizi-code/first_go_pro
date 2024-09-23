# Text Manipulation Tool

This Go application extracts text from a specified input file, processes it according to a series of transformations, and appends the resulting text to an output file. The program offers features such as manipulating case, handling punctuation, and expanding or contracting whitespace.

## Features

- Reads content from an input file (`semple.txt`).
- Processes text to manipulate case (capitalize, uppercase, lowercase).
- Handles binary and hexadecimal conversions.
- Removes extra spaces and formats punctuation correctly.
- Writes the final processed text to an output file (`result.txt`).

## Requirements

- Go 1.16 or higher
- An input file named `semple.txt` containing the text to be processed.

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd go-reloaded
2. Ensure you have a semple.txt file in the project root with the text you wish to manipulate.

## Usage

Run the program with the input and output file names as command-line arguments:
`go run main.go semple.txt result.txt`

## Note

- The program checks for the correct number of command-line arguments and validates file names.
- It will terminate and display error messages if the specified input file does not exist or if there are issues reading or writing files.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.
