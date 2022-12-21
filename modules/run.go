package modules

import (
	"jsfinder/common"
)

func Run() {
	if common.FilePath == "" {
		if common.DeepFind {

		} else {
			FindByUrlDeep()
		}
	} else {
		if common.JSFile {

		} else {

		}
	}
}
