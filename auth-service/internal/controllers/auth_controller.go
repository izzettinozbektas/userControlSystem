package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/izzettinozbektas/userControlSystem/auth-service/internal/config"
	"github.com/izzettinozbektas/userControlSystem/auth-service/internal/models"
	"github.com/izzettinozbektas/userControlSystem/auth-service/internal/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = config.GetCollection(config.ConnectDB(), "users")

func Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Fail(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	// Check if user exists
	count, _ := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if count > 0 {
		response.Fail(w, http.StatusConflict, "User already exists", nil)
		return
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, "Error hashing password", err.Error())
		return
	}

	user.Password = string(hashed)

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		response.Fail(w, http.StatusInternalServerError, "Database error", err.Error())
		return
	}

	// Mongo'dan dönen ObjectID'yi User struct'a ata
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""

	response.Success(w, "Kullanıcı oluşturuldu", user)

}
