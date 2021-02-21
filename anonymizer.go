package main

import (
	"strings"
)

func anonymizeResultSet(resultSet []map[string]interface{}) ([]map[string]interface{}, error) {
	dataColumnAggregated := make(map[string][]interface{}, 0)

	// get the data sample to be used to generate the anonymized data
	for _, row := range resultSet {
		for columnName, columnValue := range row {
			dataColumnAggregated[columnName] = append(dataColumnAggregated[columnName], columnValue)
		}
	}

	anonymizedResult := make([]map[string]interface{}, len(resultSet))
	for columnName, columnValuesArray := range dataColumnAggregated {
		anonymizedList := anonymizeList(columnValuesArray)

		for i := 0; i < len(anonymizedResult); i++ {
			if anonymizedResult[i] == nil {
				anonymizedResult[i] = make(map[string]interface{}, 0)
			}
			anonymizedResult[i][columnName] = anonymizedList[i]
		}
	}

	return anonymizedResult, nil
}

func anonymizeList(listSet []interface{}) []interface{} {
	firstTermsUnique := make(map[string]bool, 0)
	firstTerms := make([]string, 0)
	remainTerms := make([]string, 0)
	remainTermsUnique := make(map[string]bool, 0)
	for _, value := range listSet {
		if value == nil {
			continue
		}
		splitedValue := strings.Split(strings.TrimSpace(value.(string)), " ")

		if len(splitedValue) > 1 { // if its a text with more than 1 word, like full name
			// specially handle the first term
			firstTerm := splitedValue[0]
			firstTermsUnique[firstTerm] = true

			// get the other terms that will be combined
			for k := range splitedValue {
				if k == 0 {
					continue
				}
				remainTermsUnique[splitedValue[k]] = true
			}

		} else { // if its a text with only 1 word
			firstTermsUnique[splitedValue[0]] = true
		}
	}
	for k := range firstTermsUnique {
		firstTerms = append(firstTerms, k)
	}
	for k := range remainTermsUnique {
		remainTerms = append(remainTerms, k)
	}

	return listSet
}
