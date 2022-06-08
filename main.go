package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	han "github.com/tusharhow/handlers"
)

func main() {

	r := mux.NewRouter()
	p := ":8080"

	r.HandleFunc("/sendmail", han.SendMail).Methods("POST")

	fmt.Println("Server is running on port: ", strings.Split(p, ":")[1])

	// han.EmailDatas = append(han.EmailDatas, han.EmailData{ID: "33", EmailTo: "kaka400068@gmail.com", EmailFrom: "kaka400069@gmail.com", EmailSub:"test",EmailBody:"testify",EmailPassword:"greenthree135790"})

	log.Fatal(http.ListenAndServe(p, r))

}
