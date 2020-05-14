package parser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	engin "spiderProject/engine"
	"strings"
)

const AgeRe = `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[2]`

func parserProfile(contents []byte) engin.ParserResult {
	root, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	res := htmlquery.FindOne(root, AgeRe)
	fmt.Println("res", htmlquery.InnerText(res))

	return engin.ParserResult{}
}
