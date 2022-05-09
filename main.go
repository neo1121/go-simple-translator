package main

import (
	"fmt"
	"os"
	"simple-translate/caiyun"
	"simple-translate/youdao"
	"sync"
)

// TODO:
// [x] 1. send a HTTP request
// [x] 2. resolve the request
// [x] 3. resolve the response
// [x] 4. input word and format output
// [x] 5. review code

func Translate(word string, translateFunc func(string), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		// timer := utils.Timer{}
		// defer timer.Close()
		// timer.Start()
		translateFunc(word)
		// fmt.Println("耗时:", timer.End(), "ms")
	}()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Usage: translate <word>")
		fmt.Println("Example: translate hello")
		return
	}
	word := os.Args[1]
	wg := &sync.WaitGroup{}
	// timer := utils.Timer{}
	// defer timer.Close()
	// timer.Start()
	Translate(word, youdao.Translate, wg)
	Translate(word, caiyun.Translate, wg)
	wg.Wait()
	// fmt.Println("=======================")
	// fmt.Println("总耗时:", timer.End(), "ms")
	// fmt.Println("=======================")
}
