package modules

import (
	"fmt"
	"jsfinder/common"
)

func Run() {
	if common.FilePath == "" {
		Result, err := GetResponseData(common.URL)
		common.Results = append(common.Results, *Result)
		fmt.Println(Result)
		fmt.Println(err)
		ExportCsv()
	} else {

	}
}
