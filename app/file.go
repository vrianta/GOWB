package Application

import (
	"fmt"
	"os"
)

// Create File with a Data in if you do not want any data to be loaded send ""(empty string) as file_data
func CreateFile(file_name string, file_data string) error {
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close() // Ensure the file is closed when done

	// Write data to the file

	if file_data != "" {
		_, err = file.WriteString(file_data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}

	return nil

}
