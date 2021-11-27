package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	fmt.Print(os.Getenv("HOST"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Print("Error ", err)
	} else {
		fmt.Print(db)

		user := Agent{Name: "Omen", Abilities: []string{"Shrouded Step", "Dark Clouds", "Paranoia"}, Ult: "Teleport"}

		r.HandleFunc("/hello/{name}", helloHandler)
		agentRouter := r.PathPrefix("/agents").Subrouter()
		agentRouter.HandleFunc("/{class}", agentHandler)

		http.ListenAndServe(":5000", r)
	}
}
