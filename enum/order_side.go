package enum

type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

func (this OrderSide) String() string {
	switch this {
	case Buy:
		return "buy"
	case Sell:
		return "sell"
	default:
		return "unknown"
	}
}

func (this OrderSide) Valid() bool {
	if this.String() == "unknown" {
		return false
	}
	return true
}
