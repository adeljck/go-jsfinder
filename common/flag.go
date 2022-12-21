package common

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func Banner() {
	banner := `
  ________                      ____. _________
 /  _____/  ____               |    |/   _____/
/   \  ___ /  _ \   ______     |    |\_____  \ 
\    \_\  (  <_> ) /_____/ /\__|    |/        \
 \______  /\____/          \________/_______  /
        \/                                  \/
                     Go-JSFinder version: ` + Version + `
`
	print(banner)
}

func ParseParma() {
	Banner()
	flag.Usage = myUsage
	flag.StringVar(&FilePath, "f", "", "The File Contains Url or JS File That You Want To Scan.")
	flag.StringVar(&Cookie, "c", "", "Target Site's Cookie.")
	flag.StringVar(&URL, "u", "", "The Target Website You Want To Scan.")
	flag.BoolVar(&JSFile, "j", false, "Find In JS File,Default Is True.")
	flag.IntVar(&Thread, "t", 10, "Web Scan Thread.Default Is 10.")
	flag.BoolVar(&DeepFind, "d", false, "Deep Find Website.")
	flag.BoolVar(&CsvOutPut, "csv", false, "Export Scan Result To A Csv File.")
	flag.BoolVar(&HtmlOutPut, "html", true, "Export Scan Result To A HTML File.")
	flag.IntVar(&TimeOut, "to", 5, "Request Timeout,Default is 5 Second.")
	flag.IntVar(&MaxCPUProc, "m", runtime.NumCPU()/2, "Max CPU Nums To Run Programs,Default Is The Half Of Your Total CPU Nums.")
	flag.Parse()
	runtime.GOMAXPROCS(MaxCPUProc)
}
func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", os.Args[0])
	flag.PrintDefaults()
}
