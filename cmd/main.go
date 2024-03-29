package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"tiagoncardoso/gotostudy/adapter/repository"
	"tiagoncardoso/gotostudy/usecase/process_transaction"
)

func main()  {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewTransactionRepositoryDb(db)
	ucase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput{
		ID:				"1",
		AccountID: 		"1",
		Amount: 		0,
	}

	output, err := ucase.Execute(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))
}
