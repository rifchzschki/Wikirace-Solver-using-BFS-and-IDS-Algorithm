package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/PuerkitoBio/goquery"
)

func BFS(startURL, targetURL string) (string, string, string, time.Duration) {
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
			duration := time.Since(startTime)
			return "Shortest path: " + strings.Join(path, " -> "), fmt.Sprintf("Articles checked: %d", articlesChecked), fmt.Sprintf("Articles in solution: %d", articlesInSolution), duration
			// fmt.Println("Shortest path:")
			// for _, p := range path {
			// 	fmt.Println(p)
			}
			// fmt.Println("Articles checked:", articlesChecked)
			// fmt.Println("Articles in solution:", articlesInSolution)
			// fmt.Println("Time taken:", time.Since(startTime))
			// return
		}
	

	return "Path not found", "", "", 0
}

func main() {
	r := gin.Default()

	r.POST("/findpath", func(c *gin.Context) {
		var input struct {
			StartURL  string `json:"startURL"`
			TargetURL string `json:"targetURL"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		path, checked, inSolution, duration := BFS(input.StartURL, input.TargetURL)
		c.JSON(200, gin.H{
			"path":           path,
			"checked":        checked,
			"inSolution":     inSolution,
			"timeTaken":      duration.String(),
		})
	})

	r.Run(":8080")
}
