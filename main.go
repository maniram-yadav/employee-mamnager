package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"employee-manager/backend/controllers"
	"employee-manager/backend/database"
	"employee-manager/backend/service"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Employee Manager")

	database.DbConn()

	Serve(4040)

}

func Serve(port int) {

	dbService := service.DbService{}
	controller := controllers.InitController(dbService)
	r := controllers.Init(controller)

	http.Handle("/", r)

	server := newServer(":"+strconv.Itoa(port), r)
	log.Printf("Starting server on localhost:%d", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newServer(addr string, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
}
