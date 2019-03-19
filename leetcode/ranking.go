package leetcode

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// fetchUserRanking fetch user's ranking
func fetchUserRanking(username string) int {
	// 获取网页数据
	URL := fmt.Sprintf("https://leetcode.com/%s/", username)
	data := fetchRaw(URL)
	str := string(data)

	// 通过不断裁剪 str 获取排名信息
	i := strings.Index(str, "ng-init")
	j := i + strings.Index(str[i:], "ng-cloak")
	str = str[i:j]

	i = strings.Index(str, "(")
	j = strings.Index(str, ")")
	str = str[i:j]
	strs := strings.Split(str, ",")
	str = strs[6]
	i = strings.Index(str, "'")
	j = 2 + strings.Index(str[2:], "'")
	str = str[i+1 : j]
	r, err := strconv.Atoi(str)
	if err != nil {
		log.Panicf("无法把 %s 转换成数字Ranking", str)
	}

	return r
}
