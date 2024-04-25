package algorithm

import (
	
	"log"
	"strings"
	"time"
	
	"sync"
	// "github.com/gin-gonic/gin"
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
	exist := make(map[string]bool) 

	doc, err := goquery.NewDocument(currentURL)
	if err != nil {
		log.Fatal(err)
	}
	if isRedirect(doc) {
		return queue
	}
	doc.Find("div#bodyContent").Each(func(i int, s *goquery.Selection) {
		s.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if strings.HasPrefix(link, "/wiki/") && !hasPrefix(unwantedWikiPrefixes[:],link) && !strings.Contains(link,":"){
			fullURL := "https://en.wikipedia.org" + link
			
			
			if !exist[fullURL] {
				queue = append(queue, fullURL)
				exist[fullURL] = true 
			}
		}
	})})
	return queue
}


const max = 50 
// jumlah go routine
func BFS(startURL, targetURL string) ([]string, int,int,string) {
	startTime := time.Now()
	visited := make(map[string]bool)
	queue := make(chan string,100000000)  
	paths := make(map[string]string)
	articleChecked :=0
	var mu sync.RWMutex
	var wg sync.WaitGroup
	ch := make(chan bool, max)
	go func() {
		queue <- startURL  
	}()
	
	for currentURL := range queue {
		articleChecked++
		if time.Since(startTime) > (5 * time.Minute) {

			return nil,articleChecked,0,time.Since(startTime).String()
		}
		here:
		mu.Lock()
		if currentURL == targetURL {
			
			path := []string{targetURL}
			currentURL = targetURL
			for currentURL != startURL {
				currentURL = paths[currentURL]
				path = append([]string{currentURL}, path...)
			}
			
			return path,articleChecked, len(path),time.Since(startTime).String()
		}
		mu.Unlock()

		if visited[currentURL] {
			continue
		}

		mu.Lock()
		visited[currentURL] = true
		mu.Unlock()
		ch <- true
		wg.Add(1)
		
		copyURL := currentURL 

		go func(url string, targetURL string, paths *map[string]string) {
			defer wg.Done()
			defer func() { <-ch }()
			neighbors := scrapping(url)
			for _, neighbor := range neighbors {
				mu.Lock()
				if neighbor == targetURL {
					if _, exist := (*paths)[neighbor]; !exist {
						(*paths)[neighbor] = url
					}
					queue <- neighbor
					copyURL = neighbor
					mu.Unlock()
					return
				}
				mu.Unlock()

				if !visited[neighbor] {
					mu.Lock()
					if _, exist := (*paths)[neighbor]; !exist {
						(*paths)[neighbor] = url
					}
					mu.Unlock()
					queue <- neighbor
				}
			}
		}(copyURL, targetURL, &paths)
		wg.Wait()
		mu.Lock()
		if copyURL == targetURL{
			currentURL = copyURL
			mu.Unlock()
			goto here
		}
		mu.Unlock()
	}

	wg.Wait()
	close(queue)  
	return nil, 0,0,time.Since(startTime).String()
}


