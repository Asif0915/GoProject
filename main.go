package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AsifIITR/mongodb-go1/Routers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	fmt.Println("starting the application..")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://ricon:ricon%4012345@data-con.m5z3f.mongodb.net/Data"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	client = c
	fmt.Println(client)
	router := mux.NewRouter()
	Routers.RegisterBookandPersonStoreRoutes(router)
	http.ListenAndServe(":12345", router)

}

//id corresponding Update people data  (60b7978c2da0323d474859b5)
