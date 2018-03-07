package main




import (
	"strconv"
	"sort"
	"time"
	"fmt"
	"github.com/whomm/a-bargain-hunting/util"
	"sync"
	"os"
	"os/signal"
)

//监控列表
var monitoringtargets = [...][]string{
	{"浦发银行", "sh600000"},
	{"民生银行", "sh600016"},
	{"中国石化", "sh600028"},
	{"南方航空", "sh600029"},
	{"中信证券", "sh600030"},
	{"招商银行", "sh600036"},
	{"保利地产", "sh600048"},
	{"中国联通", "sh600050"},
	{"同方股份", "sh600100"},
	{"上汽集团", "sh600104"},
	{"北方稀土", "sh600111"},
	{"华夏幸福", "sh600340"},
	{"信威集团", "sh600485"},
	{"康美药业", "sh600518"},
	{"贵州茅台", "sh600519"},
	{"山东黄金", "sh600547"},
	{"绿地控股", "sh600606"},
	{"海通证券", "sh600837"},
	{"伊利股份", "sh600887"},
	{"江苏银行", "sh600919"},
	{"东方证券", "sh600958"},
	{"招商证券", "sh600999"},
	{"大秦铁路", "sh601006"},
	{"中国神华", "sh601088"},
	{"兴业银行", "sh601166"},
	{"北京银行", "sh601169"},
	{"中国铁建", "sh601186"},
	{"东兴证券", "sh601198"},
	{"国泰君安", "sh601211"},
	{"上海银行", "sh601229"},
	{"农业银行", "sh601288"},
	{"中国平安", "sh601318"},
	{"交通银行", "sh601328"},
	{"新华保险", "sh601336"},
	{"中国中铁", "sh601390"},
	{"工商银行", "sh601398"},
	{"中国太保", "sh601601"},
	{"中国人寿", "sh601628"},
	{"中国建筑", "sh601668"},
	{"华泰证券", "sh601688"},
	{"中国中车", "sh601766"},
	{"光大证券", "sh601788"},
	{"中国交建", "sh601800"},
	{"光大银行", "sh601818"},
	{"中国石油", "sh601857"},
	{"中国银河", "sh601881"},
	{"方正证券", "sh601901"},
	{"中国核电", "sh601985"},
	{"中国银行", "sh601988"},
	{"中国重工", "sh601989"},
}



var blist = struct{  
    sync.RWMutex  
    m map[string]util.Bargain  
}{m: make(map[string]util.Bargain)} 

var c chan os.Signal


func ndaylow(code string, day int){
	xp,_:=util.Get_k_daily(code)
	
	length  := len(xp) -1
	theb := util.Bargain{}
	theb.Low = 10000000
	theb.Code=code
	theb.Day = day

	for i :=0; i<day; i++  {
		if length>=i && xp[length-i].Low < theb.Low  {
			theb.Low = xp[length-i].Low
		}
	}

	for {
		g, err := util.Get_real_time_data(code)
		if err == nil{

			theb.Update(g)

			blist.Lock()
			blist.m[code] = theb
			blist.Unlock()
		}
		time.Sleep(time.Second)
	}
}


type Tolow struct {
	code  string
	tolow float64
}

type Tolowlist []Tolow

func (s Tolowlist) Len() (int) {
	return len(s)
}

func (s Tolowlist) Less(i, j int) (bool) {
	return s[i].tolow < s[j].tolow
}

func (s Tolowlist) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getbydes(){

	for {

		tolowlist := Tolowlist{}
		blist.RLock()
		for _,i:= range blist.m{
			tolowlist = append(tolowlist, Tolow{code:i.Code, tolow: i.Tolow})
		}
		blist.RUnlock()
		sort.Sort(tolowlist)

		lineno := 0
		fmt.Printf("\033[2J")
		for _,j := range tolowlist {
			blist.RLock()
			fmt.Printf("\033["+strconv.Itoa(lineno)+";0H")
			fmt.Print(blist.m[j.code].Tosting()+"\r")
			blist.RUnlock()

			lineno += 1
		}
		

		time.Sleep(time.Second)

	}
}




func main(){
	c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	

	for _,stockc := range monitoringtargets {
		go ndaylow(stockc[1],30)
	}
	
	go getbydes()

	LOOP:
	for{
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