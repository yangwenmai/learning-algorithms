package build

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/yangwenmai/learning-algorithms/util"

	"github.com/yangwenmai/learning-algorithms/leetcode"
)

type DocsBuilder struct {
	*Builder
}

func (d *DocsBuilder) Build() {
	log.Println("开始，重建 README 文档")

	renderReadmeFile(d.LC)
	renderFavoriteFile(d.LC)

	log.Println("完成，重建 README 文档")
}

func renderReadmeFile(lc *leetcode.Leetcode) {
	file := "README.md"
	os.Remove(file)

	var b bytes.Buffer
	tmpl := template.Must(template.New("readme").Parse(readTMPL("template.markdown")))
	err := tmpl.Execute(&b, lc)
	if err != nil {
		log.Fatal(err)
	}

	// 保存 README.md 文件
	util.WriteFile(file, string(b.Bytes()))
}

func renderFavoriteFile(lc *leetcode.Leetcode) {
	file := "Favorite.md"
	os.Remove(file)

	var b bytes.Buffer
	tmpl := template.Must(template.New("favorite").Parse(readTMPL("favorite.markdown")))
	err := tmpl.Execute(&b, lc)
	if err != nil {
		log.Fatal(err)
	}

	// 保存 Favorite.md 文件
	util.WriteFile(file, string(b.Bytes()))
}

func readTMPL(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
