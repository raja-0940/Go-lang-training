package DAL

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Id     int
	Name   string
	Isbn   string
	Author string
	Pages  int //`json:",omitempty"`
}

func InsertBook(book Book) {

	session := Connect()

	collection := session.Database("webinardb").Collection("books") //

	result, err := collection.InsertOne(context.TODO(), book)

	log.Println(result)
	log.Println(err)

}

func GetBookByIsbn(isbn string) Book {

	var book Book
	session := Connect()

	collection := session.Database("webinardb").Collection("books") // from a given db access books collection(Table)

	fmt.Println("After connection...")
	filter := bson.M{"isbn": isbn} // select * from books where isbn =100

	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		fmt.Println(err)

	}
	return book
}

func GetBooks() []Book {

	var books []Book

	// //var book = Book{Id:10}

	// books = append(books, Book{Id: 10, Name: "JS"})
	// books = append(books, Book{Id: 11, Name: "GoLang"})
	// books = append(books, Book{Id: 12, Name: "Node.js"})

	// return books

	session := Connect()

	collection := session.Database("webinardb").Collection("books") // from a given db access books collection(Table)

	cur, _ := collection.Find(context.TODO(), bson.M{}) // emtpy filter (select * from books)

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) { // Iterating through records

		var book Book
		//COnvert mongo DB record into book object
		err := cur.Decode(&book) // convert book document(record) to Book struct object

		if err != nil {
			log.Fatal(err)
		}

		//Add book object to Slice
		books = append(books, book)
	}
	return books

}

func DeleteBookByIsbn(isbn string) {
	//var book Book
	session := Connect()

	collection := session.Database("webinardb").Collection("books") // from a given db access books collection(Table)

	filter := bson.M{"isbn": isbn}

	result, err := collection.DeleteOne(context.TODO(), filter)

	fmt.Println(result, err)

}

func Transaction() {

	session := Connect()
	db := session.Database("webinardb")
	defer db.Client().Disconnect(context.TODO())
	col := db.Collection("books")

	// transaction
	db.Client().UseSession(context.TODO(), func(sessionContext mongo.SessionContext) error {

		err := sessionContext.StartTransaction()
		fmt.Println("Transaction started..")
		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err1 := col.InsertOne(sessionContext, bson.M{"name": "ABC", "author": "XYZ"})
		if err1 != nil {
			fmt.Println(err1)
			return err
		}

		_, err2 := col.InsertOne(sessionContext, bson.M{"name": "PQR", "author": "MNP"})

		if err1 != nil || err2 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
			sessionContext.AbortTransaction(sessionContext)

			return err
		} else {
			sessionContext.CommitTransaction(sessionContext)
		}
		return nil
	})
}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connected...")
	return client
}
