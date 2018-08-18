package githubtrending

import (
	"testing"
)

func TestGithubTrendingContent(t *testing.T) {
	for _, one := range []string{"python", "go"} {
		for _, since := range []string{"day", "weekly", "month"} {
			TrendingContent(one, since)
		}
	}

}
