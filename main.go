package main

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func generateAlias(firstName, lastName string) string {
	companyTag := ""

	return fmt.Sprintf("%c%s"+companyTag,
		strings.ToLower(firstName)[0],
		strings.ToLower(lastName))
}

func main() {
	excelFile := ""

	f, err := excelize.OpenFile(excelFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	rows, err := f.GetRows("Sheet1")

	if err != nil {
		fmt.Printf("Unable to get rows: %v", err)
	}

	for i, row := range rows {
		firstName := strings.TrimSpace(row[0])
		lastName := strings.TrimSpace(row[1])

		if firstName == "" || lastName == "" {
			fmt.Printf("Row %d has empty first or last name, skipping", i+1)
		}

		alias := generateAlias(firstName, lastName)
		fmt.Println(alias)
	}
}
