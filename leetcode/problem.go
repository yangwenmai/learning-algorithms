package leetcode

import (
	"fmt"
	"strings"
)

type Problem struct {
	ID                                 int
	Title                              string
	TitleSlug                          string
	PassRate                           string
	Difficulty                         string
	IsAccepted, IsPaid, IsFavor, IsNew bool
	HasNoGoOption                      bool // 不能够使用 Go 语言解答
}

func newProblem(ps problemStatus) Problem {
	level := []string{"", "Easy", "Medium", "Hard"}

	p := Problem{
		ID:        ps.State.ID,
		Title:     ps.State.Title,
		TitleSlug: ps.State.TitleSlug,
		// p.Submitted + 1 是因为刚刚添加的新题的 submitted 为 0
		PassRate:   fmt.Sprintf("%d%%", ps.ACs*100/(ps.Submitted+1)),
		Difficulty: level[ps.Difficulty.Level],
		IsAccepted: ps.Status == "ac",
		IsPaid:     ps.IsPaid,
		IsFavor:    ps.IsFavor,
		IsNew:      ps.State.IsNew,
	}

	return p
}

func (p Problem) IsAvailable() bool {
	if p.ID == 0 || p.IsPaid || p.HasNoGoOption {
		return false
	}
	return true
}

func (p Problem) Dir() string {
	path := "practive/leetcode"
	return fmt.Sprintf("./%s/%04d.%s", path, p.ID, p.TitleSlug)
}

func (p Problem) Link() string {
	return fmt.Sprintf("https://leetcode.com/problems/%s/", p.TitleSlug)
}

func (p Problem) tableLine() string {
	// 题号
	res := fmt.Sprintf("|[%04d](%s)|", p.ID, p.Link())

	// 标题
	t := ""
	if p.IsAccepted {
		t = fmt.Sprintf(`✅[%s](%s)`, strings.TrimSpace(p.Title), p.Dir())
	} else {
		t = fmt.Sprintf(` * %s`, p.Title)
	}
	if p.IsNew {
		t += " :new: "
	}
	res += t + "|"

	// 通过率
	res += fmt.Sprintf("%s|", p.PassRate)

	// 难度
	res += fmt.Sprintf("%s|", p.Difficulty)

	// 收藏
	f := ""
	if p.IsFavor {
		f = "[❤](https://leetcode.com/list/517142r)"
	}
	res += fmt.Sprintf("%s|\n", f)

	return res
}

func (p Problem) listLine() string {
	return fmt.Sprintf("- [%d.%s](%s)\n", p.ID, p.Title, p.Link())
}

func (p Problem) DidaTask(prefix string) string {
	return fmt.Sprintf("#%s - %04d - #%s - %s - %s - %s", prefix, p.ID, p.Difficulty, p.PassRate, p.Title, p.Link())
}

func (p Problem) PackageName() string {
	return fmt.Sprintf("problem%04d", p.ID)
}
