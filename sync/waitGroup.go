package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	/*
		waitGroup是结构体类型，拷贝类型为深拷贝
		在goroutine中使用wg.add可能直接导致提前wg.wait()
		wg.Add(100) wg.Add(-100) 效率低 循环100遍
	*/
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			//resp, err := http.Get(url)
			_, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			//resp.Write(os.Stdout)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
