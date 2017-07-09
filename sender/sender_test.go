package sender

import (
	"net/http"
	"net/url"
	"testing"
)

func TestSender_Send(t *testing.T) {
	s := NewSender(&http.Client{}, UrlAndAuth{})
	s.Send([]File4send{
		{
			Url: FileUrl{url.URL{
				Host: "sdfsdf.ru",
			}},
			Sum: "sdf",
		},
		{
			Url: FileUrl{url.URL{
				Host: "sdfsdf.com",
			}},
			Sum: "sdf22",
		},
	})
}
