package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

type URLConfig struct {
	URLs []string `json:"urls"`
}

func main() {
	ticker := time.NewTicker(time.Minute*5)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := os.ReadFile("urls.json")
			if err != nil {
				fmt.Println("Error reading JSON file:", err)
				return
			}

			var config URLConfig
			if err := json.Unmarshal(data, &config); err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}

			var wg sync.WaitGroup

			for _, url := range config.URLs {
				wg.Add(1)
				go makeHTTPRequest(url, &wg)
			}

			wg.Wait()

			fmt.Println("All HTTP requests completed.")
		}
	}
}

func makeHTTPRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request to %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Printf("GET request successful for %s\n", url)
	} else {
		fmt.Printf("Failed to make GET request to %s. Status code: %d\n", url, response.StatusCode)
	}
}
