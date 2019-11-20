package parser

import (
	"go_project/crawler/engine"
	"go_project/crawler/model"
	"regexp"
	"strings"
)

//var ageRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([\d])+岁</div>`)
//var marriageRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([^<]+)</div>`)
var basicInfoRe = regexp.MustCompile(`"basicInfo":\[(".+?)\]`)
var detailInfoRe = regexp.MustCompile(`"detailInfo":\[(".+?)\]`)
var sexInfoRe = regexp.MustCompile(`"gender":([0-9])`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.House = ""
	profile.Car = ""
	sexInfo := extractString(contents, sexInfoRe)
	if sexInfo == "1" {
		profile.Gender = "女"
	}else {
		profile.Gender = "男"
	}
	basicInfo := extractString(contents, basicInfoRe)
	basicInfo = strings.Replace(basicInfo, `"`, ``, -1)
	basicInfoSlice := strings.Split(basicInfo, ",")
	detailInfo := extractString(contents, detailInfoRe)
	detailInfoSlice := strings.Split(detailInfo, ",")
	for _, basic := range detailInfoSlice {
		if strings.Contains(basic, "房") {
			profile.House = strings.Trim(basic, `"`)
		}
		if strings.Contains(basic, "车") {
			profile.Car = strings.Trim(basic, `"`)
		}
	}
	if len(basicInfoSlice) < 9 {
		tmpSlice := basicInfoSlice[4:]
		basicInfoSlice = append(basicInfoSlice, "")
		basicInfoSlice = append(basicInfoSlice, tmpSlice...)
	}
	profile.Age =  basicInfoSlice[1]
	profile.Marriage = basicInfoSlice[0]
	profile.Xinzuo = basicInfoSlice[2]
	profile.Height = basicInfoSlice[3]
	profile.Weight = basicInfoSlice[4]
	profile.Income = basicInfoSlice[6]
	profile.Occupation = basicInfoSlice[7]
	profile.Education = basicInfoSlice[8]

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