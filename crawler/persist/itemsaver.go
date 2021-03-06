package persist

import (
	"context"
	"github.com/olivere/elastic"
	"go_project/crawler/model"
	"log"
)

func ItemSaver() (chan model.Profile, error) {

	out := make(chan model.Profile)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			//_, err := save(item)
			//if err != nil {
			//	log.Printf("Item Saver: error saving item %v: %v", item, err)
			//}
		}
	}()
	return out, nil
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false), elastic.SetURL("http://192.168.126.130:9200"))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}