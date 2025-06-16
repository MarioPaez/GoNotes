package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url   string
	valid bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		go func() {
			resultsChannel <- result{url, wc(url)}
		}()
	}

	for range len(urls) {
		r := <-resultsChannel
		results[r.url] = r.valid
	}

	return results
}
