package modules

import (
	"jsfinder/common"
	"log"
)

func FindByUrlDeep() {
	result, err := GetResponseData(common.URL)
	if err != nil {
		log.Fatalln(err)
	}
	links, _ := Extract(result.RawData, "a", "href")
	if len(links) != 0 {
		log.Printf("Find %d links\n", len(links))
	}
	for _, v := range links {
		FindByUrl(v)
	}
}
func FindByUrl(url string) {

}
