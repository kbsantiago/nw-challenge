package util

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

//CsvReader read CSV file type
func CsvReader(filename string, hasHeader bool, out chan string) {
	defer close(out)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if hasHeader {
			hasHeader = false
			continue
		}

		out <- record[0]
	}
}
