package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	agents "github.com/sajeelwaien/gopro/Agents"
	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/migrations"
	"github.com/sajeelwaien/gopro/schemas"
)

type body struct {
	query string
}

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
	// fmt.Printf("QUERY2 {}", query)
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result. {}", result.Errors)
	}
	return result
}

func main() {
	r := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	err := godotenv.Load(".env")

	database.InitDB()

	migrations.Migrate()

	if err != nil {
		fmt.Println("Error ", err)
	} else {
		r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
			var query map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&query)
			if err != nil {
				fmt.Printf("Error decoding request body", err)
			}
			// fmt.Printf("QUERY: {}", query["query"].(string), err)
			result := executeQuery(query["query"].(string), schemas.AgentSchema)
			json.NewEncoder(w).Encode(result)
		}).Methods("GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH")

		r.HandleFunc("/hello/{name}", helloHandler)
		agentRouter := r.PathPrefix("/agents").Subrouter()
		// agentRouter.HandleFunc("/{class}", agentHandler)
		agentRouter.HandleFunc("/add", agents.AddAgent).Methods("POST")
		fmt.Print("Running on Port 5000")
		http.ListenAndServe(":5000", handlers.CORS(header, methods, origins)(r))
	}
}
