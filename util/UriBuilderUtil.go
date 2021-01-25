package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/url"
)

type query struct {
	key   string
	value string
}

func UriBuilder(protocol string, host string, port string, path string, qry []*query)  string {
	urlstring := fmt.Sprintf("%s://%s:%s", protocol, host, port)
	baseurl, err := url.Parse(urlstring)
	if err != nil {
		logrus.Panic("Bad formed URL", err)
		return ""
	}
	baseurl.Path = path
	if qry != nil {
		params := url.Values{}
		for i := 0; i < len(qry); i++ {
			params.Add(qry[i].key, qry[i].value)
		}
		baseurl.RawQuery= params.Encode()
	}
	return baseurl.String()
}
