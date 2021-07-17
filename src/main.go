package main

import (
	"encoding/json"
	"fmt"
	"github.com/coopstools-homebrew/users_api/persistance"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Id 		int64
	Name 	string
}

var table = persistance.Connect()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/person/", GetData)
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*"},
	}).Handler(mux)
	addr := ":" + os.Args[1]
	println(addr)
	http.ListenAndServe(addr, handler)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(table.GetAll())
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Server error: %+v", err)
		fmt.Fprint(w, "Internal server error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	count, _ := w.Write(data)
	log.Printf("%d bytes returned", count)
}
