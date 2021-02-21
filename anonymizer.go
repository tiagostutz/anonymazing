package main

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
	return listSet
}
