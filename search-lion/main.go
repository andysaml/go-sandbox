package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var africaString string

//var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func recurciveSearch(s string) {

	if len(s) < 5 {
		return
	}

	half1 := s[0 : len(s)/2]
	half2 := s[len(s)/2 : len(s)]
	fmt.Println(len(half1), len(half2))

	if strings.Contains(half1, "lion") {

		fmt.Println("Lion is found recurcively!")
		return
	} else {
		recurciveSearch(half2)
	}
	return
}

func main() {
	type T string
	var slice []string
	var wg sync.WaitGroup

	queue := make(chan string, 1)

	// Create our data and send it into the queue.
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			// defer wg.Done()  <- will result in the last int to be missed in the receiving channel
			queue <- randSeq(10e2)
		}()
	}

	go func() {
		// defer wg.Done() <- Never gets called since the 100 `Done()` calls are made above, resulting in the `Wait()` to continue on before this is executed
		for t := range queue {
			slice = append(slice, t)
			africaString += t
			wg.Done() // ** move the `Done()` call here
		}
	}()

	wg.Wait()

	// now prints off all 100  values
	//fmt.Println(slice)
	//	fmt.Println(africaString)
	fmt.Println(strings.Contains(africaString, "lion"))
	recurciveSearch(africaString)
}
