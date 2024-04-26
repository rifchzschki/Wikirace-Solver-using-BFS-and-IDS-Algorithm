package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"wikirace/algorithm"
)


var (
	
	path map[string]string
)



func mergeMaps(map1, map2 map[string]string) map[string]string {
    mergedMap := make(map[string]string)
    for key, value := range map1 {
        mergedMap[key] = value
    }
    for key, value := range map2 {
        mergedMap[key] = value
    }

    return mergedMap
}



func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	var results []string
	var tempPath map[string]string
	query = strings.ReplaceAll(query, " ", "_")
	results,tempPath,_= algorithm.FetchSuggestions(query)
	path = mergeMaps(tempPath, path)
	response, _:= json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}


func handlerProcess(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
		var filepath = "Frontend/src/wikirace.html"
        var tmpl = template.Must(template.New("result").ParseFiles(filepath))

        if err := r.ParseForm(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        var start = r.FormValue("start")
        var startURL = path[start]
        var target = r.FormValue("target")
        var targetURL = path[target]
        var algo = r.Form.Get("algo")
		path, checked, solutionLength, duration:= algorithm.Test(algo, startURL, targetURL)   
		if(len(path)<=0){
			path = append(path, "Rute tidak ditemukan") 
		}   
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
	var filepath = "Frontend/src/wikirace.html"
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
			http.FileServer(http.Dir("Frontend"))))
	
	var address = "0.0.0.0:8080"
	// var address = "localhost:8080"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}



func main() {
		startServer()
}


