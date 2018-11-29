package main

import (
	"fmt"
	"sync"
	"time"
)

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok == true {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %v", url)
}

func Crawl(url string, depth int, f fakeFetcher) {
	if depth <= 0 {
		return
	}
	myFoundUrls.mut.Lock()
	if myFoundUrls.v[url] {
		myFoundUrls.mut.Unlock()
		return
	}
	myFoundUrls.v[url] = true
	myFoundUrls.mut.Unlock()
	body, urls, err := f.Fetch(url)
	if err != nil {
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, url := range urls {
		go Crawl(url, depth-1, f)
	}
}

type foundUrls struct {
	v   map[string]bool
	mut sync.Mutex
}

var myFoundUrls = foundUrls{v: make(map[string]bool)}

func main() {
	go Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(3 * time.Second)
}
