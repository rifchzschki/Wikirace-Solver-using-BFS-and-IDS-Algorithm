package algorithm

import (
	"log"
	"strconv"
	"strings"
	"time"

	"sync"

	"github.com/PuerkitoBio/goquery"
)

// function for check if the page redirected
func isRedirect(doc *goquery.Document) bool {
	return doc.Find("title").Text() == "Redirect" 
}

// function for scrap information from wiki
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
		if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link,":"){
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

// function for run BFS algorithm
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
		if len(queue)>=100000000{
			duration := time.Since(startTime)
			durationMS := duration.Milliseconds()
			return nil,articleChecked,0,strconv.Itoa(int(durationMS)) + "ms"
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
			duration := time.Since(startTime)
			durationMS := duration.Milliseconds()
			return path,articleChecked, len(path),strconv.Itoa(int(durationMS)) + "ms"
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
	duration := time.Since(startTime)
	durationMS := duration.Milliseconds()
	return nil, 0,0,strconv.Itoa(int(durationMS)) + "ms"
}


