package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/migrations"
	"github.com/sajeelwaien/gopro/models"
	"github.com/sajeelwaien/gopro/proto"
	"google.golang.org/grpc"
)

type body struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
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

func executeQuery(query string, schema graphql.Schema, variables map[string]interface{}) *graphql.Result {
	// fmt.Printf("QUERY2 %s", query)
	fmt.Println("Variables", variables)
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		RootObject:     nil,
		VariableValues: variables,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result. %+v\n", result.Errors)
	}
	return result
}

// func main() {
// 	r := mux.NewRouter()
// 	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
// 	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
// 	origins := handlers.AllowedOrigins([]string{"*"})

// 	err := godotenv.Load(".env")

// 	database.InitDB()

// 	migrations.Migrate()

// 	if err != nil {
// 		fmt.Println("Error ", err)
// 	} else {
// 		r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
// 			var requestBody body

// 			errBody := json.NewDecoder(r.Body).Decode(&requestBody)

// 			if errBody != nil {
// 				fmt.Printf("Error decoding request body: %+v\n", errBody)
// 			}

// 			result := executeQuery(requestBody.Query, schemas.AgentSchema, requestBody.Variables)
// 			fmt.Printf("RESULT: %+v", result)
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(&result)
// 		}).Methods("GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH")

// 		r.HandleFunc("/hello/{name}", helloHandler)
// 		agentRouter := r.PathPrefix("/agents").Subrouter()
// 		agentRouter.HandleFunc("/add", agents.AddAgent).Methods("POST")
// 		fmt.Print("Running on Port 5000")
// 		http.ListenAndServe(":5000", handlers.CORS(header, methods, origins)(r))
// 	}
// }

type Server struct {
	proto.UnimplementedSelectAgentServer
}

func (s *Server) SelectAgent(ctx context.Context, message *proto.AgentName) (*proto.SelectedAgentReply, error) {
	log.Printf("Message from client: %s", message.Name)

	var agent models.Agent
	result := database.DBCon.Where("name = ?", message.Name).First(&agent)

	log.Printf("===> %+v", &agent)

	if result.Error != nil {
		return nil, result.Error
	}

	return &proto.SelectedAgentReply{
		Message: agent.Class,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Cannot listen")
	}

	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		fmt.Println("Error ", err)
	}

	database.InitDB()

	migrations.Migrate()

	s := Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterSelectAgentServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to ser grpc server")
	}
}
