package ticker

import "time"

type Ticker struct {
	Symbol       string
	Volume       float64
	AveragePrice float64
	OpenPrice    float64
	ClosePrice   float64
	HighestPrice float64
	LowestPrice  float64
	Timestamp    time.Time
	Transactions int32
}
