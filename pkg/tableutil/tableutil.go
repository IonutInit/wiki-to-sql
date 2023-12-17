package tableutil

import (
	"fmt"
	"strings"

	"github.com/IonutInit/wiki-to-sql/pkg/config"
)

func FindColumnIndex(header []string, columnName string) (int, error) {
	cleanColumnName := strings.TrimSpace(strings.ToLower(columnName))
	for i, name := range header {
		if strings.TrimSpace(strings.ToLower(name)) == cleanColumnName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("column %s not found", columnName)
}

// looks for the countries column in input
func FindCountryColumnIndex(inputData [][]string, sourceCountries [][]string) (int, error) {
	if len(inputData) == 0 || len(inputData[0]) == 0 {
		return -1, fmt.Errorf("input data is empty")
	}

	countrySet := make(map[string]struct{})
	for _, row := range sourceCountries {
		if len(row) < 1 {
			continue
		}
		countrySet[strings.ToLower(strings.TrimSpace(row[0]))] = struct{}{}
	}

	bestMatchIndex := -1
	bestMatchCount := 0

	for colIndex := range inputData[0] {
		matchCount := 0
		for _, row := range inputData[1:] {
			if len(row) <= colIndex {
				continue
			}
			_, exists := countrySet[strings.ToLower(strings.TrimSpace(row[colIndex]))]
			if exists {
				matchCount++
			}
		}
		if matchCount > bestMatchCount {
			bestMatchCount = matchCount
			bestMatchIndex = colIndex
		}
	}

	if bestMatchIndex == -1 {
		return -1, fmt.Errorf("no country column found")
	}
	return bestMatchIndex, nil
}

func FindAttributeId(dataName string, attributeList []string) (int, error) {
	for i, name := range config.AttributeList {
		if name == dataName {
			return i + 1, nil
		}
	}

	return 0, fmt.Errorf("%s not found in attribute list", dataName)
}
