package paser

import (
	"encoding/json"
	"errors"
	"strings"
)

type PlsOpen struct {
	Api  string `json:"api"`
	Data map[string]string `json:"data"`
	Ret  []string `json:"ret"`
	V    string `json:"v"`
}

func (p *PlsOpen) PaserConent() (error,interface{}){
	var content string
	var flag bool
	if content,flag = p.Data["result"];!flag{
		return errors.New("pls paser result err"),nil
	}

	//解析content的内容
	var lotteryContent LotteryContent

	if err := json.Unmarshal([]byte(content),&lotteryContent);err != nil {
		return err,nil
	}

	if lotteryContent.Header.Status == -1 {
		return errors.New(lotteryContent.ErrMsg),nil
	}

	var resutle PlsResult

	docoder := json.NewDecoder(strings.NewReader(string(lotteryContent.Content)))

	docoder.UseNumber()

	if err := docoder.Decode(&resutle);err != nil {
		return err,nil
	}

	return nil,resutle
}