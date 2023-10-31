package fileutils

import "os"

func ReadTextFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		// we couldn't read the file
		return "", err
	} else {
		// read operation successful
		return string(content), nil
	}
}