package parser

import (
	"spiderProject/fetcher"
	"testing"
)

const TestUrl = `http://www.zhenai.com/zhenghun`

func TestParserCityList(t *testing.T) {

	_, err := fetcher.Fetch(TestUrl)
	if err != nil {
		t.Errorf("获取错误")
		return
	}
	//t.Logf("%s", body)
}
