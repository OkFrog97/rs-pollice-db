package csvRepository

import (
	"encoding/csv"
	"os"
)

type Table = [][]string

// Нужно читать из конфига
const DBPath = "./criminal_db.csv"

func AddRow(row []string) {
	file, err := os.OpenFile(DBPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writeErr := writer.Write(row)
	defer writer.Flush()
	if writeErr != nil {
		panic(writeErr)
	}
}

func GetRow(value string, column string) (row []string) {
	return
}

func DeleteRow(value string, column string) {
	return
}

func UpdateRow(value string, column string) {
	return
}

func ReadAllCSV() (data Table, err error) {
	file, err := os.Open(DBPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}
