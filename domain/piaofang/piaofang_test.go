package piaofang

import (
	"Go-Spider/infra/url"
	"fmt"
	"testing"
)

func TestGetPFangRankInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.PFangURL,
		},
	}

	for _, test := range tests {

		fmt.Println(GetPFangRankInfo(test.url))

	}
}
