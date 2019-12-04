package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"go_project/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	info := model.Profile{
		Age:        "34",
		Height:     "162",
		Weight:     "57",
		Income:     "3001-5000元",
		Gender:     "女",
		Name:       "安静的雪",
		Xinzuo:     "牡羊座",
		Occupation: "人事/行政",
		Marriage:   "离异",
		House:      "已购房",
		Hukou:      "山东菏泽",
		Education:  "大学本科",
		Car:        "未购车",
	}
	id, err := save(info)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false), elastic.SetURL("http://192.168.126.130:9200"))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal([]byte(resp.Source), &actual)
	if err != nil {
		panic(err)
	}
	if actual != info {
		t.Errorf("got %v; expected %v", actual, info)
	}
}
