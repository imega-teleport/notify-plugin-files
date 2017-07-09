package main // import "github.com/imega-teleport/notify-plugin-files"

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/imega-teleport/notify-plugin-files/fileman"
	"github.com/imega-teleport/notify-plugin-files/sender"
)

func main() {
	pluginUrl := flag.String("url", "", "Set url to connect plugin")
	storageUrlStr := flag.String("storageUrl", "", "Set storage url")
	timeout := flag.Int("timeout", 60, "timeout connection (seconds)")
	user := flag.String("user", "", "User auth")
	pass := flag.String("pass", "", "pass auth")
	path := flag.String("path", "", "path")
	flag.Parse()

	storageUrl, err := url.Parse(*storageUrlStr)
	if err != nil {
		fmt.Printf("Cound not parse storage url: %s", err)
		os.Exit(1)
	}

	/*
		1. найти все файлы по указанному пути и высчитать контрольную сумму каждого
		2. сформировать json с файлами и суммами
		3. Отправить json на указанный url
	*/

	fm := fileman.NewFileMan()
	files, err := fm.Search(*path)
	if err != nil {
		fmt.Printf("File not found: %s", err)
		os.Exit(1)
	}

	items := []sender.File4send{}

	for _, v := range files {
		sum, err := fm.Calculate(v)
		if err != nil {
			fmt.Printf("Cound not to calculate sum of file: %s", err)
			os.Exit(1)
		}

		u := url.URL{
			Scheme: storageUrl.Scheme,
			Host:   storageUrl.Host,
			Path:   fmt.Sprintf("%s/%s", storageUrl.Path, v.Name()),
		}
		item := sender.File4send{
			Url: sender.FileUrl{u},
			Sum: sum,
		}

		items = append(items, item)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   time.Duration(*timeout) * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).Dial,
			TLSHandshakeTimeout: 2 * time.Second,
		},
	}

	s := sender.NewSender(client, sender.UrlAndAuth{
		Url:  *pluginUrl,
		User: *user,
		Pass: *pass,
	})
	err = s.Send(items)
	if err != nil {
		fmt.Printf("Cound not send files: %s", err)
		os.Exit(1)
	}
}
