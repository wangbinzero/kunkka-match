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
		this.parentList.PushFrontList(newList)

		//将子链表添加到map中
		this.elementMap[order.Price.String()] = newList.Front()
	} else {

		//读取map中该价格的链表
		val, ok := this.elementMap[order.Price.String()]
		if ok {
			//存在
			val.Value.(*list.List).PushBack(order)
		} else {
			newList := list.New()
			newList.PushFront(order)
			for e := this.parentList.Front(); e != nil; e = e.Next() {
				v := e.Value.(Order)
				//升序
				if this.sortBy == enum.SortDirectionAsc {
					//如果新订单价格小于订单价格
					if order.Price.LessThan(v.Price) {
						e.Value.(*list.List).PushFrontList(newList)
						this.parentList.PushBack(newList)
						break
					}
					continue
				} else {
					//降序
					//如果新订单价格小于订单价格
					if order.Price.GreaterThan(v.Price) {
						e.Value.(*list.List).PushFrontList(newList)
						this.parentList.PushBack(newList)
					}
				}
			}
		}
	}

	for e := this.parentList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
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
