package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	agents "github.com/sajeelwaien/gopro/Agents"
	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/migrations"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(writer, "Hello From %s", vars["name"])
}

func agentHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	category := vars["class"]

	fmt.Fprintf(writer, "%s", category)
}

func main() {
	r := mux.NewRouter()

	err := godotenv.Load(".env")

	database.InitDB()

	migrations.Migrate()

	if err != nil {
		fmt.Println("Error ", err)
	} else {
		r.HandleFunc("/hello/{name}", helloHandler)
		agentRouter := r.PathPrefix("/agents").Subrouter()
		// agentRouter.HandleFunc("/{class}", agentHandler)
		agentRouter.HandleFunc("/add", agents.AddAgent).Methods("POST")
		fmt.Print("Running on Port 5000")
		http.ListenAndServe(":5000", r)
	}
}
