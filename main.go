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
	"github.com/rs/cors"
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



	



	c := cors.New(cors.Options{
		
	AllowCredentials: true,
	AllowedMethods: []string{"GET","POST", "OPTIONS","PUT","DELETE"},
    AllowedOrigins: []string{"*"},
    AllowedHeaders: []string{"Content-Type","Authorization","Bearer","Bearer ","content-type","authorization","Origin","Accept"},
    OptionsPassthrough: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	})
	handler := c.Handler(r)
	fmt.Println("Server is running on port: ", strings.Split(":"+p, ":")[1])
	log.Fatal(http.ListenAndServe(":"+p, handler))
}
