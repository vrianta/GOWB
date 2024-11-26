package Application

// File is holding packages to read .config files and get an object

import (
	"bufio"
	"os"
	"strings"
)

func (application *Application) LoadConfig(FilePath string) (map[string]string, error) {

	file, err := application.assets.Open(FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a map to store the configuration key-value pairs
	config := make(map[string]string)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore empty lines and lines that are comments (e.g., starting with #)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[key] = value
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func LoadConfig(FilePath string) (map[string]string, error) {

	file, err := os.Open(FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a map to store the configuration key-value pairs
	config := make(map[string]string)

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore empty lines and lines that are comments (e.g., starting with #)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[key] = value
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}
