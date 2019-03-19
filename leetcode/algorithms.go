package leetcode

import (
	"encoding/json"
	"log"

	"github.com/aQuaYi/GoKit"
	"github.com/yangwenmai/learning-algorithms/util"
)

const (
	unavailableFile = "unavailable.json"
)

// algorithms 保存API信息
type algorithms struct {
	Name     string          `json:"category_slug"`
	User     string          `json:"user_name"`
	ACEasy   int             `json:"ac_easy"`
	ACMedium int             `json:"ac_medium"`
	ACHard   int             `json:"ac_hard"`
	AC       int             `json:"num_solved"`
	Problems []problemStatus `json:"stat_status_pairs"`
}

type problemStatus struct {
	Status     string `json:"status"`
	State      `json:"stat"`
	IsFavor    bool `json:"is_favor"`
	IsPaid     bool `json:"paid_only"`
	Difficulty `json:"difficulty"`
}

// State 保存单个问题的解答状态
type State struct {
	ACs       int    `json:"total_acs"`
	Title     string `json:"question__title"`
	IsNew     bool   `json:"is_new_question"`
	Submitted int    `json:"total_submitted"`
	ID        int    `json:"frontend_question_id"`
	TitleSlug string `json:"question__title_slug"`
}

// Difficulty 问题的难度
type Difficulty struct {
	Level int `json:"level"`
}

func fetchAlgorithms(username string) *algorithms {
	URL := "https://leetcode.com/api/problems/Algorithms/"

	raw := fetchRaw(URL)
	res := new(algorithms)
	if err := json.Unmarshal(raw, res); err != nil {
		log.Panicf("无法把json转换成Category: %s\n", err.Error())
	}

	// 如果没有登录的话，也能获取数据，但是用户名，就不是本人
	if res.User != username {
		log.Fatal("没有获取到本人的数据")
	}

	return res
}

func parseAlgorithms(alg *algorithms) (*problems, Record) {
	hasNoGoOption := readUnavailable()
	problems := &problems{}
	r := Record{}

	for _, ps := range alg.Problems {
		p := newProblem(ps)
		if hasNoGoOption[p.ID] {
			p.HasNoGoOption = true
		}
		problems.add(p)
		r.Update(p)
	}

	return problems, r
}

func readUnavailable() map[int]bool {
	type unavailable struct {
		List []int
	}

	if !GoKit.Exist(unavailableFile) {
		log.Panicf("%s 不存在，没有不能解答的题目", unavailableFile)
	}

	raw := util.ReadFile(unavailableFile)
	u := unavailable{}
	if err := json.Unmarshal(raw, &u); err != nil {
		log.Panicf("获取 %s 失败：%s", unavailableFile, err)
	}

	res := make(map[int]bool, len(u.List))
	for i := range u.List {
		res[u.List[i]] = true
	}

	return res
}
