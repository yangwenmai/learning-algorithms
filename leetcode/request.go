package leetcode

import (
	"io/ioutil"
	"log"
)

func fetchRaw(URL string) []byte {
	log.Printf("开始下载 %s 的数据", URL)

	req := newLeetCodeReq()
	resp, err := req.Get(URL)
	if err != nil {
		log.Fatal("getRaw: Get Error: " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("getRaw: Read Error: " + err.Error())
	}
	return body
}
