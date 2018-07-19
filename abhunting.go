package main

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/whomm/a-bargain-hunting/util"
)

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var blist sync.Map
var c chan os.Signal

func ndaylow(code string, day int) {

	var xp []util.IfengKdata
	for retry := 0; retry < 3; retry++ {
		var err error
		xp, err = util.Get_k_daily(code)
		if err == nil && len(xp) > 0 {
			break
		}
		if retry == 2 {
			fmt.Println("get history data error :" + code)
			return
		}

		time.Sleep(time.Second)
	}

	length := len(xp) - 1
	theb := util.Bargain{}
	theb.Low = 10000000
	theb.Code = code
	theb.Day = day

	var volumendaylist Int64Slice
	for i := 0; i < day; i++ {
		if length >= i {
			if xp[length-i].Low < theb.Low {
				theb.Low = xp[length-i].Low
			}
			volumendaylist = append(volumendaylist, xp[length-i].Volume)
		}
	}
	sort.Sort(volumendaylist)
	//计算中位数
	if day&1 == 1 {
		theb.VolumeMe = volumendaylist[day/2]
	} else {
		theb.VolumeMe = (volumendaylist[day/2] + volumendaylist[day/2-1]) / 2
	}
	//fmt.Println(volumendaylist, theb.VolumeMe)

	for {
		g, err := util.Get_real_time_data(code)
		if err == nil {
			theb.Update(g)
			blist.Store(code, theb)
		}
		time.Sleep(time.Second)
	}
}

//Tolow 最低价涨幅
type Tolow struct {
	code  string
	tolow float64
	tome  float64
}

//Tolowlist 最低价涨幅排序列表
type Tolowlist []Tolow

func (s Tolowlist) Len() int {
	return len(s)
}

func (s Tolowlist) Less(i, j int) bool {
	return s[i].tolow < s[j].tolow
}

func (s Tolowlist) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getbydes() {

	fmt.Print("\033[2J")
	for {

		tolowlist := Tolowlist{}
		blist.Range(func(key, i interface{}) bool {
			tolowlist = append(tolowlist, Tolow{code: i.(util.Bargain).Code, tolow: i.(util.Bargain).Tolow})
			return true
		})

		sort.Sort(tolowlist)

		lineno := 2
		fmt.Print("\033[1;0H\033[K股票代码\t 股票名称\t价格\t价变\t成交量    \t量变\t最低价\t 量中位数    \t周期/天\t更新时间")
		for _, j := range tolowlist {
			if now, ok := blist.Load(j.code); ok {
				fmt.Print("\033[" + strconv.Itoa(lineno) + ";0H\033[K" + now.(util.Bargain).Tosting())
			}

			lineno++
			if lineno >= 50 {
				break
			}
		}

		time.Sleep(time.Second)

	}
}

func main() {
	c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	for _, stockc := range util.Zz500 {
		go ndaylow(stockc[0], 30)
	}

	go getbydes()

LOOP:
	for {
		select {
		case s := <-c:
			fmt.Println()
			fmt.Println("interf", s)
			break LOOP
		default:
		}
		time.Sleep(500 * time.Millisecond)
	}
}
