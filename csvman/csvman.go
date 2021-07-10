package csvman

import (
	"encoding/csv"
	"log"
	"mime/multipart"
	"os"
)

func Parse_csv(file multipart.File) [][]string {

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 5
	reader.Comment = '#'
	reader.LazyQuotes = true

	var data [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			log.Println(err)
			break
		}
		data = append(data, record)
	}
	return data
}

func Export_csv(data [][]string) multipart.File {
	csv_file, err := os.Create("MyPasswords.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(csv_file)

	writer.WriteAll(data)
	writer.Flush()
	csv_file.Close()

	return csv_file
}
