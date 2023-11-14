package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchWiki(city string, wg *sync.WaitGroup) {
	defer wg.Done()
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", city)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching page for city %s: %s\n", city, err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	start := time.Now()

	cities := []string{"Detroit", "Seoul", "Paris", "Manila", "Denver"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWiki(city, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("Time: ", time.Since(start))
}
