package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var outputSQLFile string

func main() {
	fmt.Println("")

	fmt.Println("==** Anonymazing Your Data **==")
	fmt.Println("")

	flag.StringVar(&postgresConnectionString, "postgres-connection-string", "", "Postgres connection URI. Example: --postgres-connection-string=postgresql://postgres:123456@localhost:5432/my_database?sslmode=disable")
	flag.StringVar(&databaseTable, "database-table", "", "Database table to have the columns anonymized. Example: --database-table=person")
	flag.StringVar(&databaseColumns, "database-columns", "", "Comma separated list of columns to anonymize. Example: --database-columns=first_name,second_name,nickname. (Those columns must present on the table passed in --database-table arg)")
	flag.StringVar(&outputSQLFile, "output", "./anonymizer_script.sql", "File that will receive the SQL. Example: --out=./anonymizer_script.sql")
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

		os.Remove(outputSQLFile)
		for i := range anonymizedData {
			for columnName, columnValue := range anonymizedData[i] {
				writeFile(outputSQLFile, fmt.Sprintf("UPDATE %s SET %s='%s' WHERE %s='%s';", databaseTable, columnName, columnValue.(string), columnName, dataToBeAnonimized[i][columnName]))
			}
		}

	}
	fmt.Println("")
}

func writeFile(path string, content string) error {

	var accContent string = ""
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		accContent = ""
	}
	accContent = string(fileContent)
	if accContent != "" && strings.Contains(accContent, content) {
		return nil
	}

	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logrus.Warnf("Erro ao abrir o arquivo de log de erros de %s para escrever. Error: %s", path, err)
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(content + "\n"); err != nil {
		logrus.Warnf("Erro ao escrever texto no arquivo de log de erros de %s para escrever. Error: %s", path, err)
		return err
	}
	return nil
}
