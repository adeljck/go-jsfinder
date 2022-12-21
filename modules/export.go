package modules

import (
	"bufio"
	"fmt"
	"jsfinder/common"
	"log"
	"os"
)

func ExportCsv() {
	filename := fmt.Sprintf("ExportResults_%d_%s_%d_%d.csv", common.RunTime.Year(), common.RunTime.Month().String(), common.RunTime.Day(), common.RunTime.Unix())
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	fmt.Fprintln(writer, "url,title,status_code,body_size,is_api")
	for _, v := range common.Results {
		line := fmt.Sprintf("%s,%s,%d,%d,%v", v.URL, v.Title, v.Code, v.BodySize, v.API)
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
	fmt.Println("Scan Result Export To File " + filename)
}
