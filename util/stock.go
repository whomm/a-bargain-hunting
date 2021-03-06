package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/axgle/mahonia"
)

type IfengRespData struct {
	Record [][]string
}

var DefaultClient = &http.Client{}

// 获取历史日线数据(不包含当日)
// http://api.finance.ifeng.com/akdaily/?code=sh600036&type=last

func Get_k_daily(code string) ([]IfengKdata, error) {

	resp, err := DefaultClient.Get("http://api.finance.ifeng.com/akdaily/?code=" + code + "&type=last")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	respjson := IfengRespData{}
	err = json.Unmarshal(body, &respjson)
	if err != nil {
		return nil, err
	}
	var hisdata = []IfengKdata{}
	for _, strs := range respjson.Record {
		tmp := IfengKdata{}
		tmp.Date = strs[0]
		tmp.Start, _ = strconv.ParseFloat(strs[1], 64)
		tmp.High, _ = strconv.ParseFloat(strs[2], 64)
		tmp.Close, _ = strconv.ParseFloat(strs[3], 64)
		tmp.Low, _ = strconv.ParseFloat(strs[4], 64)
		tvlum, _ := strconv.ParseFloat(strs[5], 64)
		tmp.Volume = int64(tvlum * 100)
		tmp.ChangeAmount, _ = strconv.ParseFloat(strs[6], 64)
		tmp.QuoteChange, _ = strconv.ParseFloat(strs[7], 64)
		tmp.Ma5, _ = strconv.ParseFloat(strs[8], 64)
		tmp.Ma10, _ = strconv.ParseFloat(strs[9], 64)
		tmp.Ma20, _ = strconv.ParseFloat(strs[10], 64)
		tmp.Vma5, _ = strconv.ParseFloat(strs[11], 64)
		tmp.Vma10, _ = strconv.ParseFloat(strs[12], 64)
		tmp.Vma20, _ = strconv.ParseFloat(strs[13], 64)
		tmp.ChangeRate = 0 // strconv.ParseFloat(strs[14], 64)

		hisdata = append(hisdata, tmp)
	}

	return hisdata, nil

}

//字符串类型转换
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//获取实时交易数据
func Get_real_time_data(code string) (*SinaRealtime, error) {

	resp, err := DefaultClient.Get("http://hq.sinajs.cn/list=" + code)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	bodystr := ConvertToString(string(body), "GBK", "utf-8")

	re := regexp.MustCompile("\"([^\"]*)\"")
	m := re.FindStringSubmatch(bodystr)
	if len(m) > 0 {
		s := strings.Split(m[1], ",")
		sinart := SinaRealtime{}
		sinart.Stockname = s[0]
		sinart.Open, _ = strconv.ParseFloat(s[1], 64)
		sinart.Yesterdayclose, _ = strconv.ParseFloat(s[2], 64)
		sinart.Now, _ = strconv.ParseFloat(s[3], 64)
		sinart.Todayhigh, _ = strconv.ParseFloat(s[4], 64)
		sinart.Todaylow, _ = strconv.ParseFloat(s[5], 64)
		sinart.Buy1, _ = strconv.ParseFloat(s[6], 64)
		sinart.Sell1, _ = strconv.ParseFloat(s[7], 64)
		sinart.Volume, _ = strconv.ParseInt(s[8], 10, 64)
		sinart.Turnover, _ = strconv.ParseFloat(s[9], 64)
		sinart.Buy1volume, _ = strconv.ParseInt(s[10], 10, 64)
		sinart.Buy1price, _ = strconv.ParseFloat(s[11], 64)
		sinart.Buy2volume, _ = strconv.ParseInt(s[12], 10, 64)
		sinart.Buy2price, _ = strconv.ParseFloat(s[13], 64)
		sinart.Buy3volume, _ = strconv.ParseInt(s[14], 10, 64)
		sinart.Buy3price, _ = strconv.ParseFloat(s[15], 64)
		sinart.Buy4volume, _ = strconv.ParseInt(s[16], 10, 64)
		sinart.Buy4price, _ = strconv.ParseFloat(s[17], 64)
		sinart.Buy5volume, _ = strconv.ParseInt(s[18], 10, 64)
		sinart.Buy5price, _ = strconv.ParseFloat(s[19], 64)
		sinart.Sell1volume, _ = strconv.ParseInt(s[20], 10, 64)
		sinart.Sell1price, _ = strconv.ParseFloat(s[21], 64)
		sinart.Sell2volume, _ = strconv.ParseInt(s[22], 10, 64)
		sinart.Sell2price, _ = strconv.ParseFloat(s[23], 64)
		sinart.Sell3volume, _ = strconv.ParseInt(s[24], 10, 64)
		sinart.Sell3price, _ = strconv.ParseFloat(s[25], 64)
		sinart.Sell4volume, _ = strconv.ParseInt(s[26], 10, 64)
		sinart.Sell4price, _ = strconv.ParseFloat(s[27], 64)
		sinart.Sell5volume, _ = strconv.ParseInt(s[28], 10, 64)
		sinart.Sell5price, _ = strconv.ParseFloat(s[29], 64)
		sinart.Date = s[30]
		sinart.Time = s[31]
		sinart.Mscond = s[32]
		return &sinart, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, nil
}
