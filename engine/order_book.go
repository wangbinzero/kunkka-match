package engine

import "kunkka-match/enum"

type OrderBook struct {
	buyOrderQueue  *orderQueue //买单队列
	sellOrderQueue *orderQueue //卖单队列
}

// 初始化订单簿
func (this *OrderBook) init() {
	buyOrderQueue := &orderQueue{}
	buyOrderQueue.init(enum.SortDirectionAsc)

	sellOrderQueue := &orderQueue{}
	sellOrderQueue.init(enum.SortDirectionDesc)

	this.sellOrderQueue = sellOrderQueue
	this.buyOrderQueue = buyOrderQueue
}

//向订单簿添加买单
func (this *OrderBook) addBuyOrder(order Order) {

}

//向订单簿添加卖单
func (this *OrderBook) addSellOrder(order Order) {

}

//获取订单簿 买单队列头部订单
func (this *OrderBook) getHeadBuyOrder() {

}

//获取订单簿 卖单队列头部订单
func (this *OrderBook) getHeadSellOrder() {

}

//获取订单簿 买单队列头部订单 并移除
func (this *OrderBook) popHeadBuyOrder() {

}

//获取订单簿  卖单队列头部订单 并移除
func (this *OrderBook) popHeadSellOrder() {

}

//从订单簿 移除买单
func (this *OrderBook) removeBuyOrder(order *Order) bool {
	return false
}

//从订单簿 移除卖单
func (this *OrderBook) removeSellOrder(order *Order) bool {
	return false
}
