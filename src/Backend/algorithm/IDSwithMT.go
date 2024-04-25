package algorithm

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var mu sync.Mutex

func DFSConcurrentMT(currentURL, targetURL string, depth int, visited map[string]bool, paths map[string]string, ch chan bool, wg *sync.WaitGroup) {
    defer wg.Done()
    if currentURL == targetURL {
        mu.Lock()
        fmt.Println("panjang ch", len(ch))
        fmt.Println()
        fmt.Println()
        fmt.Println(currentURL)
        if(len(ch)<=0){
            ch <- true
        }
        mu.Unlock()
        return
    }
    if depth == 0 {
        return
    }


	mu.Lock()
	visited[currentURL] = true
	mu.Unlock()

	resp, err := http.Get(currentURL)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %v\n", currentURL, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Unexpected status code %d for URL %s\n", resp.StatusCode, currentURL)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing document from URL %s: %v\n", currentURL, err)
		return
	}

	doc.Find("div#bodyContent").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if link != "" && link[0] == '/' && len(link) > 1 && link[1] != '#' {
				fullURL := "https://en.wikipedia.org" + link
				mu.Lock()
				if !visited[fullURL] {
					paths[fullURL] = currentURL
                    wg.Add(1)
					go DFSConcurrentMT(fullURL, targetURL, depth-1, visited, paths, ch, wg)
				}
				mu.Unlock()
			}
		})
	})
}

func IDSConcurrentMT(startURL, targetURL string) {
	depth := 1
	visited := make(map[string]bool)
	paths := make(map[string]string)
	found := false
    var wg sync.WaitGroup

	for !found {
        fmt.Println(depth)
		ch := make(chan bool,10000)
        wg.Add(1)
		go DFSConcurrentMT(startURL, targetURL, depth, visited, paths, ch, &wg)
        if(len(ch)>0){
            found = <-ch
            fmt.Println("woy anjing")
        }
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
        wg.Wait()
        fmt.Println("woy anjing")
		depth++
		if depth >= 100 {
			fmt.Println("Tidak ditemukan")
			break
		}
	}

}
