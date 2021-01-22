package util

import (
	"bytes"
	_ "github.com/motemen/go-loghttp/global"
	config "http-client/configurations"
	"http-client/constants"
	"net/http"
)

type headers struct {
	key   string
	value string
}

func HttpRequest(url string, method string, hdrs []*headers, data []byte) (*http.Response, error) {
	var response *http.Response

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		return response, err
	}

	request.Header.Set(constants.Authorization_Constant, config.GetToken())
	if data != nil {
		request.Header.Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	}

	if hdrs != nil {
		for i := 0; i < len(hdrs); i++ {
			request.Header.Add(hdrs[i].key, hdrs[i].value)
		}
	}

	client := config.HttpClientCustomConfig()

	response, err = client.Do(request)
	return response, err
}
