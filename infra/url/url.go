package url

import (
	"fmt"
	"time"
)

var (
	PFangURL         = "http://www.piaofang.biz/"                                           // 正则表达式范例
	MaoYanURLByYear  = "https://piaofang.maoyan.com/rankings/year?year=%d&limit=100&tab=%d" // 正则表达式范例
	CboURL           = "http://www.cbooo.cn/movies"
	CboRealURL       = "http://www.cbooo.cn/Mdata/getMdata_movie?area=%s&type=0&year=0&initial=全部&pIndex=%d"
	CboRootURL       = "http://www.cbooo.cn/"
	ChinaFilmURL     = "http://i.chinafilm.com/index/Movie/movielist"
	ChinaFilmRealURL = "http://i.chinafilm.com/index/Movie/ajaxList"
	NuoMiURL         = "https://dianying.nuomi.com/movie/boxoffice?cityId=289"
	NuoMiRealURL     = "https://mdianying.baidu.com/api/discovery/ranking"
	NuoMiRankURL     = "https://mdianying.baidu.com/api/rank/rankPage"
	XinBangURL       = "https://www.newrank.cn/xdnphb/list/day/rank"
	MeiZiURL         = "http://www.mmjpg.com/"
	GuShiWenURL      = "https://www.gushiwen.org/"
	GithubTrend      = "https://github.com/trending/go?since=daily"
	PexlesURL        = "https://www.pexels.com/?dark=false&page=%d"
)

func GetMaoYanURL(year int) string {
	nowYear := time.Now().Year()
	tab := nowYear - year + 1
	return fmt.Sprintf(MaoYanURLByYear, year, tab)
}
