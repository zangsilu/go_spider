package parser

import (
	"regexp"
	engin "spiderProject/engine"
	"spiderProject/model"
	"strconv"
)

const UserListRe = `<div class="list-item">.+?<div class="item-btn">打招呼</div></div>`
const NextPageRe = `<a href="(http://www.zhenai.com/zhenghun/[\w]+/[\d]+)">下一页</a>`
const UserNameRe = `<th><a href="http://album.zhenai.com/u/[\d]+" target="_blank">(.+?)<`
const SexRe = `<td width="180"><span class="grayL">性别：</span>(.+?)</td>`
const marriageRe = `<td width="180"><span class="grayL">婚况：</span>(.+?)</td>`
const ageRe = `<td width="180"><span class="grayL">年龄：</span>(.+?)</td>`
const addressRe = `<td><span class="grayL">居住地：</span>(.+?)</td>`
const educationRe = `<td><span class="grayL">学   历：</span>(.+?)</td>`
const moneyRe = `<td><span class="grayL">月   薪：</span>(.+?)</td>`
const heightRe = `<td width="180"><span class="grayL">身   高：</span>(.+?)</td>`
const introRe = `<div class="introduce">(.+?)</div>`

func ParserCity(contents []byte) engin.ParserResult {
	result := engin.ParserResult{}

	matchNextUrl := regexp.MustCompile(NextPageRe).FindSubmatch(contents)
	if len(matchNextUrl) == 0 {
		return result
	}
	nextUrl := string(matchNextUrl[1])

	userListSlice := regexp.MustCompile(UserListRe).FindAll(contents, -1)

	profile := model.Profile{}
	for _, userInfo := range userListSlice {
		userNameSlice := regexp.MustCompile(UserNameRe).FindSubmatch(userInfo)
		userName := string(userNameSlice[1])

		sexSlice := regexp.MustCompile(SexRe).FindSubmatch(userInfo)
		marriageSlice := regexp.MustCompile(marriageRe).FindSubmatch(userInfo)
		ageSlice := regexp.MustCompile(ageRe).FindSubmatch(userInfo)
		addressSlice := regexp.MustCompile(addressRe).FindSubmatch(userInfo)
		educationSlice := regexp.MustCompile(educationRe).FindSubmatch(userInfo)
		moneySlice := regexp.MustCompile(moneyRe).FindSubmatch(userInfo)
		heightSlice := regexp.MustCompile(heightRe).FindSubmatch(userInfo)
		introSlice := regexp.MustCompile(introRe).FindSubmatch(userInfo)

		profile.Name = userName
		profile.Sex = string(sexSlice[1])
		profile.Marriage = string(marriageSlice[1])
		profile.Address = string(addressSlice[1])
		if len(educationSlice) > 0 {
			profile.Education = string(educationSlice[1])
		}
		if len(moneySlice) > 0 {
			profile.Money = string(moneySlice[1])
		}
		profile.Intro = string(introSlice[1])
		profile.Height, _ = strconv.Atoi(string(heightSlice[1]))
		profile.Age, _ = strconv.Atoi(string(ageSlice[1]))

		result.Items = append(result.Items, profile)
	}
	result.Requests = append(
		result.Requests, engin.Request{
			Url:        nextUrl,
			ParserFunc: ParserCity,
		})

	return result
}
