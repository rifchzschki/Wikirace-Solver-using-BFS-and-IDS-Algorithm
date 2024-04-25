package algorithm

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)



func DFSConcurrent(currentURL, targetURL string, depth int, visited map[string]bool, paths map[string]string, articlesChecked *int) bool {
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
	*articlesChecked ++

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

	found := false
	doc.Find("div#bodyContent").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if strings.HasPrefix(link, "/wiki/") && !hasPrefix(unwantedWikiPrefixes[:],link) && !strings.Contains(link,":") {
				fullURL := "https://en.wikipedia.org" + link
				if !visited[fullURL] && !found {
					paths[fullURL] = currentURL
					found = DFSConcurrent(fullURL, targetURL, depth-1, visited, paths, articlesChecked)
				}
			}
		})
	})
	return found
}

func IDSConcurrent(startURL, targetURL string) ([] string, int, int, string) {
	startTime := time.Now()
	depth := 1
	visited := make(map[string]bool)
	paths := make(map[string]string)
	found := false
	checked :=0

	for !found {
        // fmt.Println(depth)
		found = DFSConcurrent(startURL, targetURL, depth, visited, paths, &checked)
        // fmt.Println(found)
		if found {
            // Path found
			path := []string{targetURL}
			for targetURL != startURL {
                targetURL = paths[targetURL]
				path = append([]string{targetURL}, path...)
			}
			return path, checked, len(path), time.Since(startTime).String()
		}
		depth++
		if depth >= 50 {
			break
		}
	}
	fmt.Println("Tidak ditemukan")
	return nil, checked, 0, time.Since(startTime).String()

}

