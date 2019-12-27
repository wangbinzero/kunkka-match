package engine

import "kunkka-match/enum"

type OrderBook struct {
	buyOrderQueue  *orderQueue //买单队列
	sellOrderQueue *orderQueue //卖单队列
}

// 初始化订单簿
func (this *OrderBook) init() {

	//买单队列 降序
	buyOrderQueue := &orderQueue{}
	buyOrderQueue.init(enum.SortDirectionDesc)

	//卖单队列 升序
	sellOrderQueue := &orderQueue{}
	sellOrderQueue.init(enum.SortDirectionAsc)

	this.sellOrderQueue = sellOrderQueue
	this.buyOrderQueue = buyOrderQueue
}

//向订单簿添加买单
func (this *OrderBook) addBuyOrder(order Order) {
	this.buyOrderQueue.addOrder(order)
}

//向订单簿添加卖单
func (this *OrderBook) addSellOrder(order Order) {
	this.sellOrderQueue.addOrder(order)
}

//获取订单簿 买单队列头部订单
func (this *OrderBook) getHeadBuyOrder() Order {
	return this.buyOrderQueue.getHeadOrder()
}

//获取订单簿 卖单队列头部订单
func (this *OrderBook) getHeadSellOrder() Order {
	return this.sellOrderQueue.getHeadOrder()
}

//获取订单簿 买单队列头部订单 并移除
func (this *OrderBook) popHeadBuyOrder() Order {
	return this.buyOrderQueue.popHeadOrder()
}

//获取订单簿  卖单队列头部订单 并移除
func (this *OrderBook) popHeadSellOrder() Order {
	return this.sellOrderQueue.popHeadOrder()
}

//从订单簿 移除买单
func (this *OrderBook) removeBuyOrder(order *Order) {
	this.buyOrderQueue.removeOrder(*order)
}

//从订单簿 移除卖单
func (this *OrderBook) removeSellOrder(order *Order) {
	this.sellOrderQueue.removeOrder(*order)
}
