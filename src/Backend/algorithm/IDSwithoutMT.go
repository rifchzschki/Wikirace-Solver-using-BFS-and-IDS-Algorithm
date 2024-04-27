package algorithm

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Function for run DFS
// Akan melakukan proses rekursi hingga ditemukan solusi atau hingga kemungkinan solusi habis
func DFS(currentURL, targetURL string, depth int, visited map[string]bool, paths map[string]string, articlesChecked *int) bool {
	if currentURL == targetURL {
        return true
    }
    if depth == 0 {
		return false
    }
	
	visited[currentURL] = true
	
	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", currentURL, err)
		return false
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		fmt.Printf("Unexpected status code %d for URL %s\n", resp.StatusCode, currentURL)
		return false
	}
	
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing document from URL %s: %v\n", currentURL, err)
		return false
	}
	
	if (isRedirect(doc)){
		return false
	}
	
	*articlesChecked ++

	found := false
	doc.Find("div#bodyContent").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link,":") {
				fullURL := "https://en.wikipedia.org" + link
				if !visited[fullURL] && !found {
					*articlesChecked ++
					paths[fullURL] = currentURL
					found = DFS(fullURL, targetURL, depth-1, visited, paths, articlesChecked)
				}
			}
		})
	})
	return found
}

//  Function for run IDS algorithm
// Akan menginisiasi proses ids dan memanggil proses dfs dengan limitasi jumlah depth
// akan mengeluarkan rute, jumlah artikel yang diperiksa, jumlah artikel solusi, durasi pencarian jika hasil ditemukan
// akan mengeluarkan nil, jumlah artikel yang diperiksa, 0, durasi pencarian
func IDS(startURL, targetURL string) ([] string, int, int, string) {
	startTime := time.Now()
	depth := 1
	paths := make(map[string]string)
	found := false
	checked :=0
	
	for !found {
		visited := make(map[string]bool)
		found = DFS(startURL, targetURL, depth, visited, paths, &checked)
		if found {
            // Path found
			path := []string{targetURL}
			for targetURL != startURL {
                targetURL = paths[targetURL]
				path = append([]string{targetURL}, path...)
			}
			duration := time.Since(startTime)
			durationMS := duration.Milliseconds()
			return path, checked, len(path), strconv.Itoa(int(durationMS)) + "ms"
		}
		depth++
		if depth >= 5 {
			break
		}
	}
	fmt.Println("Tidak ditemukan")
	duration := time.Since(startTime)
	durationMS := duration.Milliseconds()
	return nil, checked, 0, strconv.Itoa(int(durationMS)) + "ms"

}


