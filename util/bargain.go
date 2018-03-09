package util

import (
	"fmt"
	"strings"
)

type Bargain struct {
	//股票名称
	Stockname string
	//股票代码
	Code string
	//n天周期
	Day int
	//n天最低价
	Low float64
	//当前价格
	Now float64
	//离最低价距离 (now - low )/low*100
	Tolow float64
	//当前价格更新时间
	Updatetime string
}

func (this Bargain) Tosting() string {
	return fmt.Sprintf("%v\t%5s\t%v\t%.2f%%\t%v\t%.2f\t%v", this.Code, this.Stockname, this.Now, this.Tolow, this.Updatetime, this.Low, this.Day)

}

func (this *Bargain) Update(rt *SinaRealtime) {
	this.Now = rt.Now
	this.Stockname = strings.Replace(rt.Stockname, " ", "", -1)
	this.Tolow = (this.Now - this.Low) / this.Low * 100
	this.Updatetime = rt.Date + " " + rt.Time

	if rt.Todayhigh < 0.0001 && rt.Open < 0.001 {
		//停盘的
		this.Tolow = 100
	}
}
