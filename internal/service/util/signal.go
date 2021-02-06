package util

type Signal struct {
	Market          string
	Side            string
	Reason          string
	Open            float64
	TakeProfit      float64
	StopLoss        float64
	UseTrailingStop bool
	Ratio           float64 // what ratio should a trader use its balance for this trade
}
