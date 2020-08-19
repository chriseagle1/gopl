package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	f, err := os.Create("text" + strconv.Itoa(rand.Intn(10)))

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer f.Close()

	outPut := bufio.NewWriter(f)

	nbytes, err := io.Copy(outPut, resp.Body)

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f    %7d    %s", secs, nbytes, url)
}
