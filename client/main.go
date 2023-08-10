package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sync"
)

const LOOP = 10

func main() {
	var wg sync.WaitGroup

	for i := 0; i < LOOP; i++ {
		wg.Add(1)
		go buy(&wg)
	}

	wg.Wait()
}

func buy(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("http://localhost:19810/buy/3")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
