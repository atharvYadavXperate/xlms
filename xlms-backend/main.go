package main

import (
	db "github.com/atharvYadavXperate/xlms/database"
)

func main() {
	db.ConnectDb()
	defer db.CloseConnection()
}
