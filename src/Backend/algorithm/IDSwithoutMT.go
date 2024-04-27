package algorithm

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)



func DFSConcurrent(currentURL, targetURL string, depth int, visited map[string]bool, paths map[string]string, articlesChecked *int, file *os.File) bool {
	
	output := fmt.Sprintln(currentURL)
	file.WriteString(output)
	
	
	
	fmt.Println(depth)
	file.WriteString(strconv.Itoa(depth))
	fmt.Println(currentURL)
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
	
	if (isRedirect(doc)){
		return false
	}
	
	*articlesChecked ++
	// file.WriteString(strconv.Itoa(*articlesChecked))
	found := false
	doc.Find("div#bodyContent").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if strings.HasPrefix(link, "/wiki/") && !hasPrefix(unwantedWikiPrefixes[:],link) && !strings.Contains(link,":") {
				fullURL := "https://en.wikipedia.org" + link
				if !visited[fullURL] && !found {
					*articlesChecked ++
					// file.WriteString(strconv.Itoa(*articlesChecked))
					paths[fullURL] = currentURL
					found = DFSConcurrent(fullURL, targetURL, depth-1, visited, paths, articlesChecked, file)
				}
			}
		})
	})
	return found
}

func IDSConcurrent(startURL, targetURL string) ([] string, int, int, string) {
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		// return false
	}
	defer file.Close()
	startTime := time.Now()
	depth := 1
	paths := make(map[string]string)
	found := false
	checked :=0
	
	for !found {
		visited := make(map[string]bool)
        // fmt.Println(depth)
		found = DFSConcurrent(startURL, targetURL, depth, visited, paths, &checked, file)
        // fmt.Println(found)
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
		file.WriteString("ahahahahahahahahahah")
	}
	fmt.Println("Tidak ditemukan")
	duration := time.Since(startTime)
	durationMS := duration.Milliseconds()
	return nil, checked, 0, strconv.Itoa(int(durationMS)) + "ms"

}

