package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	han "github.com/tusharhow/handlers"
)

func main() {
	// Load .env file to use the environment variable
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := mux.NewRouter()
	p := os.Getenv("PORT")

	r.HandleFunc("/", han.GetFormat).Methods("GET")
	r.HandleFunc("/sendmail", han.SendMail).Methods("POST")

	fmt.Println("Server is running on port: ", strings.Split(":"+p, ":")[1])

	log.Fatal(http.ListenAndServe(":"+p, r))

}
