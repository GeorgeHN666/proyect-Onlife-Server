package main

import (
	"log"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/handlers"
)

func main() {
	// Here we Check If We Are Connected With The Mongo Database
	if db.DBConnectionCheck() == 1 {
		log.Fatalln("Fail to Connect With Mongo Database")
	}

	// Finally We Put The server to listen and serve

	handlers.Handler()

}
