package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	agents "github.com/sajeelwaien/gopro/Agents"
	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/migrations"
	"github.com/sajeelwaien/gopro/schemas"
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

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result.")
	}
	return result
}

func main() {
	r := mux.NewRouter()

	err := godotenv.Load(".env")

	database.InitDB()

	migrations.Migrate()

	if err != nil {
		fmt.Println("Error ", err)
	} else {
		r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			result := executeQuery(r.URL.Query().Get("query"), schemas.AgentSchema)
			json.NewEncoder(w).Encode(result)
		})

		r.HandleFunc("/hello/{name}", helloHandler)
		agentRouter := r.PathPrefix("/agents").Subrouter()
		// agentRouter.HandleFunc("/{class}", agentHandler)
		agentRouter.HandleFunc("/add", agents.AddAgent).Methods("POST")
		fmt.Print("Running on Port 5000")
		http.ListenAndServe(":5000", r)
	}
}
