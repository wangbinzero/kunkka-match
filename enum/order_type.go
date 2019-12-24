package enum

type OrderType string

const (
	Limit          OrderType = "limit"           //普通限价
	LimitIoc       OrderType = "limit-ioc"       //IOC限价-即时成交，剩余被撤
	Market         OrderType = "market"          //默认市价-及时成交剩余撤单
	MarketTop5     OrderType = "market-top5"     //市价-最优5档及时成交，剩余被撤
	MarketTop10    OrderType = "market-top10"    //市价-最优10档及时成交，剩余被撤
	MarketOpponent OrderType = "market-opponent" //市价-对手方最优价
)
