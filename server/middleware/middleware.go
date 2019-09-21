package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb://localhost:27017"

// DB name
const dbName = "dbCTCI"

// Collection name
const collName = "exercises"

// Collection obj/instance
var collection *mongo.Collection

// Create connection with mongo
func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created")
}

// GetAllExs route
func GetAllExs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllExs()
	json.NewEncoder(w).Encode(payload)
}

// ExCheck route
func ExCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	exCheck(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// ExUncheck route
func ExUncheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	exUncheck(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// getAllExs from db
func getAllExs() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// exCheck from db
func exCheck(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// exUncheck from db
func exUncheck(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// Insert exercises in DB (be careful to not everride, i.e. call it ONLY ONCE)
func FillExercises(chapter int, quantity int) {
	records := 0
	exChapter := strconv.Itoa(chapter)
	var exercises []interface{}
	var exNumber string

	for i := 1; i <= quantity; i++ {
		exNumber = exChapter + "." + strconv.Itoa(i)
		exercises = append(exercises, bson.M{"number": exNumber, "status": false})
		records++
	}

	insertedResults, err := collection.InsertMany(context.Background(), exercises)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted ", len(insertedResults.InsertedIDs), " records")
}
