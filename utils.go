package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func md5encrypt(s ...string) string {
	h := md5.New()
	for _, v := range s {
		io.WriteString(h, v)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
