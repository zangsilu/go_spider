package save

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSave() chan interface{} {
	out := make(chan interface{})
	go func() {
		//保存
		count := 0
		for {
			item := <-out
			log.Printf("Got Save item #%d : %v", count, item)
			count++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) {

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	response, err := client.Index().Index("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("response", response)

}
