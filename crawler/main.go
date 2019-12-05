package main

import (
	"go_project/crawler/engine"
	"go_project/crawler/persist"
	"go_project/crawler/scheduler"
	"go_project/crawler/zhenai/parser"
)

func main() {
	saver, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 1,
		ItemChan:    saver,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
