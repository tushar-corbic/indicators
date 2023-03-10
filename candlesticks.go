package indicators

import (
	"fmt"
	"sync"
)

type Interval string

const (
	IntervalOneMinute Interval = "1m"
	IntervalOneHour            = "1h"
	IntervalFourHours          = "4h"
	IntervalOneDay             = "1d"
)

type CandlestickMeta interface {
	Total() int
	Interval() Interval
	ItemAtIndex(idx int) *Candlestick
}
type CandlestickIndicators interface {
	AppendEMA(arg IndicatorInputArg) error
	AppendSMA(arg IndicatorInputArg) error
	AppendMACD(arg IndicatorInputArg) error
	AppendIchimokuCloud(arg IndicatorInputArg) error
}

type Candlesticks struct {
	CandlestickMeta
	CandlestickIndicators

	mux        sync.Mutex
	interval   Interval
	items      []*Candlestick
	maxCandles int
}

func (cs *Candlesticks) AppendCandlestick(c *Candlestick) error {
	cs.items = append(cs.items, c)
	if len(cs.items) > cs.maxCandles {
		cs.items = cs.items[1:]
	}
	return nil
}

func NewCandlesticks(i Interval, maxCandles int) (*Candlesticks, error) {
	return &Candlesticks{
		maxCandles: maxCandles,
		interval:   i,
		items:      make([]*Candlestick, 0),
	}, nil
}

// Total returns total candlesticks
func (cs *Candlesticks) Total() int {
	return len(cs.items)
}

// ItemAtIndex returns the item at specific index
func (cs *Candlesticks) ItemAtIndex(idx int) *Candlestick {
	if idx < 0 {
		return nil
	}
	if len(cs.items) > idx {
		return cs.items[idx]
	}
	return nil
}

// Interval returns currently set interval for the series of candlesticks
func (cs *Candlesticks) Interval() Interval {
	return cs.interval
}

// GenerateIndicator generates requested signals on that series of candlesticks
func (cs *Candlesticks) GenerateIndicator(i IndicatorType, arg IndicatorInputArg) error {
	switch i {
	case IndicatorTypeSMA:
		return cs.AppendSMA(arg)
	case IndicatorTypeEMA:
		return cs.AppendEMA(arg)
	// case IndicatorTypeMACD:
	// 	return cs.AppendMACD(arg)
	// case IndicatorTypeIchimokuCloud:
	// 	return cs.AppendIchimokuCloud(arg)
	case IndicatorTypeATR:
		return cs.AppendATR(arg)
	case IndicatorTypeSuperTrend:
		return cs.AppendSuperTrend(arg)
		// case IndicatorTypeHeikinAshi:
		// 	return cs.AppendHeikinAshi(arg)
		// case IndicatorTypeStdDev:
		// 	return cs.AppendStdDev(arg)
		// case IndicatorTypeHighest:
		// 	return cs.AppendHighest(arg)
		// case IndicatorTypeLowest:
		// 	return cs.AppendLowest(arg)
	}
	return fmt.Errorf("Error unsupported indicator type %+v", i)
}

// GetLastItem returns the candlestick that was most recently added
func (cs *Candlesticks) GetLastItem() *Candlestick {
	t := cs.Total()
	if t == 0 {
		return nil
	}
	return cs.items[t-1]
}
