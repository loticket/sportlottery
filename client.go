package sportlottery

import (
	"errors"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"net/http"
)

type Client struct {
	appSecret string
	sinSecret string
	appId     string
	Httpclient *httpclient.HttpClient
}

//发起网路请求
func  (c *Client) doPostJson(param map[string]string,apiUrl string) ([]byte,error) {
	var request *Request = NewRequest(c.appSecret,c.sinSecret)
    var (
    	bodyStr string
    	err error
    )

	if bodyStr,err = request.PostBosyJson(param);err != nil {
       return nil,err
	}

	var postUrl string = fmt.Sprintf("%s%s/%s/",BaseUrlDev,apiUrl,VARSION)

    req,errs := c.Httpclient.WithHeaders(c.header()).WithOption(httpclient.OPT_TIMEOUT, 4).PostJson(postUrl,bodyStr)
    if errs != nil {
       return nil, errs
    }

    if req.StatusCode != http.StatusOK {
		return nil,errors.New("请求失败")
	}

	reqStr,errReq := req.ReadAll()

	if errReq != nil {
		return nil, errReq
	}

	return reqStr,nil

}

func (c *Client) header () map[string]string {
	return map[string]string{
		"Content-Type":"application/json",
		"appId":c.appId,
	}
}

func NewClient(appSecret string,sinSecret string,appId string) *Client {
	return &Client{
		appSecret:appSecret,
		sinSecret:sinSecret,
		appId:appId,
		Httpclient: httpclient.NewHttpClient(),
	}
}


