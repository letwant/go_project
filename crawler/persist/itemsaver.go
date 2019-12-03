package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {

	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}

func save(item interface{})  {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
}