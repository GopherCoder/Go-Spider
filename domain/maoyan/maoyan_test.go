package maoyan

import (
	"Go-Spider/infra/url"
	"fmt"
	"testing"
)

func TestGetMaoYanRankInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.GetMaoYanURL(2018),
		},
	}

	for _, test := range tests {
		fmt.Println(GetMaoYanRankInfo(test.url))
	}
}
