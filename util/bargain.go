package util

import (
	"fmt"
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

func (this Bargain) Tosting() (string){
	return fmt.Sprintf("%v\t%s\t%v\t%.2f%%\t%v\t%v\t%v", this.Code, this.Stockname, this.Now, this.Tolow, this.Updatetime, this.Low, this.Day)

}

func (this *Bargain) Update(rt *SinaRealtime) {
	this.Now = rt.Now
	this.Stockname = rt.Stockname
	this.Tolow = (this.Now - this.Low) /this.Low * 100
	this.Updatetime =  rt.Date + " " + rt.Time
}
