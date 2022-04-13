package csvHelper

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
)

func ReadCSV(fileHeader *multipart.FileHeader) ([][]string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}(file)

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("err")
		return nil, err
	}

	var result [][]string

	for _, line := range lines[1:] {
		data := []string{line[0], line[1]}
		result = append(result, data)
	}

	return result, nil
}
