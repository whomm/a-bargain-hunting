package util

//k线数据
type Kdata struct {
	//日期和时间
	//低频数据时为：YYYY-MM-DD
	//高频数为：YYYY-MM-DD HH:MM
	Date string
	//开盘价
	Open float64
	//收盘价
	Close float64
	//最高价
	High float64
	//最低价
	Low float64
	//成交量
	Volume int64
	//证券代码
	Code string
}

//凤凰网api接口日k线数据定义
type IfengKdata struct {
	//日期
	Date string
	//开盘价
	Start float64
	//最高价
	High float64
	//收盘价
	Close float64
	//最低价
	Low float64
	//成交量
	Volume int64
	//涨跌额
	ChangeAmount float64
	//涨跌幅
	QuoteChange float64

	//5日均价
	Ma5 float64
	//10日均价
	Ma10 float64
	//20日均价
	Ma20 float64
	//5日均量
	Vma5 float64
	//10日均量
	Vma10 float64
	//20日均量
	Vma20 float64

	//换手率
	ChangeRate float64
}

//新浪实时股票价格（实时数据接口 http://hq.sinajs.cn/list=sh601006）
type SinaRealtime struct {
	//股票名字
	Stockname string
	//今日开盘价
	Open float64
	//昨日收盘价
	Yesterdayclose float64
	//当前价格
	Now float64
	//今日最高价
	Todayhigh float64
	//今日最低价
	Todaylow float64
	//竞买价，即“买一”报价
	Buy1 float64
	//竞卖价，即“卖一”报价
	Sell1 float64
	//成交的股票数 由于股票交易以一百股为基本单位，所以在使用时，通常把该值除以一百
	Volume int64
	//成交金额，单位为“元”，为了一目了然，通常以“万元”为成交金额的单位，所以通常把该值除以一万
	Turnover float64

	//“买一” 申请4695股，即47手
	Buy1volume int64
	//“买一”报价；
	Buy1price  float64
	Buy2volume int64
	Buy2price  float64
	Buy3volume int64
	Buy3price  float64
	Buy4volume int64
	Buy4price  float64
	Buy5volume int64
	Buy5price  float64

	Sell1volume int64
	Sell1price  float64
	Sell2volume int64
	Sell2price  float64
	Sell3volume int64
	Sell3price  float64
	Sell4volume int64
	Sell4price  float64
	Sell5volume int64
	Sell5price  float64

	//日期
	Date string
	//时间
	Time string
	//未知
	Mscond string
}
