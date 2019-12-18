package main

import (
	"context"
	"final-project/controller"
	"final-project/handler"
	"final-project/mongodb"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbAddress = "127.0.0.1:27017"
	dbName    = "MyDataBase"
)

func initMongoDB() *mongo.Database {

	var conf mongodb.Config
	conf.Address = dbAddress
	conf.Database = dbName
	db, err := mongodb.Init(context.Background(), &conf)
	if err != nil {
		log.Panicf("could not connect to the database: %v", err)
	}
	return db
}

func main() {
	db := initMongoDB()
	defer func() {
		if err := db.Client().Disconnect(context.Background()); err != nil {
			log.Printf("could not disconnect from database: %v", err)
		}
	}()

	employeeController := controller.Employee{
		DbEmployee: &mongodb.Employee{
			db.Collection("Employee"),
		},
	}

	departmentController := controller.Department{
		DbDepartment: &mongodb.Department{
			db.Collection("Department"),
		},
	}

	mux := handler.CreateMux(employeeController, departmentController)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
