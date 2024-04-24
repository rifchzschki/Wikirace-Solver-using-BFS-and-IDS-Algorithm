package algorithm

import (
	"fmt"
	"strings"
)

func getTitle(path []string) ([] string) {
	var titles []string
    for _, url := range path { // Loop melalui setiap elemen di slice path
		title := strings.TrimPrefix(url, "https://en.wikipedia.org/wiki/")
		titles = append(titles, title)
	}
	return titles
}

func Test(algorithm, startURL, targetURL string) ([] string, int, int, string) {
	var(
		rute []string
		titles []string
		banyakArtic int
		panjangRute int
		durasi string
	)
	if(strings.EqualFold(algorithm, "ids")){
		fmt.Println("Pencarian menggunakan algoritma IDS")
		rute,banyakArtic, panjangRute, durasi = IDSConcurrent(startURL, targetURL)
	}else{
		fmt.Println("Pencarian menggunakan algoritma BFS")
		rute,banyakArtic, panjangRute, durasi = BFS(startURL, targetURL)
	}
	titles = getTitle(rute)
	return titles,banyakArtic, panjangRute, durasi
}