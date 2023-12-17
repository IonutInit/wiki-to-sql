package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/IonutInit/wiki-to-sql/pkg/config"
	"github.com/IonutInit/wiki-to-sql/pkg/csvutil"
	"github.com/IonutInit/wiki-to-sql/pkg/dateutil"
	"github.com/IonutInit/wiki-to-sql/pkg/sqlutil"
	"github.com/IonutInit/wiki-to-sql/pkg/tableutil"
)

func main() {
	dataName := config.DataName

	attributeIdInt, err := tableutil.FindAttributeId(dataName, config.AttributeList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	attributeId := strconv.Itoa(attributeIdInt)

	sourceCountriesPath := "data/source_countries.csv"
	inputFilePath := "data/input/" + dataName + ".csv"
	outputDir := "data/output"
	outputFileName := outputDir + "/" + dataName + ".csv"
	sqlFileName := outputDir + "/" + dataName + ".sql"
	inputDir := "data/input"

	var year string
	var yearIndex int
	// yearFound := false

	// reading inputs
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		fmt.Println("Input directory does not exist")
		os.Exit(1)
	}

	sourceCountries, err := csvutil.ReadCsv(sourceCountriesPath)
	if err != nil {
		panic(err)
	}

	inputData, err := csvutil.ReadCsv(inputFilePath)
	if err != nil {
		panic(err)
	}

	// destructuring input
	// reading all columns
	varIndex, err := tableutil.FindColumnIndex(inputData[0], dataName)
	if err != nil {
		varIndex, err = tableutil.FindColumnIndex(inputData[0], config.CustomValueName)
		if err != nil {
			fmt.Println("Available columns:", inputData[0])
			panic(err)
		}
	}

	// finding the country column
	countryColumIndex, err := tableutil.FindCountryColumnIndex(inputData, sourceCountries)
	if err != nil {
		panic(err)
	}

	// attempting to find the date column
	for _, colName := range config.PossibleDateColumns {
		yearIndex, err = tableutil.FindColumnIndex(inputData[0], colName)
		if err == nil {
			break
		}
	}

	// creating a years map if date column is found, otherwise defaulting
	yearMap := make(map[string]string)
	if yearIndex != -1 {
		for _, row := range inputData[1:] {
			countryName := strings.ToLower(strings.TrimSpace(row[countryColumIndex]))
			dateStr := row[yearIndex]
			extractedYear := dateutil.ExtractYearFromDate(dateStr)

			if extractedYear != "0" {
				yearMap[countryName] = extractedYear
			} else {
				yearMap[countryName] = year
			}

		}
	} else {
		year = dateutil.DefaultYear()
	}

	// creating an input map
	inputMap := make(map[string]string)
	for _, row := range inputData[1:] {
		countryName := strings.ToLower(strings.TrimSpace(row[countryColumIndex]))
		inputMap[countryName] = row[varIndex]
	}

	// checking if output directory exists; if not, it creates one
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	// creating output data
	outputData := [][]string{{"country_id", "attributeId", "value", "year"}}

	// population output
	for _, row := range sourceCountries[1:] { // starts from the second row, in order to skip the header
		if len(row) < 2 {
			continue // skips if the row doesn't have enough columns
		}
		countryName, countryID := strings.ToLower(strings.TrimSpace(row[0])), row[1]

		value, ok := inputMap[countryName]

		// checks against country name variations
		if !ok {
			found := false
			if variations, exists := config.CountryVariations[countryName]; exists {
				for _, altName := range variations {
					if val, ok := inputMap[altName]; ok {
						value = val
						found = true
						break
					}
				}
			}

			if !found {
				value = "n/a"
				fmt.Printf("No data found for: %s\n", countryName)
			}
		} else {
			if config.PrintAllData {
				fmt.Printf("Data found for: %s, Value: %s\n", countryName, value)
			}
		}

		rowYear := year
		if yearData, found := yearMap[countryName]; found {
			rowYear = yearData
		}

		outputData = append(outputData, []string{countryID, attributeId, value, rowYear})
	}

	// writing output
	err = csvutil.WriteCsv(outputFileName, outputData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data written to %s.csv\n", dataName)

	err = sqlutil.WriteSQL(outputData, sqlFileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data written to %s.sql\n", dataName)
}
