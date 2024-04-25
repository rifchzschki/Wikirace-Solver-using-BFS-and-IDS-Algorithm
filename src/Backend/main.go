package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"wikirace/algorithm"
)

// Data dummy untuk autocomplete
var data = []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape"}

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	var results []string
	for _, item := range data {
		fmt.Println(item)
		if strings.Contains(item, query) {
			results = append(results, item)
			fmt.Println("mashokk")
		}
		fmt.Println("hitung")
	}
	
	fmt.Print(results)
	response, _:= json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}


func handlerProcess(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
		var filepath = filepath.Join("../", "Frontend", "src", "wikirace.html")
        var tmpl = template.Must(template.New("result").ParseFiles(filepath))

        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		
		// algoritma pengambilan data hasil pencarian
        var startURL = "https://en.wikipedia.org/wiki/" + r.FormValue("start")
        var targetURL = "https://en.wikipedia.org/wiki/" + r.FormValue("target")
        var algo = r.Form.Get("algo")
		fmt.Println("Nilai algo: ", algo)

		path, checked, solutionLength, duration:= algorithm.Test(algo, startURL, targetURL)

        var data = map[string]interface{}{
			"start":           startURL,
			"target":          targetURL,
			"path":            path,
			"checked":         checked,
			"solutionLength":  solutionLength,
			"duration":        duration,
		}

        if err := tmpl.Execute(w, data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
	var filepath = filepath.Join("../", "Frontend", "src", "wikirace.html")
	if r.Method == "GET" {
        var tmpl = template.Must(template.New("form").ParseFiles(filepath))
        var err = tmpl.Execute(w, nil)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}


func startServer(){
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/process", handlerProcess)
	http.HandleFunc("/autocomplete", autocompleteHandler)
	
	http.Handle("/static/", 
		http.StripPrefix("/static/", 
			http.FileServer(http.Dir("../Frontend/"))))
	
	// Menjalankan server di port 8080
	var address = "localhost:8080"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}



func main() {
	var command string

	fmt.Println("Mau ngapain? (\"start\" untuk start server):")
	fmt.Println("1. \"start\" untuk start server")
	fmt.Println("2. \"test\" untuk testing algoritma:")
	fmt.Print("Jawaban: ")
    fmt.Scanln(&command)
	
	if (strings.EqualFold(command, "start")){
		startServer()
	}else if (strings.EqualFold(command, "test")){	
		var algo string
		fmt.Println("Pilih algoritma:")
		fmt.Scanln(&algo)
		startURL := "https://en.wikipedia.org/wiki/Tennis"
		targetURL := "https://en.wikipedia.org/wiki/Stephen_Curry"
		path, checked, solutionLength, duration := algorithm.Test(algo, startURL, targetURL)
		fmt.Println(path)
		fmt.Println(checked)
		fmt.Println(solutionLength)
		fmt.Println(duration)
	}
}
