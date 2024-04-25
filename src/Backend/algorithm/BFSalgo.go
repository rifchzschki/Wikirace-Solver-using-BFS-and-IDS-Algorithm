package algorithm

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var unwantedWikiPrefixes = [...]string{"Wikipedia:", "Category:", "File:",
	"Portal:", "Template:", "CS1_maint:", "Special:", "Template_talk:", "Help:", "Talk:"}

func hasPrefix(names []string, str string) bool {
	for _, name := range names {
		if strings.HasPrefix(str, name) {
			return true
		}
	}
	return false
}
func isRedirect(doc *goquery.Document) bool {
	return doc.Find("title").Text() == "Redirected" 
}

func scrapping(currentURL string)([]string){
	queue := []string{}
	visited := make(map[string]bool)

	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", currentURL, err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Unexpected status code %d for URL %s\n", resp.StatusCode, currentURL)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if isRedirect(doc) {
		return queue
	}
	
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if strings.HasPrefix(link, "/wiki/") && !hasPrefix(unwantedWikiPrefixes[:],link) && !strings.Contains(link,":"){
			fullURL := "https://en.wikipedia.org" + link
			
			
			if !visited[fullURL] {
				// fmt.Printf("Parent : %s, Child : %s\n", currentURL, fullURL)
				queue = append(queue, fullURL)
				visited[fullURL] = true 
			}
		}
	})
	return queue
}
func BFS(startURL, targetURL string) ([] string, int, int, string) {
	startTime := time.Now()

	visited := make(map[string]bool)
	queue := []string{startURL}
	paths := make(map[string]string)
	articlesChecked := 0

	for len(queue) > 0 {
		currentURL := queue[0]
		queue = queue[1:]
		found:

		if currentURL == targetURL{
			path := []string{targetURL}
			currentURL = targetURL
			for currentURL != startURL {
				currentURL = paths[currentURL]
				path = append([]string{currentURL}, path...)
			}
			

			return path, articlesChecked, len(path), time.Since(startTime).String()
		}
		
		if visited[currentURL] {
			continue
		}

		visited[currentURL] = true
		articlesChecked++

		neighbors := scrapping(currentURL)
		for _, neighbor := range neighbors {
			if neighbor == targetURL{
				paths[neighbor] = currentURL
				currentURL = neighbor
				goto found
			}
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				
				if _, ok := paths[neighbor]; !ok {
					paths[neighbor] = currentURL
				}
				
			}
		}
	}
	

	return nil , articlesChecked, 0, time.Since(startTime).String()
}

