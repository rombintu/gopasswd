package csvman

import (
	"encoding/csv"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func Parse_csv(file multipart.File) [][]string {

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4
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

func Export_csv(res http.ResponseWriter, data [][]string) http.ResponseWriter {
	csv_file, err := os.Create("MyPasswords.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(res)

	writer.WriteAll(data)

	defer writer.Flush()
	defer csv_file.Close()

	// err := writer.Error()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return res
}
