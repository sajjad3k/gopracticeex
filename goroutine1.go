package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var g sync.WaitGroup

func Readresponse(url string, num chan int) {

	defer g.Done()
	//var n int = 0
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(len(body))

	num <- len(body)
	//return n
}

func main() {

	num := make(chan int)
	g.Add(1)
	go Readresponse("https://www.golangprograms.com", num)
	fmt.Println(<-num)
	//g.Add(2)
	//	go Readresponse("https://coderwall.com", num)
	//g.Add(3)
	//	go Readresponse("https://stackoverflow.com", num)
	g.Wait()
	//Readresponse()

	fmt.Println("completed")

	//time.Sleep(15 * time.Second)

}
