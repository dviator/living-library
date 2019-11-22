package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cfssl/log"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/sirupsen/logrus"
)

func main() {
	requestBooks()
	// My ID: 105516703
}

func requestBooks() {
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

	// "https://www.goodreads.com/review/list/105516703.xml?key=Jt1pNyo35UbSkSlSTnA2Sg&v=2"
	base, err := url.Parse("https://www.goodreads.com/review/list/")
	if err != nil {
		log.Error("error encoding URL: ", err)
	}
	params := url.Values{}

	params.Add("key", config.GoodreadsDeveloperKey)
	params.Add("v", "2")
	params.Add("id", "105516703")
	base.RawQuery = params.Encode()
	print("base URL: ", base)
	req, err := retryablehttp.NewRequest("GET", base.String(), nil)
	if err != nil {
		log.Error("request err: ", err)
	}
	print(req)
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	data := GoodreadsResponse{}
	err = xml.Unmarshal(body, &data)
	if err != nil {
		log.Error("couldn't unmarshal xml: ", err)
	}
	log.Info("raw body: ", string(body))
	log.Info(data)
}

type GoodreadsResponse struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Request string
	Reviews []Review `xml:"reviews>review"`
	Shelves []Shelf  `xml:"shelves>shelf`
}

type Review struct {
	XMLName xml.Name `xml:"review"`
	ID      string   `xml:"id"`
	Book    Book     `xml:"book"`
}

type Book struct {
	XMLName   xml.Name `xml:"book"`
	Title     string   `xml:"title"`
	Image_url string   `xml:"image_url"`
}

type Shelf struct {
	XMLName xml.Name `xml:"shelf`
	//XML tags and shit
}
