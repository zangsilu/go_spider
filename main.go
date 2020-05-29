package main

import (
	"fmt"
	"regexp"
	engin "spiderProject/engine"
	"spiderProject/save"
	"spiderProject/scheduler"
	"spiderProject/zhenai/parser"
)

func main() {

	//单任务版
	//engin.SimpleEngine{}.Run(engin.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParserCityList,
	//	//Url:        "http://www.zhenai.com/zhenghun/huairou",
	//	//ParserFunc: parser.ParserCity,
	//})

	//并发版
	engin.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
		ItemChan:    save.ItemSave(),
	}.Run(engin.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}

func getCity(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[\w]+?)" data-v-2cb5b6a2>(.+?)</a>`)
	matchList := re.FindAllSubmatch(contents, -1)
	for _, v := range matchList {
		fmt.Printf("%s:%s", v[2], v[1])
	}
}

func match() {
	str := "this is email zhangshilu@theduapp.com zhangshilu2@theduapp.com"
	re := regexp.MustCompile(`([\w]+)@[\w]+.com`)
	res := re.FindAllStringSubmatch(str, -1)
	fmt.Println(res)
}
