package youdao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"simple-translate/utils"
	"strings"
	"time"
)

type ResponseBody struct {
	TranslateResult [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	ErrorCode   int    `json:"errorCode"`
	Type        string `json:"type"`
	SmartResult struct {
		Entries []string `json:"entries"`
		Type    int      `json:"type"`
	} `json:"smartResult"`
}

func Translate(word string) {
	salt := time.Now().UnixMilli()
	lts := salt / 10
	bv := utils.Md5encrypt("5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36")
	sign := utils.Md5encrypt("fanyideskweb", word, fmt.Sprint(salt), "Ygy_4c=r#e#4EX^NUGUc5")
	dataStr := fmt.Sprintf(
		"i=%v&from=en&to=zh-CHS&smartresult=dict&client=fanyideskweb&salt=%v&sign=%v&lts=%v&bv=%v&doctype=json&version=2.1&keyfrom=fanyi.web&action=FY_BY_CLICKBUTTION",
		word, salt, sign, lts, bv,
	)

	client := &http.Client{}
	var data = strings.NewReader(dataStr)
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
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode: ", resp.StatusCode, "body", string(bodyText))
	}

	respBody := ResponseBody{}

	err = json.Unmarshal(bodyText, &respBody)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========youdao=========")
	fmt.Println(word)
	entires := respBody.SmartResult.Entries
	for _, v := range entires {
		fmt.Print(v)
	}
	fmt.Println("=======================")
}
