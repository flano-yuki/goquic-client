package main

import (
	"fmt"
	"net/http"
	"strings"
	"net/http/httputil"

	"github.com/devsisters/goquic"
	"github.com/jteeuwen/go-pkg-optarg"
)

func init() {
}

func log(str string) {
	fmt.Printf(" " + str + "\n")
}

func main() {
	var url string
	var logLevel int
	var bodyString string
	var showBody bool
	var quiet bool
	var headers []string

	optarg.Add("l", "liblog", "libquic Log level", 2)
	optarg.Add("s", "sring", "POST body string", "")
	optarg.Add("h", "headers", "HTTP Request headers (-h key=value)", "")
	optarg.Add("b", "body", "Show body", false)
	optarg.Add("q", "quiet", "no output", false)
	//optarg.Add("o", "out", "File save (-o file)", "")
	//optarg.Add("c", "continue", "Loop http request", "")
	//optarg.Add("k", "keepalive", "do keepalive on same domain", "")
	//optarg.Add("a", "all", "get all resource", "")
	//optarg.Add("t", "tcp", "firstly get alt-svc using https", "")
	//optarg.Add("version", "version", "version", "")
	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "l":
			logLevel = opt.Int()
		case "s":
			bodyString = opt.String()
		case "h":
			headers = append(headers, opt.String())
		case "b":
			showBody = true
		case "q":
			quiet = true
		case "version":
			log(Name + " " + Version)
			return
		}
	}

	url = optarg.Remainder[0]

	goquic.SetLogLevel(logLevel)

	client := &http.Client{
		Transport: goquic.NewRoundTripper(false),
	}

	var err error
	var req *http.Request
	if bodyString == "" {
		req, err = http.NewRequest("GET", url, nil)
	} else {
		req, err = http.NewRequest("POST", url, strings.NewReader(bodyString))
	}
	if err != nil {
		panic(err)
	}

	for _, kv := range headers {
		key := strings.SplitN(kv, "=", 2)[0]
		value := strings.SplitN(kv, "=", 2)[1]
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	dumpResp, _ := httputil.DumpResponse(resp, showBody)
	if (!quiet) {
		log(string(dumpResp))
	}
}
