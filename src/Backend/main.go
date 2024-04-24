package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
    var message = "Welcome"
    w.Write([]byte(message))
	}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	var filepath = filepath.Join("../", "Frontend", "html_project", "wikirace.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func handlerHello(w http.ResponseWriter, r *http.Request) {
    var message = "Hello world!"
    w.Write([]byte(message))
}

func handlerData(w http.ResponseWriter, r *http.Request){
	// Membuat data yang akan dikirim ke frontend
	data := map[string]string{"message": "Hello from Go!"}

	// Mengubah data ke format JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengatur header HTTP
	w.Header().Set("Content-Type", "application/json")
	
	// Mengirimkan response ke frontend
	fmt.Fprintf(w, "%s", jsonData)
}


func main() {
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/Index", handlerIndex)
	http.HandleFunc("/Hello", handlerHello)
    http.HandleFunc("/Data", handlerData)
	
	http.Handle("/static/", 
        http.StripPrefix("/static/", 
            http.FileServer(http.Dir("../Frontend/html_project"))))

    // Menjalankan server di port 8080
	var address = "localhost:8080"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}