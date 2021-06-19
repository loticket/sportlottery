package sportlottery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/loticket/sportlottery/paser"
)

type Lottery struct {
	client *Client
}

//大乐透
func (l *Lottery) Dlt(drawNo string) (err error,res paser.DltResult) {
	var dlt paser.DltOpen = paser.DltOpen{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"drawNo":drawNo},DltApi,&dlt);err != nil || resInter == nil {
		return
	}
    var ok bool
	if res,ok = resInter.(paser.DltResult);!ok {
		err = errors.New("dlt paser fail")
        return
	}

	return
}

//排列三
func (l *Lottery) Pls(drawNo string) (err error,res paser.PlsResult)  {
	var pls paser.PlsOpen = paser.PlsOpen{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"drawNo":drawNo},PlsApi,&pls);err != nil || resInter == nil {
		return
	}

	var ok bool
	if res,ok = resInter.(paser.PlsResult);!ok {
		err = errors.New("pls paser fail")
		return
	}

	return
}

//排列五
func (l *Lottery) Plw(drawNo string) (err error,res paser.PlwResult)  {
	var pls paser.PlwOpen = paser.PlwOpen{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"drawNo":drawNo},PlwApi,&pls);err != nil || resInter == nil {
		return
	}

	var ok bool
	if res,ok = resInter.(paser.PlwResult);!ok {
		err = errors.New("plw paser fail")
		return
	}

	return
}

//七星彩
func (l *Lottery) Qxc(drawNo string) (err error,res paser.QxcResult)  {
	var qxc paser.QxcOpen = paser.QxcOpen{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"drawNo":drawNo},QxcApi,&qxc);err != nil || resInter == nil {
		return
	}

	var ok bool
	if res,ok = resInter.(paser.QxcResult);!ok {
		err = errors.New("plw paser fail")
		return
	}

	return
}

//获取正在开售的期号
//足球胜平负：90；胜平负任选9场：900129；足球4场进球：94；足球6场半全场胜平负：98。
func (l *Lottery) ZuCaiIssue(game string) (err error,res []paser.ZuCaiIssueResult)  {
	var zucai paser.ZuCaiIssue = paser.ZuCaiIssue{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"gameNo":game},ZqIssueApi,&zucai);err != nil || resInter == nil {
		return
	}

	var ok bool
	if res,ok = resInter.([]paser.ZuCaiIssueResult);!ok {
		err = errors.New("zucai issue paser fail")
		return
	}

	return
}

//获取正在开售的期号
//足球胜平负：90；胜平负任选9场：900129；足球4场进球：94；足球6场半全场胜平负：98。
func (l *Lottery) ZuCai(drawNo string,game string) (err error,res paser.ZuCaiResult)  {
	var zucai paser.ZuCai = paser.ZuCai{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"drawNo":drawNo,"gameNo":game},ZqPrizeApi,&zucai);err != nil || resInter == nil {
		return
	}

	var ok bool
	if res,ok = resInter.(paser.ZuCaiResult);!ok {
		err = errors.New("zucai paser fail")
		return
	}

	return
}

//竞彩获取赛程
//运动编号。足球：1；篮球2。
func (l *Lottery) JcMatchList(sportNo string) (err error,res paser.JcMatchResult)  {
	var zucai paser.JcMatch = paser.JcMatch{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"sportNo":sportNo},JcMatchListApi,&zucai);err != nil || resInter == nil {
		return
	}

	var ok bool
	if res,ok = resInter.(paser.JcMatchResult);!ok {
		err = errors.New("js match fail")
		return
	}

	return
}

//竞彩获取赛程
//运动编号。足球：1；篮球2。
//足球让球胜平负：1；足球半全场胜平负：2；足球比分：3；足球总进球：4，足球胜平负：10，篮球胜负：7，
//篮球让分胜负：6，篮球胜分差：8，篮球大小分：9。
func (l *Lottery) JcMatchOdds(sportNo string,poolNo string) (err error,res paser.JcOddsResult)  {
	var zucai paser.JcOdds = paser.JcOdds{}
	var resInter interface{}
	if err,resInter = l.lotCommon(map[string]string{"sportNo":sportNo,"poolNo":poolNo},JcMatchOddsApi,&zucai);err != nil || resInter == nil {
		return
	}

	var ok bool
	if res,ok = resInter.(paser.JcOddsResult);!ok {
		err = errors.New("js odds paser fail")
		return
	}

	return
}

//解析公共数据
func  (l *Lottery) lotCommon(param map[string]string,api string,lotter paser.Paser) (error,interface{}) {
	body,err := l.client.doPostJson(param,api)
	if err != nil {
		return err,nil
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body,lotter)
	if err != nil {
		return err,nil
	}

	return lotter.PaserConent()
}



func NewLottery(appSecret string,sinSecret string,appId string) *Lottery {
   return &Lottery{
       client: NewClient(appSecret,sinSecret,appId),
    }
}