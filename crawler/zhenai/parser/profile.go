package parser

import (
	"go_project/crawler/engine"
	"go_project/crawler/model"
	"regexp"
	"strings"
)

//var ageRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([\d])+岁</div>`)
//var marriageRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([^<]+)</div>`)
var profileRe = regexp.MustCompile(`"basicInfo":\[(".+?)\]`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	basicInfo := extractString(contents, profileRe)
	basicInfoList := strings.Split(basicInfo, ",")
	if len(basicInfoList) < 9 {
		tmpSlice := basicInfoList[4:]
		basicInfoList = append(basicInfoList, "")
		basicInfoList = append(basicInfoList, tmpSlice...)
	}
	profile.Age = basicInfoList[1]
	profile.Marriage = basicInfoList[0]
	profile.Xinzuo = basicInfoList[2]
	profile.Height = basicInfoList[3]
	profile.Weight = basicInfoList[4]
	profile.Income = basicInfoList[6]
	profile.Occupation = basicInfoList[7]
	profile.Education = basicInfoList[8]
	profile.House = ""
	profile.Car = ""
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}