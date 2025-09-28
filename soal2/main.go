package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Url struct {
	nameURL string
}

// Define function untuk mengecek Url ada atau tidak dengan http.Client
func checkUrl (url string, wg *sync.WaitGroup) {
	defer wg.Done()

	client:= http.Client{
		Timeout: 500 * time.Millisecond,
	}

	ok, err := client.Get(url)
	if err != nil {
		fmt.Printf("URL: %s is not reachable (error).\n", url)
		return
	}
	defer ok.Body.Close()

	// Periksa kode status (kode 200-299)
	if ok.StatusCode >= 200 && ok.StatusCode < 300 {
		fmt.Printf("URL: %s is reachable.\n", url)
	} else {
		fmt.Printf("URL: %s is not reachable.\n", url)
	}
}

func main() {
	urls := []Url{
		{"https://h4cktiv8.com"},
		{"https://google.com"},
		{"https://nonexistenturl.com"},
		{"https://hacktiv8.com"},
		{"https://github.com"},
	}

	var wg sync.WaitGroup

	//Panggil satu persatu atau looping url yang ada di List dengan goroutine
	for _, i := range urls {
		wg.Add(1)
		go checkUrl(i.nameURL, &wg)
	}

	wg.Wait()
}