package main

import (
	"flag"
	"log"
	"os"

	"github.com/yangwenmai/learning-algorithms/config"
	"github.com/yangwenmai/learning-algorithms/leetcode"

	"github.com/yangwenmai/learning-algorithms/build"
)

const (
	// Version leet code generator version
	Version = "2019-05-18 11:47:47 +0800 @v0.0.3"
	Compile = "2019-05-18 11:47:47 +0800 by go version go1.12 darwin/amd64"

	// Usage 程序的使用方法
	Usage = `使用方法：
	gen_leetcode readme 刷新 learning-algorithms 项目的 README.md 和 Favorite.md 文件
	gen_leetcode prepare -number N 生成第 N 题的答题文件夹
	gen_leetcode task -prefix PRE -first FIRST -last LAST
		从 LeetCode 的 [FIRST, LAST] 选取你未完成的题目，保存到 tasks.txt，格式：
		# TODO: - 770 - #hard - 43% - Basic Calculator IV
`
	banner = `
 __       _______      ___      .______      .__   __.  __  .__   __.   _______              ___       __        _______   ______   .______       __  .___________. __    __  .___  ___.      _______.
|  |     |   ____|    /   \     |   _  \     |  \ |  | |  | |  \ |  |  /  _____|            /   \     |  |      /  _____| /  __  \  |   _  \     |  | |           ||  |  |  | |   \/   |     /       |
|  |     |  |__      /  ^  \    |  |_)  |    |   \|  | |  | |   \|  | |  |  __   ______    /  ^  \    |  |     |  |  __  |  |  |  | |  |_)  |    |  |  ---|  |---- |  |__|  | |  \  /  |    |   (----
|  |     |   __|    /  /_\  \   |      /     |  .    | |  | |  .    | |  | |_ | |______|  /  /_\  \   |  |     |  | |_ | |  |  |  | |      /     |  |     |  |     |   __   | |  |\/|  |     \   \
|   ----.|  |____  /  _____  \  |  |\  \----.|  |\   | |  | |  |\   | |  |__| |          /  _____  \  |   ----.|  |__| | |   --   | |  |\  \----.|  |     |  |     |  |  |  | |  |  |  | .----)   |
|_______||_______|/__/     \__\ | _|  ._____||__| \__| |__| |__| \__|  \______|         /__/     \__\ |_______| \______|  \______/  | _|  ._____||__|     |__|     |__|  |__| |__|  |__| |_______/
`
)

func main() {
	log.Printf("Hi, %s. I'm LeetCode generator!\n", config.GetConfig().Username)
	log.Printf("Let's start to learning algorithms!\n")
	log.Printf(banner)
	log.Printf("Git commit:%s\n", Version)
	log.Printf("Build time:%s\n", Compile)

	cli := new(CommandClient)
	cli.Run()
}

// CommandClient command client
type CommandClient struct{}

func (c *CommandClient) printUsage() {
	log.Printf("%s\n", Usage)
}

func (c *CommandClient) checkArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (c *CommandClient) Run() {
	// c.checkArgs()

	readmeCmd := flag.NewFlagSet("readme", flag.ExitOnError)

	prepareCmd := flag.NewFlagSet("prepare", flag.ExitOnError)
	dir := prepareCmd.String("dir", "practive/leetcode", "请输入你想要准备的题目的目录")
	problemNumber := prepareCmd.Int("number", 0, "请输入你想要准备的题目的题号")
	language := prepareCmd.String("lang", "golang", "请输入你想要准备的语言")

	taskCmd := flag.NewFlagSet("task", flag.ExitOnError)
	prefix := taskCmd.String("prefix", "", "答题任务的前缀")
	first := taskCmd.Int("first", 0, "答题任务的最小题号")
	last := taskCmd.Int("last", 0, "答题任务的最大题号")

	lc := leetcode.NewLeetCode(config.GetConfig().Username)
	builder := &build.Builder{
		LC: lc,
	}
	switch os.Args[1] {
	case "readme":
		err := readmeCmd.Parse(os.Args[2:])
		check(err)
		docsBuilder := &build.DocsBuilder{Builder: builder}
		docsBuilder.Build()
	case "prepare":
		err := prepareCmd.Parse(os.Args[2:])
		check(err)
		build.CheckProblemNumValid(lc, *problemNumber)
		problemBuilder := &build.ProblemBuilder{
			Builder:    builder,
			Dir:        *dir,
			ProblemNum: *problemNumber,
			Language:   *language,
		}
		problemBuilder.Build()
	case "task":
		err := taskCmd.Parse(os.Args[2:])
		check(err)
		taskBuilder := &build.TaskBuilder{
			Builder: builder,
			First:   *first,
			Last:    *last,
			Prefix:  *prefix,
		}
		taskBuilder.Build()
	default:
		c.printUsage()
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
