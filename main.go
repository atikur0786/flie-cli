package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	lisCmd := flag.String("list", "", "Path to list all the files in the directory")

	flag.Parse()

	if *lisCmd != "" {
		err := listFiles(*lisCmd)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println(("Usage: filecli -list /path/to/directory"))
	}
}

// Function to list files in a directory
func listFiles(path string) error {
	// Open the directory
	dir, err := os.Open(path)
	
	if err != nil {
		return fmt.Errorf("Error opening directory: %v\n", err)
	}

	// Close the directory
	defer dir.Close()

	// Read the directory contents
	files, err := dir.Readdirnames(0)

	if err != nil {
		return fmt.Errorf("Error reading directory contents: %v\n", err)
	}

	// Print the list files
	for _, file := range files{
		fmt.Println(file)
	}

	return nil
}