package cbooo

import (
	"Go-Spider/infra/url"
	"fmt"
	"testing"
)

func TestGetCboRankList(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: "",
		},
	}

	for _, test := range tests {
		fmt.Println(test)
		getCboRankList()

	}
}

func TestGetAreaList(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.CboURL,
		},
	}

	for _, test := range tests {
		getAreaList(test.url)
	}
}

func TestGetRealTimeList(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.CboRootURL,
		},
	}

	for _, test := range tests {
		getRealTimeList(test.url)
	}
}
