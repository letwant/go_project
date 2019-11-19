package parser

import (
	"go_project/crawler/engine"
	"go_project/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([\d])+岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([^<]+)</div>`)
var profileRe = regexp.MustCompile(`"basicInfo":\[(.+)[^]]]`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe)


}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}