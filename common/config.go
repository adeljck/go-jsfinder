package common

import "time"

var (
	FilePath   string
	URL        string
	Cookie     string
	Thread     int
	DeepFind   bool
	JSFile     bool
	CsvOutPut  bool
	HtmlOutPut bool
	Version    string = "0.0.1"
	MaxCPUProc int
	UserAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.54"
	TimeOut    int
	Results    []Result = make([]Result, 0)
	RunTime             = time.Now()
	SubDomains []string = make([]string, 0)
)

type Result struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Code        int    `json:"code"`
	BodySize    int64  `json:"bodySize"`
	API         bool   `json:"api"`
	RawData     []byte
	HtmlContent string `json:"htmlContent,omitempty"`
}
