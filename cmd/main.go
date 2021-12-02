package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ThaSMorato/go-escalavel/adapter/repository"
	"github.com/ThaSMorato/go-escalavel/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
		
	}

	repo := repository.NewTransactionRepositoryDb(db)

	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput{
		ID: "3",
		AccountID: "1",
		Amount: 0,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		fmt.Println(err)
	}

	outputJson, _ := json.Marshal(output)

	fmt.Println(string(outputJson))
}