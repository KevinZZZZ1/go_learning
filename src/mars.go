package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}

	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}

		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func filterGopher2(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}

	close(downstream)
}

func printGopher(upstream chan string) {
	for {
		item, ok := <-upstream

		if !ok {
			return
		}
		fmt.Println(item)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
