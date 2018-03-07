package util

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"regexp"
	"github.com/axgle/mahonia"
)







type IfengRespData struct {
	Record [][]string
}

var DefaultClient = &http.Client{}


// 获取历史日线数据
// http://api.finance.ifeng.com/akdaily/?code=sh600848&type=last

func Get_k_daily(code string) ([]IfengKdata,error){

	
	resp,err := DefaultClient.Get("http://api.finance.ifeng.com/akdaily/?code="+code+"&type=last")
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	respjson := IfengRespData{}
	err = json.Unmarshal(body,&respjson)
	if err != nil {
		return nil,err
	}
	var hisdata = []IfengKdata{}
	for _,strs :=range respjson.Record{
		tmp := IfengKdata{}
		tmp.Date =strs[0]
		tmp.Start, _ = strconv.ParseFloat(strs[1] ,64) 
		tmp.High , _ = strconv.ParseFloat(strs[2] ,64)
		tmp.Close , _ = strconv.ParseFloat(strs[3] ,64)
		tmp.Low , _ = strconv.ParseFloat(strs[4] ,64)
		tmp.Volume, _ = strconv.ParseInt(strs[5], 10, 64)
		tmp.ChangeAmount , _ = strconv.ParseFloat(strs[6] ,64)
		tmp.QuoteChange , _ = strconv.ParseFloat(strs[7] ,64)
		tmp.Ma5 , _ = strconv.ParseFloat(strs[8] ,64)
		tmp.Ma10 , _ = strconv.ParseFloat(strs[9] ,64)
		tmp.Ma20 , _ = strconv.ParseFloat(strs[10] ,64)
		tmp.Vma5 , _ = strconv.ParseFloat(strs[11] ,64)
		tmp.Vma10 , _ = strconv.ParseFloat(strs[12] ,64)
		tmp.Vma20 , _ = strconv.ParseFloat(strs[13] ,64)
		tmp.ChangeRate , _ = strconv.ParseFloat(strs[14] ,64) 

		hisdata = append(hisdata,tmp)
	}

	return hisdata,nil

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
func Get_real_time_data(code string) (*SinaRealtime,error){

	resp,err := DefaultClient.Get("http://hq.sinajs.cn/list="+code)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	bodystr := ConvertToString(string(body),"GBK","utf-8")

	re := regexp.MustCompile("\"([^\"]*)\"")
	m := re.FindStringSubmatch(bodystr)
	if len(m) >0{
		s := strings.Split(m[1], ",")
		sinart := SinaRealtime{}
		sinart.Stockname =s[0]
		sinart.Open  , _ = strconv.ParseFloat(s[1] ,64)
		sinart.Yesterdayclose  , _ = strconv.ParseFloat(s[2] ,64)
		sinart.Now  , _ = strconv.ParseFloat(s[3] ,64)
		sinart.Todayhigh  , _ = strconv.ParseFloat(s[4] ,64)
		sinart.Todaylow  , _ = strconv.ParseFloat(s[5] ,64)
		sinart.Buy1, _ = strconv.ParseFloat(s[6] ,64)
		sinart.Sell1  , _ = strconv.ParseFloat(s[7] ,64)
		sinart.Volume , _ = strconv.ParseInt(s[8], 10, 64)
		sinart.Turnover  , _ = strconv.ParseFloat(s[9] ,64)
		sinart.Buy1volume , _ = strconv.ParseInt(s[10], 10, 64)
		sinart.Buy1price  , _ = strconv.ParseFloat(s[11] ,64)
		sinart.Buy2volume , _ = strconv.ParseInt(s[12], 10, 64)
		sinart.Buy2price  , _ = strconv.ParseFloat(s[13] ,64)
		sinart.Buy3volume , _ = strconv.ParseInt(s[14], 10, 64)
		sinart.Buy3price  , _ = strconv.ParseFloat(s[15] ,64)
		sinart.Buy4volume , _ = strconv.ParseInt(s[16], 10, 64)
		sinart.Buy4price  , _ = strconv.ParseFloat(s[17] ,64)
		sinart.Buy5volume , _ = strconv.ParseInt(s[18], 10, 64)
		sinart.Buy5price  , _ = strconv.ParseFloat(s[19] ,64)
		sinart.Sell1volume , _ = strconv.ParseInt(s[20], 10, 64)
		sinart.Sell1price  , _ = strconv.ParseFloat(s[21] ,64)
		sinart.Sell2volume , _ = strconv.ParseInt(s[22], 10, 64)
		sinart.Sell2price  , _ = strconv.ParseFloat(s[23] ,64)
		sinart.Sell3volume , _ = strconv.ParseInt(s[24], 10, 64)
		sinart.Sell3price  , _ = strconv.ParseFloat(s[25] ,64)
		sinart.Sell4volume , _ = strconv.ParseInt(s[26], 10, 64)
		sinart.Sell4price  , _ = strconv.ParseFloat(s[27] ,64)
		sinart.Sell5volume , _ = strconv.ParseInt(s[28], 10, 64)
		sinart.Sell5price  , _ = strconv.ParseFloat(s[29] ,64)
		sinart.Date =s[30]
		sinart.Time =s[31]
		sinart.Mscond =s[32]
		return &sinart,nil	
	}
	if err != nil {
		return nil,err
	}
	return nil,nil
}




/*
code
证券代码：
支持沪深A、B股
支持全部指数
支持ETF基金


ktype

数据类型：
默认为D日线数据
D=日k线 W=周 M=月 
5=5分钟 15=15分钟 
30=30分钟 60=60分钟


autype

复权类型：
qfq-前复权 hfq-后复权 None-不复权，默认为qfq


index

是否为指数：
默认为False
设定为True时认为code为指数代码

start

开始日期
 format：YYYY-MM-DD 为空时取当前日期

end

结束日期 ：
format：YYYY-MM-DD
*/
func get_k_data(code string, ktype string, autype string, index bool, start string, end string){

}

