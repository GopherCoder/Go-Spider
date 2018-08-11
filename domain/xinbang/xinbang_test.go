package xinbang

import (
	"Go-Spider/infra/url"
	"fmt"
	"testing"
)

func TestGetXinBangReport(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.XinBangURL,
		},
	}

	for _, test := range tests {
		fmt.Println(test.url)
		GetXinBangReport(test.url)
	}
}
