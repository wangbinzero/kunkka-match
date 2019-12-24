package process

import "kunkka-match/middleware/cache"

func Init() {
	symbols := cache.GetSymbols()
	for _, symbol := range symbols {
		price := cache.GetPrice(symbol)
		NewEngine(symbol, price)

		//rderIds:=
	}
}
