package nuomi

import (
	"Go-Spider/infra/url"
	"testing"
)

func TestGetRankNuoMiInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.NuoMiRankURL,
		},
	}

	for _, test := range tests {
		GetRankNuoMiInfo(test.url)
	}
}
