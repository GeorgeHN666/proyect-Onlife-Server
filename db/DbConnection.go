package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This Variable Its Gonna be Use To Connect To Our Database Anywhere
var DBConnection = connection()

// This Variable Its Responsable To take our Mongo Token Of our Database Along with the password
var clientOptions = options.Client().ApplyURI(" Here Goes The Database String ")

// This is The Function That Will Do The Connection With The Database
func connection() *mongo.Client {

	// Here we Create The Connection Passing The Context and The Client Options
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln("Failed To Create The Mongo Connection", err)
	}

	// Here we Make Sure That We Create The Connection With Mongo
	err = client.Ping(context.TODO(), nil)
	// Check Error
	if err != nil {
		log.Fatalln("Were Not Connected With Mongo Database")
	}

	// Here we Display A Message To the log that the connection is succesfull
	log.Println("Successful Connection With Mongo")

	// Here we Return The Pointer to Mongo Client
	return client

}

// This Function Will Allow us To Simply Check If The Connection With Mongo Its Good It's Gonna Be Exportable
func DBConnectionCheck() int {

	err := DBConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 1
	}

	return 0

	// If The Connection Its Good, Were Gonna Return 0 as an Integer
}
