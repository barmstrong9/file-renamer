package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Define the root directory to search for files in
	root := "/Users/brandonarmstrong/Documents/Music"

	// Recursively search the root directory for all files
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current file is a regular file (i.e., not a directory)
		if !info.Mode().IsRegular() {
			return nil
		}

		// Check if the filename contains spaces or hyphens
		if strings.Contains(info.Name(), " ") || strings.Contains(info.Name(), "-") || strings.Contains(info.Name(), "__") {
			// Replace spaces with underscores and hyphens with nothing
			re := regexp.MustCompile("[ -]")
			newName := re.ReplaceAllString(info.Name(), "_")

			// Replace consecutive underscores with a single underscore
			newName = strings.ReplaceAll(newName, "__", "_")

			// Get the full path to the current file
			oldPath := filepath.Join(root, info.Name())

			// Get the full path to the new file with underscores
			newPath := filepath.Join(root, newName)

			// Rename the file
			err := os.Rename(oldPath, newPath)
			if err != nil {
				log.Printf("Error renaming file %s to %s: %s", oldPath, newPath, err)
				return err
			}

			// Print a message to indicate that the file was renamed
			fmt.Printf("Renamed %s to %s\n", oldPath, newPath)
		}

		return nil
	})

	// Check if there was an error while walking the directory
	if err != nil {
		log.Fatal(err)
	}

	// Print a message to indicate that the script has completed
	fmt.Println("Finished renaming files.")
}
