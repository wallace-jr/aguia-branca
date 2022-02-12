package main

import (
	"github.com/wallace-jr/aguia-branca/database"
	"github.com/wallace-jr/aguia-branca/server"
)

func main(){
	database.StartDB()

	server := server.NewServer()
	
	server.Run()
}