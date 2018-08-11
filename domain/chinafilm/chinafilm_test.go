package chinafilm

import (
	"Go-Spider/infra/url"
	"testing"
)

func TestGetChinaFilmResponse(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.ChinaFilmRealURL,
		},
	}

	for _, test := range tests {
		GetChinaFilmResponse(test.url)
	}
}
