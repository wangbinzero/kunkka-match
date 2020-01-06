package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {

	//payUrl := "http://pay.vin/Pay_Index.html" //支付地址
	payTradeUrl := "http://pay.vin/Pay_Trade_query.html"
	pay_memberid := "10127" //商户号
	//secretKey := "70k6fe7f1kxdhqj6ongprc4jjq1uznwl" //秘钥
	//pay_orderid := generateOrderId()                //订单号
	//pay_applydate := generateTime()                 //支付时间
	//pay_bankcode := "923"                           //银行编码  922 支付宝H5+扫码  923支付宝咸鱼
	//pay_notifyurl := "http://www.baidu.com"         //后端回调地址
	//pay_callbackurl := "http://www.baidu.com"       //页面回调地址
	//pay_amount := 120                               //支付金额
	//
	//paySign := "pay_amount=" + strconv.Itoa(pay_amount) +
	//	"&pay_applydate=" + pay_applydate +
	//	"&pay_bankcode=" + pay_bankcode +
	//	"&pay_callbackurl=" + pay_callbackurl +
	//	"&pay_memberid=" + pay_memberid +
	//	"&pay_notifyurl=" + pay_notifyurl +
	//	"&pay_orderid=" + pay_orderid +
	//	"&key=" + secretKey
	//pay_md5sign := strings.ToUpper(generateSign(paySign))
	//fmt.Println("签名Sign:", pay_md5sign)
	//fmt.Println("订单号:",pay_orderid)
	//res, err := http.PostForm(payUrl, url.Values{"pay_amount": {strconv.Itoa(pay_amount)},
	//	"pay_applydate":   {pay_applydate},
	//	"pay_bankcode":    {pay_bankcode},
	//	"pay_callbackurl": {pay_callbackurl},
	//	"pay_memberid":    {pay_memberid},
	//	"pay_notifyurl":   {pay_notifyurl},
	//	"pay_orderid":     {pay_orderid},
	//	"pay_md5sign":     {pay_md5sign},
	//	"pay_productname": {"test"}})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//resBytes, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(resBytes))

	getPayStatus("20200103125926799523", pay_memberid, payTradeUrl)

}

//Deprecated
func generateOrderId() string {
	date := time.Now().Format("20060102150405")
	random := strconv.Itoa(rand.Intn(899999) + 100000)
	return date + random
}

//Deprecated
func generateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//Deprecated
func generateSign(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

//查询支付结果
func getPayStatus(pay_orderid, pay_memberid, pay_url string) {
	md5signTemp := "pay_memberid=" + pay_memberid + "&pay_orderid=" + pay_orderid + "&key=70k6fe7f1kxdhqj6ongprc4jjq1uznwl"
	md5sign := strings.ToUpper(generateSign(md5signTemp))
	res, err := http.PostForm(pay_url, url.Values{"pay_memberid": {pay_memberid}, "pay_orderid": {pay_orderid}, "pay_md5sign": {md5sign}})
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}
