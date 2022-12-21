package modules

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"jsfinder/common"
	"log"
	"strings"
	"time"
)

func GetResponseData(Url string) (*common.Result, error) {
	result := new(common.Result)
	client := resty.New()
	client.SetHeader("User-Agent", common.UserAgent)
	client.SetHeader("cookie", common.Cookie)
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(time.Duration(common.TimeOut) * time.Second)
	resp, err := client.R().Get(Url)
	if err != nil {
		return nil, err
	}
	result.API = CheckIsApi(resp.Body())
	result.HtmlContent = string(resp.Body())
	result.URL = Url
	result.BodySize = resp.Size()
	result.Code = resp.StatusCode()
	result.RawData = resp.Body()
	result.Title, _, _ = Extract(resp.Body(), "title", "")
	if result.Title == "" {
		result.Title = "None"
	}
	return result, nil
}
func Extract(RawData []byte, Tag string, Attr string) (string, string, bool) {
	targetAttr := ""
	InnerText := ""
	exists := false
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(RawData)))
	if err != nil {
		log.Fatalln(err)
	}
	target := doc.Find(Tag)
	if Attr == "" {
		InnerText = target.Text()
	} else {
		targetAttr, exists = target.Attr(Attr)
	}
	return InnerText, targetAttr, exists
}
func CheckIsApi(RawData []byte) bool {
	if string(RawData[:2]) == "{\"" {
		return true
	}
	return false
}
