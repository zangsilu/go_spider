package parser

import (
	"regexp"
	engin "spiderProject/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[\w]+?)" data-v-2cb5b6a2>(.+?)</a>`

func ParserCityList(contents []byte) engin.ParserResult {

	re := regexp.MustCompile(cityListRe)
	matchList := re.FindAllSubmatch(contents, -1)

	result := engin.ParserResult{}
	for _, v := range matchList {
		result.Items = append(result.Items, "City:"+string(v[2]))
		result.Requests = append(
			result.Requests, engin.Request{
				Url:        string(v[1]),
				ParserFunc: ParserCity,
			})
	}

	return result
}
