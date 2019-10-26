package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//
type Database struct {
	Db *mongo.Database
}

var db *Database

func Getdb() *Database {
	if db == nil {
		db = getDb()
	}
	return db
}

func getDb() *Database {
	db := Database{}
	// fmt.Println("get started with db")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	fmt.Printf("%T\n", cancel)
	// fmt.Println(cancel)
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, error := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if error != nil {
		log.Fatal(error)
	}

	error = client.Connect(ctx)
	if error != nil {
		log.Fatal(error)
	}
	error = client.Ping(context.TODO(), nil)
	fmt.Println("Database connected")
	database := client.Database("rule")
	db.Db = database
	return &db
}
