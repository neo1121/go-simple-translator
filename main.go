package main

import (
	"fmt"
	"os"
	"simple-translate/caiyun"
	"simple-translate/youdao"
	"sync"
	"time"
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
		startTime := time.Now().UnixMilli()
		translateFunc(word)
		endTime := time.Now().UnixMilli()
		fmt.Println("耗时:", endTime-startTime, "ms")
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
	Translate(word, youdao.Translate, wg)
	Translate(word, caiyun.Translate, wg)
	wg.Wait()
}
