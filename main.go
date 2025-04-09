package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func main() {
	//  Configuration Variables
	const (
		targetUsername = "any_username"                 // This is me // Target NGL username
		message        = "We have read your blog Ayush" // Message to send
		deviceID       = "e092577e-f5f596a7b202"        //any uuid
		// Number of messages to send
		requestCount = 10 //change it to 1000 if you are evil
	)

	var wg sync.WaitGroup
	endpoint := "https://ngl.link/api/submit"
	method := "POST"

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			data := url.Values{}
			data.Set("username", targetUsername)
			data.Set("question", message)
			data.Set("deviceId", deviceID)
			data.Set("gameSlug", "")
			data.Set("referrer", "")

			payload := strings.NewReader(data.Encode())

			client := &http.Client{}
			req, err := http.NewRequest(method, endpoint, payload)
			if err != nil {
				fmt.Println("Request creation error:", err)
				return
			}

			// Headers (copied from real browser request)
			// you may not be able to use these headers as they may get expired by the time you are reading this blog
			// replace your own headers from postman,
			// these headers will be visible on postman golang code too

			req.Header.Add("accept", "*/*")
			req.Header.Add("accept-language", "en-US,en;q=0.9")
			req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
			req.Header.Add("origin", "https://ngl.link")
			req.Header.Add("priority", "u=1, i")
			req.Header.Add("referer", "https://ngl.link/"+targetUsername)
			req.Header.Add("sec-ch-ua", `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`)
			req.Header.Add("sec-ch-ua-mobile", "?1")
			req.Header.Add("sec-ch-ua-platform", `"Android"`)
			req.Header.Add("sec-fetch-dest", "empty")
			req.Header.Add("sec-fetch-mode", "cors")
			req.Header.Add("sec-fetch-site", "same-origin")
			req.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Mobile Safari/537.36")
			req.Header.Add("x-requested-with", "XMLHttpRequest")
			req.Header.Add("Authorization", "Basic TUtGV0ZZTjJRMtlhGWUdORlFVVzpFbWNJc4dFJVS2ZuZyQWa2RR")

			/////////////////////////////////////////

			res, err := client.Do(req)
			if err != nil {
				fmt.Println("HTTP request error:", err)
				return
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Println("Response reading error:", err)
				return
			}

			fmt.Println("Response:", string(body))
		}()
	}

	wg.Wait() // Waits for all goroutines to finish
}
