package paser

type Paser interface {
	PaserConent() (error,interface{})
}

type LotHeader struct {
	ServerTime string `json:"serverTime"`
	Status     int    `json:"status"`
}


type LotteryContent struct {
	Content string                 `json:"content"`
	Header  LotHeader              `json:"header"`
	Sign    string                 `json:"sign"`
	ErrMsg  string                 `json:"errMsg"`
	ErrNo   string                 `json:"errNo"`
}

type LevelPrize struct {
	PriovinceName string `json:"priovinceName"`  //省份名
	FirstLevelStakeCount int `json:"firstLevelStakeCount"` //一等奖基本中奖注数
	FirstLevelAddStakeCount int `json:"first_level_add_stake_count"` //一等奖追加中奖注数
}


//奖等
type PrizeStat struct {
	PrizeAmount float64    `json:"prizeAmount"` //中奖金额，单位：元。
	PrizeLevel  string     `json:"prizeLevel"`  //奖等名称
	StakeAmount float64    `json:"stakeAmount"` //单注奖金，单位：元。
	StakeCount  int        `json:"stakeCount"`  //中奖注数
}

//派奖信息
type Promotion struct {
	NextValidate int `json:"nextValidate"` //下一个奖期是否有派奖  1为有，0为没有
	Validate     int `json:"validate"` //当前奖期是否有派奖  1为有，0为没有
}

type DltResult struct {
	DrawNo     string `json:"drawNo"` //奖期编号
	DrawResult string `json:"drawResult"` //开奖结果
	UnSortDrawResult  string `json:"unSortDrawResult"` //出球顺序
	DrawTime   string `json:"drawTime"`   //开奖日期
	GameName   string `json:"gameName"`   //游戏名称
	PaidEndTime      string   `json:"paidEndTime"`   //截止兑奖日期
	PoolBalance      float64  `json:"poolBalance"`   //本期开奖后奖池余额，单位：元。
	PrizeAmountSum   float64   `json:"prizeAmountSum"`   //应派奖金总计, 单位：元。
	FirstLevelPrizeInfos []LevelPrize `json:"firstLevelPrizeInfos"`  //一等奖统计信息
	LotteryEquipmentCount   int `json:"lotteryEquipmentCount"`   //摇奖球套数
	PrizeStat        []PrizeStat  `json:"prizeStat"`   //奖等
	PromotionInfo    Promotion  `json:"promotionInfo"` //派奖信息
	TotalSaleAmount  float64    `json:"totalSaleAmount"`   //全国本期投注总额

}


type PlsPrizeStat struct{
	PrizeAmount float64          `json:"prizeAmount"` //中奖金额，单位：元。
	PrizeLevel  string           `json:"prizeLevel"`  //奖等名称
	StakeAmount float64          `json:"stakeAmount"` //单注奖金，单位：元。
	StakeCount  int              `json:"stakeCount"`  //中奖注数
	provincePrizeAmount float64  `json:"provincePrizeAmount"`  //本省中奖注数
	provinceStakeCount  int      `json:"provinceStakeCount"`   //本省中奖金额, 单位：元。
}

type PlsResult struct {
	DrawNo     string `json:"drawNo"` //奖期编号
	DrawResult string `json:"drawResult"` //开奖结果
	DrawTime   string `json:"drawTime"`   //开奖日期
	GameName                string    `json:"gameName"`         //游戏名称
	PaidEndTime             string    `json:"paidEndTime"`       //截止兑奖日期
	PoolBalance             float64   `json:"poolBalance"`       //本期开奖后奖池余额，单位：元。
	PrizeAmountSum          float64   `json:"prizeAmountSum"`   //应派奖金总计, 单位：元。
	TotalSaleAmount         float64   `json:"totalSaleAmount"`   //全国本期投注总额
	ProvincePrizeAmountSum  float64   `json:"provincePrizeAmountSum"` //本省应派奖金总计，单位：元。
	ProvinceSaleAmount      float64   `json:"provinceSaleAmount"`   //本省本期投注总额, 单位：元。
	PrizeStat              []PlsPrizeStat  `json:"prizeStat"`
}

type PlwResult struct {
	DrawNo     string `json:"drawNo"` //奖期编号
	DrawResult string `json:"drawResult"` //开奖结果
	DrawTime   string `json:"drawTime"`   //开奖日期
	GameName                string    `json:"gameName"`         //游戏名称
	PaidEndTime             string    `json:"paidEndTime"`       //截止兑奖日期
	PoolBalance             float64   `json:"poolBalance"`       //本期开奖后奖池余额，单位：元。
	PrizeAmountSum          float64   `json:"prizeAmountSum"`   //应派奖金总计, 单位：元。
	TotalSaleAmount         float64   `json:"totalSaleAmount"`   //全国本期投注总额
	PrizeStat              []PlsPrizeStat  `json:"prizeStat"`
}

type QxcResult struct {
	DrawNo     string `json:"drawNo"` //奖期编号
	DrawResult string `json:"drawResult"` //开奖结果
	DrawTime   string `json:"drawTime"`   //开奖日期
	GameName                string    `json:"gameName"`         //游戏名称
	PaidEndTime             string    `json:"paidEndTime"`       //截止兑奖日期
	PoolBalance             float64   `json:"poolBalance"`       //本期开奖后奖池余额，单位：元。
	TotalSaleAmount         float64   `json:"totalSaleAmount"`   //全国本期投注总额
	PrizeAmountSum          float64   `json:"prizeAmountSum"`   //应派奖金总计, 单位：元。
	PrizeStat              []PlsPrizeStat  `json:"prizeStat"`
}


type ZuCaiMatch struct {
	GuestTeam string `json:"guestTeam"`
	HostTeam  string `json:"hostTeam"`
	DrawResult string `json:"drawResult"`
}

type ZuCaiIssueResult struct {
	DrawNo         string        `json:"drawNo"`            //奖期编号
	GameName       string        `json:"gameName"`         //游戏名称
	SaleBeginTime  string        `json:"saleBeginTime"`    //开售时间
	SaleEndTime    string        `json:"saleEndTime"`      //止售时间
	DrawTime       string        `json:"drawTime"`        //开奖时间
	Against        []ZuCaiMatch  `json:"against"`
}

type ZuCaiResult struct {
	DrawNo                  string `json:"drawNo"` //奖期编号
	DrawResult              string `json:"drawResult"` //开奖结果
	DrawTime                string `json:"drawTime"`   //开奖日期
	GameName                string    `json:"gameName"`         //游戏名称
	PaidEndTime             string    `json:"paidEndTime"`       //截止兑奖日期
	PoolBalance             float64   `json:"poolBalance"`       //本期开奖后奖池余额，单位：元。
	TotalSaleAmount         float64   `json:"totalSaleAmount"`   //全国本期投注总额
	PrizeAmountSum          float64   `json:"prizeAmountSum"`   //应派奖金总计, 单位：元。
	PrizeStat      []PrizeStat  `json:"prizeStat"`   //奖等
	Against        []ZuCaiMatch  `json:"against"`    //一组对阵
}


type Matchs struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
}

type MatchsTeam struct {
	Id           int    `json:"id"`
	AwayTeam     Matchs `json:"awayTeam"`
	BusinessDate string `json:"businessDate"`
	HomeTeam     Matchs `json:"homeTeam"`
	League       Matchs `json:"league"`
	MatchDate    string `json:"matchDate"`
	MatchTime    string `json:"matchTime"`
	Num          int    `json:"num"`
	Status       string `json:"status"`
}

type JcMatchResult struct {
	Matchs []MatchsTeam `json:"matchs"`
}

type JcOddsResult struct {

}
