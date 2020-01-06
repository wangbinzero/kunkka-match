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
	//pay()

	//查询订单直接结果
	queryPayStatus("202001031406608799523", PAY_TRADE_URL)
}

const (
	PAY_URL          = "http://pay.vin/Pay_Index.html"       //支付地址
	PAY_TRADE_URL    = "http://pay.vin/Pay_Trade_query.html" //支付查询地址
	PAY_MEMBERID     = "10127"                               //商户号
	KEY              = "70k6fe7f1kxdhqj6ongprc4jjq1uznwl"    //授权key
	PAY_NOTIFY_URL   = "wwww.baidu.com"                      //通知地址
	PAY_CALLBACK_URL = "www.baidu.com"
)

var (
	AliPayH5AndScan = "922" //支付宝H5 + 扫码
	AliPayAndXianyu = "923" //支付宝咸鱼
)

//支付
func pay() {

	//金额范围  100/5000 每笔
	order := generateOrder("101", AliPayH5AndScan)
	sign := paySign(order)
	fmt.Println(sign)
	response, err := http.PostForm(PAY_URL, url.Values{
		"pay_amount":      {order["pay_amount"]},
		"pay_applydate":   {order["pay_applydate"]},
		"pay_bankcode":    {order["pay_bankcode"]},
		"pay_callbackurl": {order["pay_callbackurl"]},
		"pay_memberid":    {order["pay_memberid"]},
		"pay_notifyurl":   {order["pay_notifyurl"]},
		"pay_orderid":     {order["pay_orderid"]},
		"pay_md5sign":     {sign},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	responseBytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseBytes))

	//支付结果返回
	//<html>
	// <head>
	// <title></title>
	// <meta http-equiv="content-Type" content="text/html; charset=utf-8" />
	// </head>
	// <body>
	// <form action='https://pay.katpay.top/index/index/payDeposit' method='post' id='frm1'>
	// <input name="mer_num" type="hidden" value="1001009464" />
	// <input name="order_id" type="hidden" value="20200103140608489798" />
	// <input name="amount" type="hidden" value="101" />
	// <input name="type" type="hidden" value="H5QrCode" />
	// <input name="notify_url" type="hidden" value="http://pay.vin/Pay_KtxyPay_notifyurl.html" />
	// <input name="return_url" type="hidden" value="http://pay.vin/Pay_KtxyPay_callbackurl.html" />
	// <input name="sign" type="hidden" value="861f44909269edd20162a1ce9446669a" />
	// <input name="quota" type="hidden" value="1" />
	// </form>
	// <script language="javascript"> document.getElementById("frm1").submit();</script></frameset></body></html></frameset>
	//
}

//查询支付结果
func queryPayStatus(pay_orderid, addr string) {
	signTemp := "pay_memberid=" + PAY_MEMBERID + "&pay_orderid=" + pay_orderid + "&key=" + KEY
	sign := generateMd5sign(signTemp)
	response, err := http.PostForm(addr, url.Values{
		"pay_memberid": {PAY_MEMBERID},
		"pay_orderid":  {pay_orderid},
		"pay_md5sign":  {sign}})

	if err != nil {
		fmt.Println(err)
		return
	}

	responseBytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseBytes))

	//返回结果
	//{
	//	"memberid": "10127",
	//	"orderid": "202001031406608799523",
	//	"amount": "101.0000",
	//	"time_end": "1970-01-01 08:00:00",
	//	"transaction_id": "20200103140608489798",
	//	"returncode": "00",
	//	"trade_state": "NOTPAY",
	//	"sign": "473F430014A8CC18DA893137869C1010"
	//}
}

// md5方法签名
func generateMd5sign(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
}

// 构建订单
func generateOrder(pay_amount, pay_bankcode string) map[string]string {
	ret := make(map[string]string)
	ret["pay_amount"] = pay_amount
	ret["pay_applydate"] = time.Now().Format("2006-01-02 15:04:05")
	ret["pay_bankcode"] = pay_bankcode
	ret["pay_callbackurl"] = PAY_CALLBACK_URL
	ret["pay_memberid"] = PAY_MEMBERID
	ret["pay_notifyurl"] = PAY_NOTIFY_URL
	ret["pay_orderid"] = generatedOrderId()
	return ret
}

//type orderResponse struct {
//	pay_orderid     string `json:"pay_orderid"`     //订单
//	pay_applydate   string `json:"pay_applydate"`   //支付时间
//	pay_bankcode    string `json:"pay_bankcode"`    //支付方式
//	pay_notifyurl   string `json:"pay_notifyurl"`   //服务端通知
//	pay_callbackurl string `json:"pay_callbackurl"` //页面通知
//	pay_amount      string `json:"pay_amount"`      //支付金额
//}

//生成订单号
func generatedOrderId() string {
	date := time.Now().Format("200601021504405")
	random := strconv.Itoa(rand.Intn(899999) + 100000)
	return date + random
}

//支付数据签名
func paySign(order map[string]string) string {
	signTem := "pay_amount=" + order["pay_amount"] +
		"&pay_applydate=" + order["pay_applydate"] +
		"&pay_bankcode=" + order["pay_bankcode"] +
		"&pay_callbackurl=" + order["pay_callbackurl"] +
		"&pay_memberid=" + order["pay_memberid"] +
		"&pay_notifyurl=" + order["pay_notifyurl"] +
		"&pay_orderid=" + order["pay_orderid"] +
		"&key=" + KEY
	return generateMd5sign(signTem)
}
