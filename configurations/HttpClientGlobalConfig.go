package configurations

import (
	cons "http-client/constants"
	"net"
	"net/http"
	"time"

)

func HttpClientCustomConfig( ) *http.Client {
   	tr:=&http.Transport{
   		ResponseHeaderTimeout: cons.ResponseHeader*time.Second,
   		Proxy: http.ProxyFromEnvironment,
   		DialContext: (&net.Dialer{
   			KeepAlive: cons.ConnKeepAlive*time.Second,
   			Timeout: cons.Connect* time.Second,
   			}).DialContext,
   			MaxIdleConns: cons.MaxAllIdleConns,
   			IdleConnTimeout: cons.IdleConn*time.Second,
   			TLSHandshakeTimeout: cons.TLSHandshake*time.Second,
   			MaxIdleConnsPerHost: cons.MaxHostIdleConns,
   			ExpectContinueTimeout: cons.ExpectContinue*time.Second,


	}

   	return &http.Client{
   		Transport: tr,

	}
}
