package main

import (
	"root/root/database"
	"root/root/internal/transport/rest"

	_ "github.com/lib/pq"
)

func main() {
	database.InitDB()
	defer database.CloseDB()
	rest.HandleFunction()
}
