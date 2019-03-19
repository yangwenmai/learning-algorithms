package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mattn/godown"
)

// FetchQuestion 参考：https://github.com/WindomZ/leetcode-graphql/blob/master/base_question.go
func FetchQuestion(p Problem) *Question {
	body := strings.NewReader(fmt.Sprintf(`{"operationName":"getQuestionDetail","variables":{"titleSlug":"%s"},"query":"query getQuestionDetail($titleSlug: String!) {\n  isCurrentUserAuthenticated\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    questionTitle\n    translatedTitle\n    questionTitleSlug\n    content\n    translatedContent\n    difficulty\n    stats\n    allowDiscuss\n    contributors {\n      username\n      profileUrl\n      __typename\n    }\n    similarQuestions\n    mysqlSchemas\n    randomQuestionUrl\n    sessionId\n    categoryTitle\n    submitUrl\n    interpretUrl\n    codeDefinition\n    sampleTestCase\n    enableTestMode\n    metaData\n    langToValidPlayground\n    enableRunCode\n    enableSubmit\n    judgerAvailable\n    infoVerified\n    envInfo\n    urlManager\n    article\n    questionDetailUrl\n    libraryUrl\n    companyTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    __typename\n  }\n  interviewed {\n    interviewedUrl\n    companies {\n      id\n      name\n      slug\n      __typename\n    }\n    timeOptions {\n      id\n      name\n      __typename\n    }\n    stageOptions {\n      id\n      name\n      __typename\n    }\n    __typename\n  }\n  subscribeUrl\n  isPremium\n  loginUrl\n}\n"}`, p.TitleSlug))
	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", body)
	if err != nil {
		panic(err)
	}
	refer := fmt.Sprintf(
		"https://leetcode.com/problems/%s/description/",
		p.TitleSlug,
	)
	req.Header.Set("x-csrftoken", csrfTokenBase)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("referer", refer)
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	req.AddCookie(&http.Cookie{
		Name:    "csrftoken",
		Value:   csrfTokenBase,
		Path:    "/",
		Domain:  ".leetcode.com",
		Secure:  true,
		Expires: time.Now(),
	})
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var raw struct {
		Data struct {
			Question rawQuestion
		}
	}

	if err = json.Unmarshal(data, &raw); err != nil {
		panic(err)
	}

	pdetail, err := detailFromRaw(raw.Data.Question)
	if err != nil {
		panic(err)
	}
	return pdetail
}

type rawQuestion struct {
	Content           string
	Stats             string
	CodeDefinition    string
	SampleTestCase    string
	EnableRunCode     bool
	MetaData          string
	TranslatedContent string
}

type Question struct {
	Content           string
	Stats             questionStat
	CodeDefinition    []codeDefinition
	SampleTestCase    string
	MetaData          string
	TranslatedContent string
	EnableRunCode     bool
}

type codeDefinition struct {
	Value       string
	Text        string
	DefaultCode string
}

type questionStat struct {
	TotalAccepted      string
	TotalSubmission    string
	TotalAcceptedRaw   int
	TotalSubmissionRaw int
	AcRate             string
}

func detailFromRaw(raw rawQuestion) (*Question, error) {
	d := &Question{}
	buf := &bytes.Buffer{}

	if err := godown.Convert(buf, strings.NewReader(raw.Content), nil); err != nil {
		return nil, err
	}
	d.Content = buf.String()
	buf.Reset()

	var c []codeDefinition
	if err := json.Unmarshal([]byte(raw.CodeDefinition), &c); err != nil {
		return nil, err
	}
	d.CodeDefinition = c

	var stats questionStat
	if err := json.Unmarshal([]byte(raw.Stats), &stats); err != nil {
		return nil, err
	}

	d.Stats = stats
	return d, nil
}
