package build

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"syscall"

	"github.com/yangwenmai/learning-algorithms/util"

	"github.com/yangwenmai/learning-algorithms/leetcode"

	"github.com/aQuaYi/GoKit"
)

type ProblemBuilder struct {
	*Builder
	ProblemNum int
}

func (p *ProblemBuilder) Build() {
	log.Printf("~~ 开始生成第 %d 题的文件夹 ~~\n", p.ProblemNum)

	prob := p.LC.Problems[p.ProblemNum]
	buildDir(prob)
	createProblemFiles(prob)
	log.Printf("~~ 第 %d 题的文件夹，已经生成 ~~\n", p.ProblemNum)
}

func CheckProblemNumValid(lc *leetcode.Leetcode, problemNum int) {
	if problemNum >= len(lc.Problems) {
		log.Panicf("%d 超出题目范围，请核查题号。", problemNum)
	}
	if lc.Problems[problemNum].ID == 0 {
		log.Panicf("%d 号题不存，请核查题号。", problemNum)
	}
	if lc.Problems[problemNum].IsPaid {
		log.Panicf("%d 号题需要付费。如果已经订阅，请注释掉本代码。", problemNum)
	}
	if lc.Problems[problemNum].HasNoGoOption {
		log.Panicf("%d 号题，没有提供 Go 解答选项。请核查后，修改 unavailable.json 中的记录。", problemNum)
	}
}

func buildDir(p leetcode.Problem) {
	if GoKit.Exist(p.Dir()) {
		log.Panicf("第 %d 题的文件夹已经存在，请 **移除** %s 文件夹后，再尝试。", p.ID, p.Dir())
	}

	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Println(err)
			log.Println("清理不必要的文件")
			os.RemoveAll(p.Dir())
		}
	}()

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	err := os.Mkdir(p.Dir(), 0755)
	if err != nil {
		log.Panicf("无法创建目录，%s ：%s", p.Dir(), err)
	}

	log.Printf("开始创建 %d %s 的文件夹...\n", p.ID, p.Title)
}

func createProblemFiles(p leetcode.Problem) {
	detail := leetcode.FetchQuestion(p)

	fc := getGoFunction(detail)
	fcName, para, ans, fc := parseFunction(fc)

	creatGo(p, detail, fc, ans)
	creatGoTest(p, fcName, para, ans)
	creatREADME(p, detail)

	log.Printf("%d.%s 的文件夹，创建完毕。\n", p.ID, p.Title)
}

var typeMap = map[string]string{
	"int":     "0",
	"float64": "0",
	"string":  "\"\"",
	"bool":    "false",
}

func creatGo(p leetcode.Problem, detail *leetcode.Question, function, ansType string) {

	fileFormat := `package %s

/*
	[%d] %s

	https://leetcode.com/problems/%s/description/

	algorithms %s (%s)
	Total Accepted:    %s(%d)
	Total Submissions: %s(%d)

%s
*/

%s`
	content := fmt.Sprintf(fileFormat,
		p.PackageName(), p.ID, p.Title, p.TitleSlug, p.Difficulty, detail.Stats.AcRate,
		detail.Stats.TotalAccepted, detail.Stats.TotalAcceptedRaw,
		detail.Stats.TotalSubmission, detail.Stats.TotalSubmissionRaw,
		strings.TrimSpace(detail.Content), function)

	returns := "\treturn nil\n}"
	if v, ok := typeMap[ansType]; ok {
		returns = fmt.Sprintf("\treturn %s\n}", v)
	}
	content = strings.Replace(content, "}", returns, -1)
	filename := fmt.Sprintf("%s/%s.go", p.Dir(), p.TitleSlug)
	util.WriteFile(filename, content)
}

func creatGoTest(p leetcode.Problem, fcName, para, ansType string) {
	testCasesFormat := `var tcs = []struct {
	%s
	ans %s
}{
	// 可以有多个 testcase

}`

	para = strings.Replace(para, ",", "\n", -1)

	testCases := fmt.Sprintf(testCasesFormat, para, ansType)

	testFuncFormat := `
func Test_%s(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, %s(%s), "输入:%s", tc)
	}
}`
	tcPara := getTcPara(para)
	testFunc := fmt.Sprintf(testFuncFormat, fcName, fcName, tcPara, `%v`)

	benchFuncFormat := `
func Benchmark_%s(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			%s(%s)
		}
	}
}`
	benchFunc := fmt.Sprintf(benchFuncFormat, fcName, fcName, tcPara)

	fileFormat := `package %s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
%s
%s
%s
`
	content := fmt.Sprintf(fileFormat, p.PackageName(), testCases, testFunc, benchFunc)
	filename := fmt.Sprintf("%s/%s_test.go", p.Dir(), p.TitleSlug)
	util.WriteFile(filename, content)
}

// 把函数的参数变成 tc 的参数
func getTcPara(para string) string {
	// 把 para 按行切分
	paras := strings.Split(para, "\n")

	// 把单个参数按空格，切分成参数名和参数类型
	temp := make([][]string, len(paras))
	for i := range paras {
		temp[i] = strings.Split(strings.TrimSpace(paras[i]), ` `)
	}

	// 在参数名称前添加 "tc." 并组合在一起
	res := ""
	for i := 0; i < len(temp); i++ {
		res += ", tc." + temp[i][0]
	}

	return res[2:]
}

func creatREADME(p leetcode.Problem, detail *leetcode.Question) {
	fileFormat := `# [%d. %s](%s)

%s
`
	questionDescription := strings.TrimSpace(detail.Content)
	content := fmt.Sprintf(fileFormat, p.ID, p.Title, p.Link(), questionDescription)
	filename := fmt.Sprintf("%s/README.md", p.Dir())
	util.WriteFile(filename, content)
}
