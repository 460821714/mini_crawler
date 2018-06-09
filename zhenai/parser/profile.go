// @Time : 2018/5/26 17:39
// @Author : minigeek
package parser

import (
	"mini_crawler/engine"
	"mini_crawler/model"
	"regexp"
	"strconv"
)

// some regexp
var ageReg = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightReg = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeReg = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageReg = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationReg = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationReg = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouReg = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoReg = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span>`)
var weightReg = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var houseReg = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carReg = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var genderReg = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)

var (
	guessReg = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)
	idUrlReg = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

// ParseProfile returns engine.ParseResult after parse profile  with contents.
func ParseProfile(contents []byte, url, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	ageStr := extractString(contents, ageReg)
	age, err := strconv.Atoi(ageStr)
	if err == nil {
		profile.Age = age
	}
	heightStr := extractString(contents, heightReg)
	height, err := strconv.Atoi(heightStr)
	if err == nil {
		profile.Height = height
	}

	profile.Income = extractString(contents, incomeReg)
	profile.Marriage = extractString(contents, marriageReg)
	profile.Education = extractString(contents, educationReg)
	profile.Occupation = extractString(contents, occupationReg)
	profile.Hokou = extractString(contents, hukouReg)
	profile.Xinzuo = extractString(contents, xinzuoReg)
	weightStr := extractString(contents, weightReg)
	weight, err := strconv.Atoi(weightStr)
	if err == nil {
		profile.Weight = weight
	}
	profile.House = extractString(contents, houseReg)
	profile.Car = extractString(contents, carReg)
	profile.Gender = extractString(contents, genderReg)

	result := engine.ParseResult{}
	result.Items = append(result.Items, engine.Item{
		Url:     url,
		Type:    "profile",
		Id:      extractString([]byte(url), idUrlReg),
		Payload: profile,
	})

	matchers := guessReg.FindAllSubmatch(contents, -1)
	for _, v := range matchers {
		name := string(v[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(v[1]),
			ParseFunc: ProfileParse(name),
		})
	}
	return result
}

// extractString extract string from contents by reg and then return.
func extractString(contents []byte, reg *regexp.Regexp) string {
	matchers := reg.FindSubmatch(contents)
	if matchers != nil {
		return string(matchers[1])
	}
	return ""
}

func ProfileParse(name string) func([]byte, string) engine.ParseResult {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
