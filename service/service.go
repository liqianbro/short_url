package service

import (
	"encoding/json"
	"log"
	"short_url/model"
	"short_url/tool"
	"time"
)

type ShortService struct{}

var (
	redisKey = "shortURL"
)

// CreateShortURL 生成短链接
func (s *ShortService) CreateShortURL(url string) (result bool, redirect model.Redirect) {
	// SET Domain
	redirect.GetDomain()
	// 生成短链接
	code := tool.Generator(tool.CHARSET_RANDOM_ALPHANUMERIC, redirect.Domain, url)
	// 进行实体封装
	redirect.Code = code
	redirect.ShortURL = redirect.Domain + redirect.Code
	redirect.URL = url
	redirect.CreatedAt = time.Now()
	// redis 进行缓存
	redirectStr, err := json.Marshal(redirect)
	if err != nil {
		log.Printf("err:%e", err)
		return
	}
	ok, _ := tool.RedisClient.HSet(redisKey, redirect.Code, redirectStr).Result()
	if !ok {
		log.Printf("错误信息:%s", "save redis fail")
		return
	}
	result = true
	return
}

// RedirectURL 301 短链接重定向
func (s *ShortService) RedirectURL(code string) (url string) {
	resultJson, redisErr := tool.RedisClient.HGet(redisKey, code).Result()
	if redisErr != nil {
		log.Fatalf("Redis查询错误：%e", redisErr)
		return
	}
	var redirect model.Redirect
	jsonErr := json.Unmarshal([]byte(resultJson), &redirect)
	if jsonErr != nil {
		log.Fatalf("JSON解析错误：%e", jsonErr)
		return
	}
	url = redirect.URL
	return
}
