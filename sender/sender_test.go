package sender

import (
	"testing"
	"net/http"
	"net/url"
)

func TestSender_Send(t *testing.T) {
	s := NewSender(&http.Client{}, &url.URL{})
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

