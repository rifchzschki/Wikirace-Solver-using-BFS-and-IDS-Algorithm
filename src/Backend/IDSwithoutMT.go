package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func DFSConcurrent(currentURL, targetURL string, depth int, visited map[string]bool, paths map[string]string) bool {
    if currentURL == targetURL {
        fmt.Println()
        fmt.Println()
        fmt.Println(currentURL)
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

	found := false
	doc.Find("div#bodyContent").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if link != "" && link[0] == '/' && len(link) > 1 && link[1] != '#' {
				fullURL := "https://en.wikipedia.org" + link
				if !visited[fullURL] && !found {
					paths[fullURL] = currentURL
					found = DFSConcurrent(fullURL, targetURL, depth-1, visited, paths)
				}
			}
		})
	})
	return found
}

func IDSConcurrent(startURL, targetURL string) {
	depth := 1
	visited := make(map[string]bool)
	paths := make(map[string]string)
	found := false

	for !found {
        fmt.Println(depth)
		found = DFSConcurrent(startURL, targetURL, depth, visited, paths)
        fmt.Println(found)
		if found {
            // Path found
			path := []string{targetURL}
			for targetURL != startURL {
                targetURL = paths[targetURL]
				path = append([]string{targetURL}, path...)
			}
			fmt.Println("Shortest path:")
			for _, p := range path {
				fmt.Println(p)
			}
			break
		}
		depth++
		if depth >= 50 {
			fmt.Println("Tidak ditemukan")
			break
		}
	}

}

func main() {
	startURL := "https://en.wikipedia.org/wiki/Jokowi"
	targetURL := "https://en.wikipedia.org/wiki/MNC_Asia_Holding"

	IDSConcurrent(startURL, targetURL)
}
