package main

import (
	"authentification/database"
	"log"
	"net/http"
)

func main() {
	// connect to DB
	conn := database.ConnectToDB()
	if conn == nil {
		log.Fatal("Cannot connect to DB")
	}

	app := NewApp(conn)

	log.Println("Starting Authentication service on port 80")

	//define server
	srv := &http.Server{
		Addr:    ":80",
		Handler: app.routes(),
	}

	//start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
