package service

import (
	"short_url/tool"
	"testing"
)

var (
	shortService ShortService
)

func BenchmarkName(b *testing.B) {
	tool.RedisInit()
	for i := 0; i < b.N; i++ {
		shortService.CreateShortURL("http://github.com")
	}
}

func TestShortService_CreateShortURL(t *testing.T) {
	tool.RedisInit()
	shortService.CreateShortURL("http://github.com")
}
