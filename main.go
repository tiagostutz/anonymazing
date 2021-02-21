package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("")

	fmt.Println("==** Anonymazing Your Data **==")
	fmt.Println("")

	flag.StringVar(&postgresConnectionString, "postgres-connection-string", "", "Postgres connection URI. Example: --postgres-connection-string=postgresql://postgres:123456@localhost:5432/my_database?sslmode=disable")
	flag.StringVar(&databaseTable, "database-table", "", "Database table to have the columns anonymized. Example: --database-table=person")
	flag.StringVar(&databaseColumns, "database-columns", "", "Comma separated list of columns to anonymize. Example: --database-columns=first_name,second_name,nickname. (Those columns must present on the table passed in --database-table arg)")
	flag.Parse()

	if postgresConnectionString == "" || databaseTable == "" || databaseColumns == "" {
		colorRed := "\033[31m"
		colorReset := "\033[0m"
		fmt.Println(string(colorRed))
		fmt.Println("Error!")
		fmt.Println("You must specify the following args:", string(colorReset))
		fmt.Println("--postgres-connection-string")
		fmt.Println("--database-table")
		fmt.Println("--database-columns")
		fmt.Println("")
		fmt.Println("Run `anonymazing --help` for more information")
		fmt.Println("")
	} else {
		fmt.Printf("Connecting to %s\n", postgresConnectionString)

		fmt.Printf("Anonymizing columns `%s` from table `%s`\n", databaseColumns, databaseTable)
		dataToBeAnonimized, err := readDatabaseData()
		if err != nil {
			panic(err)
		}

		anonymizedData, err := anonymizeResultSet(dataToBeAnonimized)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v", anonymizedData)

	}
	fmt.Println("")
}
