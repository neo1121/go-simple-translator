package main

import (
	"fmt"
	"os"
)

// TODO:
// [x] 1. send a HTTP request
// [x] 2. resolve the request
// [x] 3. resolve the response
// [x] 4. input word and format output
// [x] 5. review code

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Usage: translate <word>")
		fmt.Println("Example: translate hello")
		return
	}
	word := os.Args[1]
	youdaoTranslate(word)
}
