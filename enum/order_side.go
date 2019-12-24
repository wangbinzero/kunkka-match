package enum

type OrderSide string

const (
	SideBuy  OrderSide = "buy"
	SideSell OrderSide = "sell"
)

func (this OrderSide) String() string {
	switch this {
	case SideBuy:
		return "buy"
	case SideSell:
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
