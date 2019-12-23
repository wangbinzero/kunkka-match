package enum

type OrderAction string

const (
	//创建订单
	ActionCreate OrderAction = "create"

	//取消订单
	ActionCancel OrderAction = "cancel"
)

func (o OrderAction) String() string {
	switch o {
	case ActionCreate:
		return "create"
	case ActionCancel:
		return "cancel"
	default:
		return "unknown"
	}
}
