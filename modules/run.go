package modules

import (
	"jsfinder/common"
	"log"
)

func Run() {
	if common.FilePath == "" {
		Result, err := GetResponseData(common.URL)
		if err != nil {
			log.Fatalln(err)
		} else {
			common.Results = append(common.Results, *Result)
			ExportCsv()
		}
	} else {

	}
}
