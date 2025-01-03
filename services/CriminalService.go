package CriminalService

import (
	"errors"
	"fmt"
	"reflect"
	"rs-police-db/domain"
	csvRepository "rs-police-db/repository"
	"strconv"
)

func AddCriminalToTable(name string, age int, gender string, crimeDate string, crimeType string, crimeDescription string, status string) {
	c := createCriminal(name, age, gender, crimeDate, crimeType, crimeDescription, status)
	saveCriminal(*c)
}

func createCriminal(name string, age int, gender string, crimeDate string, crimeType string, crimeDescription string, status string) *domain.Criminal {
	return &domain.Criminal{
		Name:             name,
		Age:              age,
		Gender:           gender,
		CrimeDate:        crimeDate,
		CrimeType:        crimeType,
		CrimeDescription: crimeDescription,
		Status:           status,
	}
}

func saveCriminal(c domain.Criminal) {
	var s []string = crimianToString(c)
	fmt.Println(s)
	csvRepository.AddRow(s)
	return
}

func crimianToString(c domain.Criminal) (s []string) {
	var r = reflect.ValueOf(c)
	for i := 0; i < r.NumField(); i++ {
		s = append(s, fmt.Sprintf("%v", r.Field(i).Interface()))
	}
	return
}

func ShowAllCriminals() {
	data, err := csvRepository.ReadAllCSV()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	printTable(data)
	fmt.Println()
}

func printTable(data csvRepository.Table) {
	for rowIndex, row := range data {
		fmt.Print(rowIndex, " ")
		for _, col := range row {
			fmt.Print(col)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func FindByName(name string) *domain.Criminal {
	row, err := findByField("name", name)
	if err != nil {
		panic(err)
	}
	return stringToCriminal(row)
}

func findByField(field string, searchValue string) ([]string, error) {
	data, err := csvRepository.ReadAllCSV()
	if err != nil {
		panic(err)
	}
	var columnNamesRow = data[0]
	var searchingValueColumnIndex int

	// Find column index for searching
	for colIndex, colName := range columnNamesRow {
		if colName == field {
			searchingValueColumnIndex = colIndex
		}
	}

	// Find Searching Value
	for _, row := range data[1:] {
		if row[searchingValueColumnIndex] == searchValue {
			return row, nil
		}
	}
	return nil, errors.New("Searching field not found")
}

func stringToCriminal(s []string) (c *domain.Criminal) {
	age, err := strconv.Atoi(s[1])
	if err != nil {
		age = 0
	}
	return createCriminal(s[0], age, s[2], s[3], s[4], s[5], s[6])
}
