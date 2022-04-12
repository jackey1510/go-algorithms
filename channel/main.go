package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)

	}

	for l := range ch {
		go func(s string) {
			time.Sleep(time.Second * 2)
			checkLink(s, ch)
		}(l)
		// go checkLink(l, ch)
	}

}

func checkLink(link string, ch chan string) {
	time.Sleep(time.Second * 10)
	_, err := http.Get(link)
	if err == nil {
		fmt.Println(link, "is up!")
		ch <- link
		return
	}
	fmt.Println(link, "might be down.", err.Error())
	ch <- link
}
