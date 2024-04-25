package algorithm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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

