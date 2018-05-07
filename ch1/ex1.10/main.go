package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type result struct {
	url  string
	body string
}

func main() {

	args := os.Args[1:]
	ch := make(chan *result)

	start := time.Now()

	// fetch
	for _, url := range args {
		go fetch(url, ch) // ゴルーチンを開始
	}

	// save fetch results to a file.
	for range args {
		save(<-ch) // chチャネルから受信
	}

	fmt.Printf("% 2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- *result) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	ch <- &result{
		url,
		string(bodyBytes),
	}
}

func save(r *result) {
	name := fmt.Sprintf("%s@%v", strings.Replace(r.url, "/", "_", -1), time.Now())
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = file.WriteString(r.body)
	if err != nil {
		log.Fatalln(err)
	}

}
