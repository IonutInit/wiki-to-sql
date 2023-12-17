package sqlutil

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WriteSQL(outputData [][]string, sqlFileName string) error {
	file, err := os.Create(sqlFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("INSERT INTO data(countryId, attributeId, value, year) VALUES\n")
	if err != nil {
		panic(err)
	}

	for i, row := range outputData {
		if i == 0 { // Skip header row
			continue
		}

		value := row[2]
		if _, err := strconv.Atoi(value); err != nil {
			value = "NULL"
		}

		sql := fmt.Sprintf("('%s', '%s', '%s', '%s')", row[0], row[1], row[2], row[3])

		if i < len(outputData)-1 {
			sql += ","
			sql += "\n"
		} else {
			sql += ";"
		}

		_, err := writer.WriteString(sql)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
