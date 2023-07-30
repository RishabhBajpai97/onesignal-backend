package main

import (
	"fmt"
	"log"
	"net/http"
	"onesignal-backend/db"
	"onesignal-backend/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	collection, ctx, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	port:=os.Getenv("PORT")

	fmt.Println("Start main")
	routes := routes.Router(collection, ctx)
	server := http.Server{Addr: "0.0.0.0:"+port, Handler: routes}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
