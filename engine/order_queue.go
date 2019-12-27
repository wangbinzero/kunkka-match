package engine

import (
	"container/list"
	"fmt"
	"kunkka-match/enum"
)

//交易委托账本由两个队列组成，使用二维链接结合Map保存订单
//1.买单队列
//2.卖单队列

type orderQueue struct {
	sortBy     enum.SortDirection       //排序方向 asc/desc
	parentList *list.List               //保存整个二维链表的所有订单  一维 以价格排序  二维以时间排序
	elementMap map[string]*list.Element //key为价格  value为第二维订单表的键值对
}

// 初始化丁单薄
func (this *orderQueue) init(sortBy enum.SortDirection) {
	this.sortBy = sortBy
	this.parentList = list.New()
	this.elementMap = make(map[string]*list.Element)
}

// 添加订单
func (this *orderQueue) addOrder(order Order) {

	//如果父链表长度为0
	if this.parentList.Len() == 0 {

		//新建子链表 并且将订单插入子链表
		newList := list.New()
		newList.PushFront(order)
		//将子链表插入父链表
		el := this.parentList.PushFront(newList)
		//将子链表添加到map中
		this.elementMap[order.Price.String()] = el
	} else {
		//读取map中该价格的链表
		el, ok := this.elementMap[order.Price.String()]
		if ok {
			//当前价格的数据已经存在，则需要根据时间序列，将此条数据放在查询出来的数据的后面
			el.Value.(*list.List).PushBack(order)
		} else {
			//创建新子链
			newList := list.New()
			newList.PushFront(order)

			//先取头和尾
			parentHeadList := this.parentList.Front()
			parentTailList := this.parentList.Back()

			frontOrder := parentHeadList.Value.(*list.List).Front().Value.(Order)
			backOrder := parentTailList.Value.(*list.List).Front().Value.(Order)

			//升序 卖单
			if this.sortBy == enum.SortDirectionAsc {

				if order.Price.LessThan(frontOrder.Price) {
					el := this.parentList.PushFront(newList)
					this.elementMap[order.Price.String()] = el
					return
				}

				if order.Price.GreaterThan(backOrder.Price) {
					el := this.parentList.PushBack(newList)
					this.elementMap[order.Price.String()] = el
					return
				}

				for el := this.parentList.Front(); el != nil; el = el.Next() {
					childList := el.Value.(*list.List)
					childPrice := childList.Front().Value.(Order).Price
					if order.Price.LessThan(childPrice) {
						el := this.parentList.InsertBefore(order, el)
						this.elementMap[order.Price.String()] = el
						break
					}
				}

			} else {
				//降序 买单
				if order.Price.GreaterThan(frontOrder.Price) {
					el := this.parentList.PushFront(newList)
					this.elementMap[order.Price.String()] = el
					return
				}

				if order.Price.LessThan(backOrder.Price) {
					el := this.parentList.PushBack(newList)
					this.elementMap[order.Price.String()] = el
					return
				}

				for el := this.parentList.Front(); el != nil; el = el.Next() {
					childList := el.Value.(*list.List)
					childPrice := childList.Front().Value.(Order).Price
					if order.Price.GreaterThan(childPrice) {
						el := this.parentList.InsertBefore(order, el)
						this.elementMap[order.Price.String()] = el
						break
					}
				}
			}
		}
	}

	for e := this.parentList.Front(); e != nil; e = e.Next() {
		for el := e.Value.(*list.List).Front(); el != nil; el = el.Next() {
			fmt.Println("打印:", el.Value)
		}

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
// 为了方便处理根据档位优先撮合的问题  market-top5 market-top10
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
