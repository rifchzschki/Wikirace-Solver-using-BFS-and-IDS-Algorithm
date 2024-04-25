package algorithm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// fetchSuggestions mengambil saran dari Wikipedia berdasarkan input.
func FetchSuggestions(input string) ([]string, map[string]string, error) {
	url := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=opensearch&limit=10&format=json&search=%s&origin=*", input)
	paths := make(map[string]string)
	resp, err := http.Get(url)
	if err != nil {
		return nil,nil, err
	}
	defer resp.Body.Close()

	var data []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil,nil, err
	}

	suggestions := make([]string, 0)
	if len(data) > 1 {
		for _, suggestion := range data[1].([]interface{}) {
			suggestions = append(suggestions, suggestion.(string))
			formattedSuffestionArticle := strings.ReplaceAll(suggestion.(string), " ", "_")
			fullSuggestArticleURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", formattedSuffestionArticle)
			paths[suggestion.(string)]=fullSuggestArticleURL
		}
	}

	return suggestions,paths, nil
}
// func printMap(m map[string]string) {
// 	for key, value := range m {
// 		fmt.Printf("Key: %s, Value: %s\n", key, value)
// 	}
// }

// func main() {
	// var mu sync.Mutex

// 	startArticle := "I"
// 	// targetArticle := "Penis"

// 	startSuggestions, pathStart,err := fetchSuggestions(startArticle)
// 	if err != nil {
// 		log.Fatalf("Error fetching start suggestions: %v", err)
// 	}

// 	// targetSuggestions,pathTarget, err := fetchSuggestions(targetArticle)
// 	if err != nil {
// 		log.Fatalf("Error fetching target suggestions: %v", err)
// 	}

// 	// Handle the spaces in article titles
// 	formattedStartArticle := strings.ReplaceAll(startArticle, " ", "_")
// 	// formattedTargetArticle := strings.ReplaceAll(targetArticle, " ", "_")

// 	// Construct full URLs
// 	fullStartArticleURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", formattedStartArticle)
// 	// fullTargetArticleURL := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", formattedTargetArticle)

// 	// Print suggestions and URLs
// 	// for _, sus := range (startSuggestions){
// 	// 	fmt.Println("Start Suggestions:", sus)
// 	// }
// 	fmt.Println("Start Suggestions:", startSuggestions)
// 	printMap(pathStart)
// 	// fmt.Println("Target Suggestions:", targetSuggestions)
// 	// printMap(pathTarget)
// 	fmt.Println("Full Start Article URL:", fullStartArticleURL)
// 	// fmt.Println("Full Target Article URL:", fullTargetArticleURL)
// }
