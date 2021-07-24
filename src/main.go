package main

import (
	"encoding/json"
	"fmt"
	"github.com/coopstools-homebrew/shadows-of-the-forgotten/persistance"
	"github.com/rs/cors"
	"net/http"
	"os"
)

var table = persistance.Connect()

func main() {
	prefix := ""
	if len(os.Args) > 2 {
		prefix = "/" + os.Args[2]
	}

	mux := http.NewServeMux()
	mux.HandleFunc(prefix + "/", GetData)
	handler := logRequestHandler(mux)
	handler = cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:*",
			"https://home.coopstools.com",
			"http://home.coopstools.com",
		},
	}).Handler(handler)
	addr := ":" + os.Args[1]
	fmt.Println(addr)
	http.ListenAndServe(addr, handler)
}

func GetData(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(table.GetAll())
	if err != nil {
		w.WriteHeader(500)
		fmt.Printf("\nServer error: %+v", err)
		fmt.Fprint(w, "Internal server error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	count, _ := w.Write(data)
	fmt.Printf("\n%d bytes returned", count)
}

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		uri := r.URL.String()
		method := r.Method
		fmt.Printf("\n%v: %v", method, uri)
	}
	return http.HandlerFunc(fn)
}
