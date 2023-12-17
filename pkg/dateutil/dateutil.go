package dateutil

import (
	"fmt"
	"strconv"
	"time"

	"github.com/IonutInit/wiki-to-sql/pkg/config"
)

// defaults to the current or custom year if no year column is found.
func DefaultYear() string {
	if config.CustomYear == 0 {
		currentYear := time.Now().Year()
		fmt.Println("No date found. Defaulting to current year.")
		return strconv.Itoa(currentYear)
	}
	fmt.Println("No date column found. Defaulting to custom year.")
	return strconv.Itoa(config.CustomYear)
}

// extracting date from various date formats
func ExtractYearFromDate(dateStr string) string {
	for _, format := range config.PossibleDateFormats {
		if parsedTime, err := time.Parse(format, dateStr); err == nil {
			return strconv.Itoa(parsedTime.Year())
		}
	}
	return "0"
}
