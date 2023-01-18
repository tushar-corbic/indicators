package indicators

type IndicatorType string

const (
	IndicatorTypeSMA           IndicatorType = "IndicatorTypeSMA"
	IndicatorTypeEMA           IndicatorType = "IndicatorTypeEMA"
	IndicatorTypeMACD          IndicatorType = "IndicatorTypeMACD"
	IndicatorTypeIchimokuCloud IndicatorType = "IndicatorTypeIchimokuCloud"
	IndicatorTypeATR           IndicatorType = "IndicatorTypeATR"
	IndicatorTypeSuperTrend    IndicatorType = "IndicatorTypeSuperTrend"
	IndicatorTypeHeikinAshi    IndicatorType = "IndicatorTypeHeikinAshi"
	IndicatorTypeStdDev        IndicatorType = "IndicatorTypeStdDev"
	IndicatorTypeHighest       IndicatorType = "IndicatorTypeHighest"
	IndicatorTypeLowest        IndicatorType = "IndicatorTypeLowest"
)

type IndicatorInputArg struct {
	Type IndicatorType

	Limit      int
	Period     int
	Multiplier float64

	// MacdLarge  int
	// MacdSmall  int
	// MacdSignal int

	// IchimokuCloudTenkan  int
	// IchimokuCloudKijun   int
	// IchimokuCloudSenkouB int
	// IchimokuCloudChikou  int
}
