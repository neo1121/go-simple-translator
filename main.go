package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// TODO:
// [x] 1. send a HTTP request

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`i=hello&from=en&to=zh-CHS&smartresult=dict&client=fanyideskweb&salt=16520888719980&sign=e0982c1b72361100301ddc7c504f0d07&lts=1652088871998&bv=a6a7eab4afbf9b019ca15a461e45e966&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_CLICKBUTTION`)
	req, err := http.NewRequest("POST", "https://fanyi.youdao.com/translate_o?smartresult=dict&smartresult=rule", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en,zh-CN;q=0.9,zh;q=0.8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", `OUTFOX_SEARCH_USER_ID_NCOO=820764440.9262143; OUTFOX_SEARCH_USER_ID="-1211735701@10.110.96.160"; JSESSIONID=aaaanulxstvTB7Ff1kOcy; fanyi-ad-id=305838; fanyi-ad-closed=1; ___rl__test__cookies=1652088871994`)
	req.Header.Set("Origin", "https://fanyi.youdao.com")
	req.Header.Set("Referer", "https://fanyi.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
