package main

import (
	"context"
	"fmt" // format
	"log"
	"net/http" //http module
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func homeRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home route")
}

func aboutRoute(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"About route")
}

func contactRoute(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Contact route")
}

func connectToDB() *mongo.Collection{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("golang").Collection("users")

	return collection
}

type User struct {
	ID string `bson:"id,omitempty"`
	Name string `bson:"name"`
	Email string `bson:"email"`
}

func createUser(collection *mongo.Collection,user User) error{
	_, err:= collection.InsertOne(context.Background(),user)
	return err
}

func findUserByID(collection *mongo.Collection, id string) (User, error) {
	var user User
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func updateUser(collection *mongo.Collection, id string, updatedUser User) error {
	_, err := collection.ReplaceOne(context.Background(), bson.M{"_id": id}, updatedUser)
	return err
}

func deleteUser(collection *mongo.Collection, id string) error {
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

// func main(){
// 	http.HandleFunc("/",homeRoute)
// 	http.HandleFunc("/about",aboutRoute)
// 	http.HandleFunc("/contact",contactRoute)
	
	
// 	fmt.Println("Server is running at http://localhost:8080")
// 	http.ListenAndServe(":8080",nil)
// }
func main() {
	collection := connectToDB()

	// Create a new user
	newUser := User{Name: "John Doe", Email: "johndoe@example.com"}
	err := createUser(collection, newUser)
	if err != nil {
		log.Fatal(err)
	}

	// Find user by ID
	user, err := findUserByID(collection, newUser.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found user:", user)

	// Update user
	user.Name = "Jane Doe"
	err = updateUser(collection, user.ID, user)
	if err != nil {
		log.Fatal(err)
	}

	// Delete user
	err = deleteUser(collection, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User deleted successfully")
}
