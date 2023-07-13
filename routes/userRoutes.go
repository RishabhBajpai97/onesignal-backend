package routes

import (
	"context"
	"onesignal-backend/controllers"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

func Router(collection *mongo.Collection, ctx context.Context) *httprouter.Router {
	routes := httprouter.New()
	routes.GET("/user/getallusers", controllers.GetAllUsers(collection, ctx))
	routes.POST("/login", controllers.Login(collection, ctx))
	routes.POST("/signup", controllers.SignUp(collection, ctx))
	routes.POST("/user/notifications", controllers.SendNotification())
	return routes
}
