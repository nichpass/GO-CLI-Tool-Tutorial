package main

// import reqs for tests to work
import (
	"flag"
	"os"
	"reflect"
	"testing"
)

// so arg is a var named 't' of type *testing.T
func Test_getFileData(t *testing.T) {
	// define test slice (like a list). Each test has the following properties:
	tests := []struct {
		name    string    // name of the test
		want    inputFile // inputFile instance we want the function to return
		wantErr bool      // whether or not we want an error
		osArgs  []string  // command line args used for the test
	}{
		// Here we define the params for each unit test: inputs and outputs
		{"Default parameters", inputFile{"test.csv", "comma", false}, false, []string{"cmd", "test.csv"}},
		{"No parameters", inputFile{}, true, []string{"cmd"}},
		{"Semicolon enabled", inputFile{"test.csv", "semicolon", false}, false, []string{"cmd", "--separator=semicolon", "test.csv"}},
		{"Pretty enabled", inputFile{"test.csv", "comma", true}, false, []string{"cmd", "--pretty", "test.csv"}},
		{"Pretty and semicolon enabled", inputFile{"test.csv", "semicolon", true}, false, []string{"cmd", "--pretty", "--separator=semicolon", "test.csv"}},
		{"Separator not identified", inputFile{}, true, []string{"cmd", "--separator=pipe", "test.csv"}},
	}

	// iterating over the above defined test slice
	for _, tt := range tests {
		// why name the parameter the same name as that above? shouldn't this throw an error?
		t.Run(tt.name, func(t *testing.T) {

			// saves the actual
			actualOSArgs := os.Args
			// func runs after the test is done
			defer func() {
				// restore original os.Args
				os.Args = actualOSArgs
				// reset flag command line to parse flags again
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.osArgs             //set specific command args for this test
			got, err := getFileData()       // running the function to test
			if (err != nil) != tt.wantErr { // assert whether we got the correct error value or not
				t.Errorf("getFileData() error %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) { // assert whether or not we got the correct wanted value
				t.Errorf("getFileData() = %v, want %v", got, tt.want)
			}
		})
	}

}
