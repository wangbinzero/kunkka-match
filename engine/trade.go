package engine

import (
	"github.com/shopspring/decimal"
	"time"
)

type Trade struct {
	MarkerId  string          `json:"markerId"`
	TakerId   string          `json:"takerId"`
	TakerSide string          `json:"takerSide"`
	Amount    decimal.Decimal `json:"amount"`
	Price     decimal.Decimal `json:"price"`
	Timestamp time.Time       `json:"timestamp"`
}

func (this Trade) toJson() {

}
