package engine

import (
	"container/list"
	"kunkka-match/enum"
)

//交易委托账本由两个队列组成，使用二维链接结合Map保存订单
//1.买单队列
//2.卖单队列

type orderQueue struct {
	sortBy     enum.SortDirection
	parentList *list.List
	elementMap map[string]*list.Element
}

// 初始化丁单薄
func (this *orderQueue) init(sortBy enum.SortDirection) {
	this.sortBy = sortBy
	this.parentList = list.New()
	this.elementMap = make(map[string]*list.Element)
}

// 添加订单
func (this *orderQueue) addOrder(order Order) {
	if this.parentList.Len() == 0 {
		this.parentList = list.New()

	}
}

// 读取头部订单
func (this *orderQueue) getHeadOrder() {

}

// 读取并删除头部订单
func (this *orderQueue) popHeadOrder() {

}

// 移除订单
func (this *orderQueue) removeOrder(order Order) {

}

// 读取深度价格
func (this *orderQueue) getDepthPrice(depth int) (string, int) {
	if this.parentList.Len() == 0 {
		return "", 0
	}
	p := this.parentList.Front()
	i := 1
	for ; i < depth; i++ {
		n := p.Next()
		if n != nil {
			p = n
		} else {
			break
		}
	}

	o := p.Value.(*list.List).Front().Value.(*Order)
	return o.Price.String(), i
}
