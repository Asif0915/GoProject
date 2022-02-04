package PersonControllers 
import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AsifIITR/mongodb-go1/Models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
var client *mongo.Client
func CreatePerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var person AllConst.Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("Asifdatabase").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person) //fmt.Println(result)
	json.NewEncoder(response).Encode(result)       //NewEncoder response ko read krega aur fir usko result me encode krega
}
func UpdatePerson(response http.ResponseWriter, request *http.Request){
	ctx, _:= context.WithTimeout(context.Background(), 10*time.Second)
	err=client.Connect(ctx)
	if err!=nil {
		log.Fatal(err)
	}
       peopleCollection:=client.Database("Asifdatabase").Collection(people)
       id,_=primitive.ObjectIDFromHex("60b7978c2da0323d474859b5")
       result,err:=peopleCollection.UpdateOne(
	       ctx,
	       bson.M{"_id":id},
	       bson.D{ {"$set",bson.D{{"firstname":"NewAsif"}}}
   },
       )
       if err!=nil{
	    log.fatal(err)   
       }
       fmt.Printf("updated %v \n",result.ModifiedCount)
}
func DeletePerson(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.Database("Asifdatabase").Collection("people")
	id, _ := primitive.ObjectIDFromHex(params["id"])
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}
func GetPerson(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var people []Models.Person
	collection := client.Database("Asifdatabase").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person AllConst.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
	//	json.Unmarshal([]byte(), &people)
}
func GetPeopleEndPoint(response http.ResponseWriter, request *http.Request) { //querry for a particular entry
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Models.Person
	collection := client.Database("Asifdatabase").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, AllConst.Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)
}
// func aggregate(response http.ResponseWriter, request *http.Request) []bson.M {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer client.Disconnect(ctx)
// 	borrowcollection := client.Database("Asifdatabase").Collection("people")
// 	lookupStage := bson.D{{"$lookup", bson.D{primitive.E{"from", "prople"}, {"localField", "Person"}, {"foreignField", "_id"}, {"as", "Person"}}}}
// 	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$Person"}, {"preserveNullAndEmptyArrays", false}}}}

// 	showLoadedStructCursor, err := borrowcollection.Aggregate(ctx, mongo.Pipeline{lookupStage, unwindStage})
// 	if err != nil {
// 		panic(err)
// 	}
// 	var showsLoadedStruct []bson.M
// 	if err = showLoadedStructCursor.All(ctx, &showsLoadedStruct); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(showsLoadedStruct)
// 	return showsLoadedStruct
// }

// func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
//         response.Header().Add("content-type", "application/json")
//         var person Person
//         json.NewDecoder(request.Body).Decode(&person)                       //fmt.Println(person) fmt.Println(client)
//         collection := client.Database("Asifdatabase").Collection("people")  //fmt.Println(collection)
//         ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) //fmt.Println(ctx)

//         result, _ := collection.InsertOne(ctx, person) //fmt.Println(result)
//         json.NewEncoder(response).Encode(result)
// }
