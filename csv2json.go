package main

import (
	"errors"
	"flag"
	"os"
)

type inputFile struct {
	filepath  string
	separator string // period, comma, etc.
	pretty    bool   // if JSON is well formatted or not
}

/*
Steps
- validate number of args
- parse command line args
- validate separator arg
- return inputFile instance if no errors
*/
func getFileData() (inputFile, error) {
	// validate we're getting the right num of args
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("A filepath arguments is required")
	}

	// defining optional flags via the Flag package from standard library
	// 3 args: flag's name, default value, and a short description for the --help section
	separator := flag.String("separator", "comma", "Column separator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse() // This will parse all the args from the command line

	fileLocation := flag.Arg(0) // only arg that isn't a flag option is the CSV file location

	// Validate whether we received "comma" or "semicolon" from parsed args (only 2 valid options)
	// if didn't get either of those, return an error

	// assuming that flag object returns pointers since the tutorial uses dereference operators on it
	if !(*separator == "comma" || *separator == "semicolon") {
		return inputFile{}, errors.New("Only comma or semicolon separators are allowed")
	}

	// if reach endpoint, program args are validated
	// return correspoinding struct instance will required data

	return inputFile{fileLocation, *separator, *pretty}, nil
}

func main() {
	fileData, err := getFileData()
}
