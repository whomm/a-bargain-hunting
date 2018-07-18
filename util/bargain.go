package util

import (
	"fmt"
	"strconv"
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
	//n天的成交量中位数
	VolumeMe int64
	//成交量离中位数距离 (volume - volumemedian )/volumemedian*100
	TomMe float64
	//当前成交量
	Volume int64
}

func (this Bargain) Tosting() string {
	return fmt.Sprintf("%v\t%5s\t%v\t%.2f%%\t%v\t%.2f\t%v", this.Code, this.Stockname, this.Now, this.Tolow, this.Updatetime, this.Low, this.Day)

}

func (this *Bargain) Update(rt *SinaRealtime) {
	this.Now = rt.Now
	this.Stockname = strings.Replace(rt.Stockname, " ", "", -1)
	this.Tolow = (this.Now - this.Low) / this.Low * 100
	this.Updatetime = rt.Date + " " + rt.Time
	this.Volume = rt.Volume

	//根据当前交易量预估今天交易量

	timesplit := strings.Split(rt.Time, ":")
	hour, _ := strconv.Atoi(timesplit[0])
	min, _ := strconv.Atoi(timesplit[1])

	nowmin := hour*60 + min
	passmin := 1
	if nowmin > 9*60+30 && nowmin <= 11*60+30 {
		//上午的交易时间段
		passmin = nowmin - 9*60 + 30
	}
	if nowmin >= 13*60 && nowmin <= 15*60 {
		//下午交易时间段
		passmin = nowmin - 13*60 + 2*60
	}

	this.TomMe = (float64(rt.Volume)*float64(5*60)/float64(passmin) - float64(this.VolumeMe)) / float64(this.VolumeMe) * 100

	if rt.Todayhigh < 0.0001 && rt.Open < 0.001 {
		//停盘的
		this.Tolow = 100
		this.TomMe = 100
	}
}
