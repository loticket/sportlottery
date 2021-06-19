package sportlottery

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/loticket/sportlottery/encry"
	"sort"
	"strings"
	"sync"
	"time"
)

type Request struct {
	appSecret string
	sinSecret string
	timestamp int64 //请求时间戳
	nonce     string
	signature string
	param     map[string]string
}

var RequestPool sync.Pool =  sync.Pool{
	New: func() interface{} {
		return &Request{}
	},
}

//实例化类型
func NewRequest(appSecret string,sinSecret string) *Request {
	return &Request{
		appSecret:appSecret,
		sinSecret:sinSecret,
	}
}


//获取请求的数据
func (c *Request) PostBosyJson(param map[string]string) (string,error)  {
	c.param = param
	c.nonce = encry.RandString(16)
	c.timestamp = time.Now().UnixNano() / 1e6
	c.sign()
	body,err := c.singBody()
	if err != nil {
		return "", err
	}

	var bodyMap map[string]string = map[string]string{
		"data":body,
	}

	bodyStr,_ := json.Marshal(&bodyMap)

    return string(bodyStr),nil
}


//获取公共加密数据
func (c *Request) sign() {
	var params []string = make([]string,0)
	params = append(params,fmt.Sprintf("nonce%s",c.nonce))
	params = append(params,fmt.Sprintf("timestamp%d",int(c.timestamp)))
    for k,v := range c.param {
		params = append(params,fmt.Sprintf("%s%s",k,v))
	}

	sort.Strings(params)
    var encryStr string = strings.Join(params,"")

    c.signature = encry.HmacSha1(c.sinSecret,encryStr)
	return
}

func (c *Request) singBody() (string,error) {
	var params []string = make([]string,0)
	params = append(params,fmt.Sprintf("nonce=%s",c.nonce))
	params = append(params,fmt.Sprintf("timestamp=%d",int(c.timestamp)))
	params = append(params,fmt.Sprintf("signature=%s",c.signature))
	for k,v := range c.param {
		params = append(params,fmt.Sprintf("%s=%s",k,v))
	}

	sort.Strings(params)
	var encryStr string = strings.Join(params,"&")
	decStr,err := encry.DesECBEncrypt([]byte(encryStr),[]byte(c.appSecret))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(decStr),nil
}






