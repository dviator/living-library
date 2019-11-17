package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cloudflare/cfssl/log"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/sirupsen/logrus"
)

func main() {
	config := GetConfig()
	defaultClient := retryablehttp.NewClient()
	client := &retryablehttp.Client{
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
		//Pass the standard logrus logger here, prints every message at Info level,
		//but integrates with the rest of the program's log settings
		Logger:     logrus.StandardLogger(),
		RetryMax:   0,
		Backoff:    defaultClient.Backoff,
		CheckRetry: defaultClient.CheckRetry,
	}
	url := "https://www.goodreads.com/search.xml?key=" + config.GoodreadsDeveloperKey + "&q=Ender%27s+Game"
	resp, _ := client.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info(string(body))
}
