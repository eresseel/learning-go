package main

import (
	"fmt"
)

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// type result struct {
// 	url    string
// 	status bool
// }

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
// 	results := make(map[string]bool)
// 	resultChannel := make(chan result)

// 	for _, url := range urls {
// 		// fontosság: átadjuk paraméterként az url-t, hogy ne legyen loop-capture bug
// 		go func(u string) {
// 			resultChannel <- result{url: u, status: wc(u)}
// 		}(url)
// 	}

// 	for i := 0; i < len(urls); i++ {
// 		r := <-resultChannel
// 		results[r.url] = r.status
// 	}

// 	return results
// }

func main() {
	urls := []string{
		"http://google.com",
		"http://badurl.com",
		"http://github.com",
	}

	checker := func(url string) bool {
		return url != "http://badurl.com"
	}

	results := CheckWebsites(checker, urls)

	fmt.Println(results)
}
