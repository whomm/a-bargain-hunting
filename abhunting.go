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



var blist = struct{  
    sync.RWMutex  
    m map[string]util.Bargain  
}{m: make(map[string]util.Bargain)} 

var c chan os.Signal


func ndaylow(code string, day int){

	var xp  []util.IfengKdata
	for retry:=0; retry<3; retry++ {
		var err error
		xp,err =util.Get_k_daily(code)
		if err == nil && len(xp) > 0 {
			break
		}
		if retry == 2 {
			fmt.Println("get history data error :"+ code)
			return
		}

		time.Sleep(time.Second)
	}

	


	
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
			if lineno>=50{
				break
			}
		}
		

		time.Sleep(time.Second)

	}
}




func main(){
	c = make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	

	for _,stockc := range util.Zz500 {
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