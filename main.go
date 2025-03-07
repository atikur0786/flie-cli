package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	lisCmd := flag.String("list", "", "Path to list all the files in the directory")
	createCmd := flag.String("create", "", "Path to create a new file")
	deleteCmd := flag.String("delete", "", "Path to delete a file")
	readCmd := flag.String("read", "", "Path to read a file")
	writeCmd := flag.String("write", "", "Path to write to a file")
	contentCmd := flag.String("content", "", "Content to write to a file")
	updateCmd := flag.String("update", "", "Path to update a file content")

	flag.Parse()

	if *lisCmd != "" {
		err := listFiles(*lisCmd)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else if *createCmd != "" {
		err := createFile(*createCmd)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else if *deleteCmd != "" {
		err := deleteFile(*deleteCmd)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else if *readCmd != "" {
		err := readFile(*readCmd) 

		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else if *writeCmd != "" && *contentCmd != "" {
		err := writeFile(*writeCmd, *contentCmd)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else if *updateCmd != "" && *contentCmd != "" {
		err := updateFile(*updateCmd, *contentCmd)

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

// Function to create a new file
func createFile(path string) error {
	// Create teh file
	file, err := os.Create(path);

	if err != nil {
		return fmt.Errorf("Error creating file: %v\n", err)
	}

	defer file.Close()

	_, err = file.WriteString("This is a new file created by file cli\n")

	if err != nil {
		return fmt.Errorf("Error writing to file: %v\n", err)
	}

	return nil
}

// Function to delete a file
func deleteFile(path string) error {
	// Delete the file 
	err := os.Remove(path)

	if err != nil {
		return fmt.Errorf("Error deleting file: %v\n", err)
	}

	return nil
}

// Function to read a file
func readFile(path string) error {
	// Read the file
	file, err := os.ReadFile(path)

	if err != nil {
		return fmt.Errorf("Error reading file: %v\n", err)
	}

	fmt.Println(string(file))

	return nil
}

// Function to write to a file
func writeFile(path string, content string) error {
	// Open the file
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		return fmt.Errorf("Error opening file: %v\n", err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		return fmt.Errorf("Error writing to file: %v\n", err)
	}

	fmt.Println("Content written to file")
	return nil
}

// Function to update a file content
func updateFile(path string, content string) error {
	// Open the file
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644);

	if err != nil {
		return fmt.Errorf("Error opening file: %v\n", err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		return fmt.Errorf("Error writing to file: %v\n", err)
	}

	fmt.Println("Content updated in file")
	return nil
}