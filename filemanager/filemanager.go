package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("Error while creating file")
	}

	encoder := json.NewEncoder(file)
	errorEnc := encoder.Encode(data)

	if errorEnc != nil {
		return errors.New("Failed to convert data to json")
	}

	return nil
}

func (fm FileManager) ReadLines() ([]string, error) {

	file, error := os.Open(fm.InputFilePath)

	if error != nil {
		file.Close()
		return nil, errors.New("Error while opening file")
	}

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Error while scaning file")
	}
	file.Close()
	return lines, nil

}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
