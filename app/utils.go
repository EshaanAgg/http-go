package main

import (
	"fmt"
	"io"
	"os"
)

func (s *Server) getFileContent(fileName string) ([]byte, error) {
	// Open the file
	file, err := os.Open(fmt.Sprintf("%s/%s", s.FilesBaseDir, fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content into a byte slice
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
