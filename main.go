package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var zineTmpl = template.Must(template.ParseFiles("templates/index.html"))

func zineHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := zineTmpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", zineHandler)
	http.ListenAndServe(":"+port, mux)
}
