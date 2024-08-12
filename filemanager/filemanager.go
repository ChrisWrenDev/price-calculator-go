package filemanager

import (
  "os"
  "bufio"
  "encoding/json"
  "errors"
)

type FileManager struct {
  InputFileName string
  OutputFileName string
}

func (fm FileManager) ReadLines() ([]string, error) {
 	file, err := os.Open(fm.InputFileName)

	if err != nil {
		return nil, errors.New("Failed to open file")
	}

  defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("Failed to read file")
	} 

  return lines, nil
}

func (fm FileManager) WriteResult(fileName string, data any) error {
  file, err := os.Create(fm.OutputFileName)

  if err != nil {
    return errors.New("Failed to create file")
  }

  defer file.Close()

  time.Sleep(3 * time.Second)

  encoder := json.NewEncoder(file)
  err = encoder.Encode(data)

  if err != nil {
    return errors.New("Failed to convert data to JSON")
  }

  return nil
}

func New(inputFileName, outputFileName string) FileManager {
  return FileManager{
    InputFileName: inputFileName,
    OutputFileName: outputFileName
  }
}
