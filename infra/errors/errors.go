package errors

import "strconv"

type ErrorCode struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (e *ErrorCode) Error() string {
	return "Code: " + strconv.Itoa(e.Code) + ", Detail " + e.Detail
}

var (
	ErrorRequest        = &ErrorCode{400, "http request fail "}
	ErrorResponse       = &ErrorCode{400, "http response fail"}
	ErrorStatusCode     = &ErrorCode{400, "http status code fail"}
	ErrorPFangRankInfo  = &ErrorCode{400, "piao fang rank info fail"}
	ErrorWebDriver      = &ErrorCode{400, "web driver fail"}
	ErrorSelenium       = &ErrorCode{400, "selenium web driver fail"}
	ErrorMaoYanRankList = &ErrorCode{400, "get mao yan data fail"}
	ErrorDB             = &ErrorCode{500, "operate database fail"}
)
