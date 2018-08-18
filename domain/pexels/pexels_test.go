package pexels

import "testing"

func TestGetPexelsImageLink(t *testing.T) {

	tests := []struct {
		url string
	}{
		{
			url: "https://www.pexels.com/?dark=false&page=1",
		},
	}

	for _, test := range tests {
		getPexelsImageLink(test.url)
	}

}

func TestGetHeightWidth(t *testing.T) {
	tests := []struct {
		value string
	}{
		{
			value: "https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260",
		},
		{
			value: "https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&cs=tinysrgb&h=650&w=940",
		},
		{
			value: "https://images.pexels.com/photos/380330/pexels-photo-380330.jpeg?auto=compress&cs=tinysrgb&fit=crop&h=1200&w=800",
		},
	}
	for _, test := range tests {
		getHeightWidth(test.value)
	}
}
