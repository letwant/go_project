package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func test() {
	resp, err := http.Get("https://m.zhenai.com/u/1525702321")
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	infoRe := `data-v-3e01facc[^>]*>([^<]`
	//infoRe := `m-btn purple`
	re := regexp.MustCompile(infoRe)
	match := re.FindAllSubmatch(body, -1)
	fmt.Println(len(match))
	for _, item := range match {
		fmt.Println(string(item[0]))
	}
}

func test2() {
	//res := `"basicInfo":\[[.+]\]`
	profileRe := regexp.MustCompile(`"basicInfo":\[(".+?)\]`)
	contents := `{"objectInfo":{"age":37,"avatarPhotoID":298523925,"avatarPraiseCount":0,"avatarPraised":false,"avatarURL":"https:\u002F\u002Fphoto.zastatic.com\u002Fimages\u002Fphoto\u002F22976\u002F91901938\u002F7835157641273300.png","basicInfo":["未婚","37岁","魔羯座(12.22-01.19)","181cm","75kg","工作地:北京朝阳区","月收入:2-5万","传媒\u002F艺术","大学本科"],"detailInfo":["汉族","籍贯:北京","体型:瘦长","不吸烟","社交场合会喝酒","已购房","已买车","没有小孩","是否想要孩子:视情况而定","何时结婚:时机成熟就结婚"],"educationString":"大学本科","emotionStatus":0,"gender":0,"genderString":"男士","hasIntroduce":true,"heightString":"181cm","hideVerifyModule":false,"introduceContent":"土生土长北京人，一直秉承宁缺毋滥的恋爱原则，所以至今未遇到那个适合结婚的人…至于性格方面，请参考O型血魔羯座性格特征，很像很典型…如果看一眼照片觉得还对路子，就加Q聊吧回头，不得不承认，第一眼的感觉还是很重要的，尤其是在网上，哪怕是照片…我说话喜欢直给，不喜欢瞎含蓄，希望你也是…","introducePraiseCount":0,"isActive":true,"isFollowing":false,"isInBlackList":false,"isStar":false,"isZhenaiMail":true,"lastLoginTimeString":"当前在线","liveAudienceCount":0,"liveType":0,"marriageString":"未婚","memberID":91901938,"momentCount":2,"nickname":"守望者","objectAgeString":"25-35岁","objectBodyString":"未填写","objectChildrenString":"没有小孩","objectEducationString":"未填写","objectHeightString":"163-175cm","objectInfo":["25-35岁","163-175cm","工作地:北京","月薪:2万以上","未婚","没有小孩"],"objectMarriageString":"未婚","objectSalaryString":"20000元以上","objectWantChildrenString":"未填写","objectWorkCityString":"北京","onlive":0,"photoCount":5,"praisedIntroduce":false,"previewPhotoURL":"","pycreditCertify":false,"salaryString":"20001-50000元","showValidateIDCardFlag":false,"totalPhotoCount":5,"validateEducation":false,"validateFace":false,"validateIDCard":true,"videoCount":0,"videoID":0,"workCity":10102005,"workCityString":"北京","workProvinceCityString":"北京朝阳区"}};(function(){var s;(s=document.currentScript||document.scripts[document.scripts.length-1]).parentNode.removeChild(s);}());`
	data := []byte(contents)
	match := profileRe.FindSubmatch(data)
	if len(match) >= 2 {
		fmt.Println(string(match[1]))
	}
	//fmt.Println(match)
}

func main() {
	//engine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	test2()
}
