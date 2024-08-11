package filemanager

import (
  "os"
  "bufio"
  "encoding/json"
  "errors"
)

func ReadLines(fileName string) ([]string, error) {
 	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read file")
	} 

  file.Close()
  return lines, nil
}

func WriteJSON(fileName string, data any) error {
  file, err := os.Create(fileName)

  if err != nil {
    return errors.New("Failed to create file")
  }

  encoder := json.NewEncoder(file)
  err = encoder.Encode(data)

  if err != nil {
    file.Close()
    return errors.New("Failed to convert data to JSON")
  }

  file.Close()
  return nil
}
