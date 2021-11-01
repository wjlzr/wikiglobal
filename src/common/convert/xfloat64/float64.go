package xfloat64

import "github.com/shopspring/decimal"

//float64乘法
func Mul(ff ...float64) float64 {
	df := decimal.NewFromFloat(ff[0])
	for i := 1; i < len(ff); i++ {
		df = df.Mul(decimal.NewFromFloat(ff[i]))
	}
	return decimalToFloat64(df)
}

//转成float64
func decimalToFloat64(d decimal.Decimal) float64 {
	f, _ := d.Float64()
	return f
}
