package build

import (
	"fmt"
	"log"
	"strings"

	"github.com/yangwenmai/learning-algorithms/leetcode"
)

func parseFunction(fc string) (fcName, para, ansType, nfc string) {
	log.Println("准备分解函数:", fc)

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fcName = "myFunc"
			para = "p int"
			ansType = "int"
			nfc = "func myFunc(p int) int {\n\n}"
		}
	}()

	funcIndex := strings.Index(fc, "func ")
	a := funcIndex + strings.Index(fc[funcIndex:], " ")
	b := funcIndex + strings.Index(fc[funcIndex:], "(")
	c := funcIndex + strings.Index(fc[funcIndex:], ")")
	d := funcIndex + strings.Index(fc[funcIndex:], "{")

	fcName = fc[a+1 : b]
	para = fc[b+1 : c]
	ansType = strings.TrimSpace(fc[c+1 : d])
	nfc = fmt.Sprintf("func %s(%s) %s {\n\n}", fcName, para, ansType)
	return
}

func getFunction(detail *leetcode.Question, lang string) string {
	for _, v := range detail.CodeDefinition {
		if v.Value == lang {
			return v.DefaultCode
		}
	}

	return ""
}
