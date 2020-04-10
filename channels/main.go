package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amazon.com",
		"https://www.foundragon.com",
		"https://www.geekudos.com",
		"http://www.wongshengchuang.me",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }
	// equipvalent to

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
