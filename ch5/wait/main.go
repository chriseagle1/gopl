package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	WaitForServer("https://github.com")
}

func WaitForServer(url string) error  {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}

		fmt.Printf("Server not responding (%s);retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
}

