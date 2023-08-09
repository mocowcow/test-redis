package main

import (
	"bufio"
	"fmt"
	"net/http"
)

const LOOP = 1

func main() {
	for i := 0; i < LOOP; i++ {
		go buy()
	}
}

func buy() {
	resp, err := http.Get("http://localhost:19810/buy")
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
