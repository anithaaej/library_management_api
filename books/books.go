package books

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anithaa19/bms/structs"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "goBasicsNewBMS"
const colName = "books"

var collection *mongo.Collection

func init() {
	// Client options
	clientOption := options.Client().ApplyURI(connectionString)

	// MongoDB Connection
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkNilError(err)

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")

}

// Mongo Helpers File

// Insert record into Collection
func insertRecord(book structs.Book) string {
	_, err := collection.InsertOne(context.Background(), book)
	checkNilError(err)
	msg := "Record Inserted Successfully"
	return msg
}

// Update record from collection
func updateRecord(bookId string, book structs.Book) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"rentedcopies": book.RentedCopies}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkNilError(err)

	fmt.Println("Record Updated Successfully, ", result.ModifiedCount)
}

// Delete a Record
func deleteOneRecord(bookId string) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	checkNilError(err)
	fmt.Printf("Deleted %d Record Successfully\n", result)
}

// Delete Multiple Records
func deleteManyRecords() {
	// id, _ := primitive.ObjectIDFromHex(bookId)
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	checkNilError(err)
	fmt.Printf("Deleted %d Records Successfully.", result.DeletedCount)
}

// Fetch All Records from DB
func getAllRecords() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkNilError(err)

	var books []primitive.M

	for cursor.Next(context.Background()) {
		var book bson.M
		err := cursor.Decode(&book)
		checkNilError(err)

		books = append(books, book)
	}

	defer cursor.Close(context.Background())
	return books
}

func getRecordbyId(getId string) bson.M {

	var (
		book bson.M
		id   primitive.ObjectID
		err  error
	)
	id, err = primitive.ObjectIDFromHex(getId)
	checkNilError(err)

	er := collection.FindOne(context.Background(), bson.D{{"_id", id}}).Decode(&book)
	checkNilError(er)

	return book
}

// Actual Controllers Functions

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	// var response structs.Response
	allMyBooks := getAllRecords()

	// response.Status = "200"
	// response.Message = "Retrieved all Books successfully"
	// response.Data = allMyBooks

	json.NewEncoder(w).Encode(allMyBooks)
}

func GetBookbyId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	// var response structs.Response
	params := mux.Vars(r)
	getBook := getRecordbyId(params["id"])

	// response.Status = "200"
	// response.Message = "Retrieved all Books successfully"
	// response.Data = allMyBooks

	json.NewEncoder(w).Encode(getBook)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var (
		book structs.Book
		// response structs.Response
	)
	_ = json.NewDecoder(r.Body).Decode(&book)
	respMsg := insertRecord(book)
	// response.Status = "201"
	// response.Message = respMsg
	// json.NewEncoder(w).Encode(response)

	json.NewEncoder(w).Encode(respMsg)
}

func UpdateBookData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	var book structs.Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	updateRecord(params["id"], book)
	json.NewEncoder(w).Encode("Data Updated Successfully")
}

func DeleteBookData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode("Data Deleted Successfully")
}

func DeleteAllBookData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteManyRecords()
	json.NewEncoder(w).Encode("Data Deleted Successfully")
}

// Errors Functions

func checkNilError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
