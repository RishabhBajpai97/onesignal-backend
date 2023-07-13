package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"onesignal-backend/model"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(collection *mongo.Collection, ctx context.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var requestBody model.RequestBody
		var responseBody model.RequestBody
		json.NewDecoder(r.Body).Decode(&requestBody)
		fmt.Println(requestBody)
		err := collection.FindOne(ctx, bson.D{primitive.E{Key: "email", Value: requestBody.Email}}).Decode(&responseBody)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode("User not found")
				return
			}
		}
		fmt.Println("Found it")
		json.NewEncoder(w).Encode(responseBody)

	}
}

func SignUp(collection *mongo.Collection, ctx context.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var requestBody model.RequestBody
		var responseBody model.RequestBody
		json.NewDecoder(r.Body).Decode(&requestBody)
		err := collection.FindOne(ctx, bson.D{primitive.E{Key: "email", Value: requestBody.Email}}).Decode(map[string]string{})
		if err == mongo.ErrNoDocuments {
			id, er := collection.InsertOne(ctx, requestBody)
			if er != nil {
				log.Fatal(er)
			}
			er = collection.FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: id.InsertedID}}).Decode(&responseBody)
			if er != nil {
				log.Fatal(er)
			}
			json.NewEncoder(w).Encode(responseBody)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "User Already Exists"})

	}
}

func GetAllUsers(collection *mongo.Collection, ctx context.Context) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var users []model.User
		cursor, err := collection.Find(ctx, bson.D{{}})
		cursor.All(ctx, &users)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"users": users})
		defer cursor.Close(ctx)
	}
}
