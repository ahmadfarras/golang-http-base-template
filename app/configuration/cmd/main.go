package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting App..")
	StartHttpApp()
}
