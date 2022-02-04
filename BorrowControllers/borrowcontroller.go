// package BorrowControllers

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/AsifIITR/mongodb-go1/Models"
// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var client *mongo.Client

// func CreateBorrowBook(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Add("content-type", "application/json")
// 	var borrow Models.Borrow
// 	json.NewDecoder(request.Body).Decode(&borrow)
// 	collection := client.Database("Asifdatabase").Collection("borrow")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	result, _ := collection.InsertOne(ctx, borrow) //fmt.Println(result)
// 	json.NewEncoder(response).Encode(result)       //NewEncoder response ko read krega aur fir usko result me encode krega
// 	fmt.Println("Create is called")
// }

// func DeleteBorrowBook(response http.ResponseWriter, request *http.Request) {
// 	params := mux.Vars(request)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	collection := client.Database("Asifdatabase").Collection("borrow")
// 	id, _ := primitive.ObjectIDFromHex(params["id"])
// 	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(result)
// }

// func GetBorrowBorrowBook(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Add("content-type", "application/json")
// 	var book []Models.Person
// 	collection := client.Database("Asifdatabase").Collection("book")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	cursor, err := collection.Find(ctx, bson.M{})
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
// 		return
// 	}
// 	defer cursor.Close(ctx)
// 	for cursor.Next(ctx) {
// 		var book1 Models.Person
// 		cursor.Decode(&book1)
// 		book = append(book, book1)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
// 		return
// 	}
// 	json.NewEncoder(response).Encode(book)
// }

// func Updatemethod(response http.ResponseWriter, request *http.Request) {
// 	var people AllConst.Person
// 	err := json.NewDecoder(request.Body).Decode(&book)
// 	params := mux.Vars(request)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	collection := client.Database("Asifdatabase").Collection("book")
// 	id, _ := primitive.ObjectIDFromHex(params["id"])
// 	if "_id"=id{
//edited
// 	}
// 	result, err := collection.UpdateOne(ctx, bson.M{"_id": id})
// 	if err != nil {
// 		respondWithError(response, http.StatusBadRequest, err.Error())
// 	} else {
// 		fmt.Println("_id: ", book.B_ID)
// 		fmt.Println("bname: ", book.Bname)
// 		fmt.Println("bprice: ", book.bprice)
// 		respondWithJSON(response, http.StatusOK, book)
// 	}
// }
// func respondWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
// 	result, _ := json.Marshal(data)
// 	response.Header().Set("Content-Type", "application/json")
// 	response.WriteHeader(statusCode)
// 	response.Write(result)
// }

// func respondWithError(response http.ResponseWriter, statusCode int, msg string) {
// 	respondWithJSON(response, statusCode, map[string]string{"error": msg})
// }
