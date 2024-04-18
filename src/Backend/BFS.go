package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func BFS(startURL, targetURL string) {
	startTime := time.Now()

	visited := make(map[string]bool)
	queue := []string{startURL}
	paths := make(map[string]string)
	articlesChecked := 0
	articlesInSolution := 0

	for len(queue) > 0 {
		currentURL := queue[0]
		queue = queue[1:]

		if visited[currentURL] {
			continue
		}

		visited[currentURL] = true
		articlesChecked++

		doc, err := goquery.NewDocument(currentURL)
		if err != nil {
			log.Fatal(err)
		}

		foundTarget := false
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if strings.HasPrefix(link, "/wiki/") {
				fullURL := "https://en.wikipedia.org" + link
				if !visited[fullURL] && !foundTarget {
					queue = append(queue, fullURL)
					paths[fullURL] = currentURL
				}
				if fullURL == targetURL {
					foundTarget = true
				}
			}
		})

		if foundTarget {
			// Path found
			path := []string{targetURL}
			currentURL = targetURL
			for currentURL != startURL {
				currentURL = paths[currentURL]
				path = append([]string{currentURL}, path...)
				articlesInSolution++
			}
			fmt.Println("Shortest path:")
			for _, p := range path {
				fmt.Println(p)
			}
			fmt.Println("Articles checked:", articlesChecked)
			fmt.Println("Articles in solution:", articlesInSolution)
			fmt.Println("Time taken:", time.Since(startTime))
			return
		}
	}

	fmt.Println("Path not found")
	fmt.Println("Time taken:", time.Since(startTime))
}

func main() {
	startURL := "https://en.wikipedia.org/wiki/Mathematics"
	targetURL := "https://en.wikipedia.org/wiki/Knowledge"

	BFS(startURL, targetURL)
}
