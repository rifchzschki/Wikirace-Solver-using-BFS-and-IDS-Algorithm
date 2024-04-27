package algorithm

import (
	"fmt"
	"strings"
)

func getTitle(path []string) ([] string) {
	var titles []string
    for _, url := range path { 
		title := strings.TrimPrefix(url, "https://en.wikipedia.org/wiki/")
		title =  strings.ReplaceAll(title, "_", " ")
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
	fmt.Println(startURL, targetURL)
	if(strings.EqualFold(algorithm, "ids")){
		fmt.Println("Pencarian menggunakan algoritma IDS")
		rute,banyakArtic, panjangRute, durasi = IDS(startURL, targetURL)
	}else{
		fmt.Println("Pencarian menggunakan algoritma BFS")
		rute,banyakArtic, panjangRute, durasi = BFS(startURL, targetURL)
	}
	titles = getTitle(rute)
	fmt.Println("ini rute: ", rute)
	return titles,banyakArtic, panjangRute, durasi
}