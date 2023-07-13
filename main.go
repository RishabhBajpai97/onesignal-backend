package main

import (
	"fmt"
	"log"
	"net/http"
	"onesignal-backend/db"
	"onesignal-backend/routes"
)

func main() {
	collection, ctx, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start main")
	routes := routes.Router(collection, ctx)
	server := http.Server{Addr: "localhost:3000", Handler: routes}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
