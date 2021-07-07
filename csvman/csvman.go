package csvman

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
)

func Parse_csv(file multipart.File) [][]string {

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 5
	reader.Comment = '#'
	reader.LazyQuotes = true

	var data [][]string
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		data = append(data, record)
	}
	return data
}
