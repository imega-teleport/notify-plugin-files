package sender

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// Sender is the interface that wraps the basic Send method
type Sender interface {
	Send([]File4send) (err error)
}

type FileUrl struct {
	url.URL
}

func (f FileUrl) MarshalJSON() ([]byte, error) {
	ret, err := json.Marshal(f.String())
	return ret, err
}

type File4send struct {
	Url FileUrl `json:"url"`
	Sum string  `json:"sum"`
}

type UrlAndAuth struct {
	Url  string
	User string
	Pass string
}

type sender struct {
	client *http.Client
	opt    UrlAndAuth
}

func NewSender(client *http.Client, opt UrlAndAuth) Sender {
	return &sender{
		client: client,
		opt:    opt,
	}
}

func (s *sender) Send(items []File4send) (err error) {
	u, err := url.Parse(s.opt.Url)
	if err != nil {
		return err
	}

	b, err := json.Marshal(items)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	req.SetBasicAuth(s.opt.User, s.opt.Pass)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Teleport/1.0")

	_, err = s.client.Do(req)

	return
}
