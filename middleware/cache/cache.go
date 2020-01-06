package cache

import (
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
	"kunkka-match/common"
	"kunkka-match/log"
	"kunkka-match/middleware"
)

func GetSymbols() []string {
	symbols, _ := middleware.RedisClient.SMembers(common.SYMBOL_KEY).Result()
	return symbols
}

// 保存交易标的缓存
func SaveSymbol(symbol string) {
	// lpush 压栈操作 将元素放入头部，允许重复
	// lpushx 压栈操作 将元素放入头部 不允许重复
	//
	middleware.RedisClient.SAdd(common.SYMBOL_KEY, symbol)
	//middleware.RedisClient.LPush("kunkka:match:symbol", symbol)
}

//移除交易对
func RemoveSymbol(symbol string) {

	//lrem 移除等于 symbol的元素
	//当 count>0 时，从表头开始查找移除count个
	//当 count=0 时，从表头开始查找，移除所有等于value的
	middleware.RedisClient.SRem(common.SYMBOL_KEY, symbol)
}

// 删除所有交易标的： 删除key
func RemoveAllSymbol() {
	middleware.RedisClient.Del(common.SYMBOL_KEY)
}

//保存交易标的以及价格
func SavePrice(symbol string, price string) {
	middleware.RedisClient.Set(common.PRICE_KEY+symbol, price, 0)
}

// 根据交易标的查询价格
func GetPrice(symbol string) decimal.Decimal {
	str := middleware.RedisClient.Get(common.PRICE_KEY + symbol).Val()
	val, _ := decimal.NewFromString(str)
	return val
}

func RemovePrice(symbol string) {
	middleware.RedisClient.Del(common.PRICE_KEY + symbol)
}

// 将订单写入缓存
func SaveOrder(order map[string]interface{}) {
	action := order["action"].(string)
	symbol := order["symbol"].(string)
	orderId := order["orderId"].(string)
	timestamp := order["timestamp"].(float64)
	//订单ID + 订单标志[ 下单/撤单 ]
	middleware.RedisClient.HMSet(common.ORDER_KEY+symbol+":"+orderId+":"+action, order)

	z := redis.Z{
		Score:  timestamp,
		Member: orderId + ":" + action,
	}
	middleware.RedisClient.ZAdd(common.ORDER_IDS_KEY+symbol, z)
}

//查询订单ID集合
func GetOrderIdsWithAction(symbol string) []string {
	return middleware.RedisClient.ZRange(common.ORDER_IDS_KEY+symbol, 0, -1).Val()
}

func GetOrder(symbol, orderId string) map[string]string {
	res := middleware.RedisClient.HGetAll(common.ORDER_KEY + symbol + ":" + orderId).Val()
	log.Info("缓存加载订单: 交易标的 [%s] 订单 %v\n", symbol, res)
	return res
}

func UpdateOrder() {

}

// 删除缓存中订单信息
func RemoveOrder(symbol, orderId, action string) {
	middleware.RedisClient.Del(common.ORDER_KEY + symbol + ":" + orderId + ":" + action)
}

// 判断缓存中是否存在订单
func OrderExist(symbol, orderId, action string) bool {
	return middleware.RedisClient.HExists(common.ORDER_KEY+symbol+":"+orderId+":"+action, "symbol").Val()
}
