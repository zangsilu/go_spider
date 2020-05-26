package engine

import (
	"log"
	"spiderProject/fetcher"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	//将种子放入请求队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		ParserResult, _ := e.worker(r)

		requests = append(requests, ParserResult.Requests...)

		for _, item := range ParserResult.Items {
			log.Printf("item : %v", item)
		}
	}
}

func (e SimpleEngine) worker(r Request) (ParserResult, error) {
	log.Printf("Fetch : %v", r.Url)

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s:%v", r.Url, err)
		return ParserResult{}, err
	}
	ParserResult := r.ParserFunc(body)
	return ParserResult, err
}
