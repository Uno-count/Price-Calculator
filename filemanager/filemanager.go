package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		// file.Close()
		return nil, errors.New("failed to read line in file")
	}

	// file.Close()

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OuputFilepath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second) //make up scenario for goroutine

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to json")
	}

	// file.Close()

	return nil
}

func New(inputPath, ouputPath string) FileManager {
	return FileManager{
		InputFilePath: inputPath,
		OuputFilepath: ouputPath,
	}
}
