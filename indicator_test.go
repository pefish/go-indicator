package go_indicator

import (
	"fmt"
	go_test_ "github.com/pefish/go-test"
	"math"
	"testing"
)

func TestIndicator_Sma(t *testing.T) {
	result := NewIndicator().SMA([]float64{
		1,
		4,
		3,
		4,
		5,
		6,
	}, 2)
	go_test_.Equal(t, true, math.IsNaN(result[0]))
	go_test_.Equal(t, 2.5, result[1])
	go_test_.Equal(t, 3.5, result[3])
	go_test_.Equal(t, 5.5, result[5])

	result1 := NewIndicator().SMA([]float64{
		1,
		4,
		3,
		4,
		5,
		6,
	}, 3)
	go_test_.Equal(t, true, math.IsNaN(result1[0]))
	go_test_.Equal(t, true, math.IsNaN(result1[1]))
	go_test_.Equal(t, 2.6666666666666665, result1[2])
	go_test_.Equal(t, 3.6666666666666665, result1[3])
	go_test_.Equal(t, 4.0, result1[4])
	go_test_.Equal(t, 5.0, result1[5])
}

func TestIndicator_Ema(t *testing.T) {
	result := NewIndicator().EMA([]float64{
		34220.4,
		34349.2,
		34352.5,
		34195.4,
		34142.6,
		34068.9,
	}, 3)
	fmt.Println(result)
	go_test_.Equal(t, true, math.IsNaN(result[0]))
	go_test_.Equal(t, true, math.IsNaN(result[1]))
	go_test_.Equal(t, 34307.36666666667, result[2])
	go_test_.Equal(t, 34251.38333333333, result[3])
	go_test_.Equal(t, 34196.99166666667, result[4])
	go_test_.Equal(t, 34132.94583333333, result[5])
}

func TestIndicator_MACD(t *testing.T) {
	diffEma, signalEma, histogram := NewIndicator().MACD([]float64{
		34220.4,
		34349.2,
		34352.5,
		34195.4,
		34142.6,
		34068.9,
	}, 2, 4, 3)
	fmt.Println(diffEma, signalEma, histogram)
	go_test_.Equal(t, true, math.IsNaN(diffEma[0]))
	go_test_.Equal(t, true, math.IsNaN(diffEma[2]))
	go_test_.Equal(t, -49.51685185185488, diffEma[4])
	go_test_.Equal(t, true, math.IsNaN(signalEma[0]))
	go_test_.Equal(t, true, math.IsNaN(signalEma[4]))
	go_test_.Equal(t, -48.89678600823148, signalEma[5])
	go_test_.Equal(t, true, math.IsNaN(histogram[0]))
	go_test_.Equal(t, true, math.IsNaN(histogram[4]))
	go_test_.Equal(t, -9.146164609056235, histogram[5])
}

func TestIndicator_BOLL(t *testing.T) {
	u, m, l := NewIndicator().BOLL([]float64{
		34220.4,
		34349.2,
		34352.5,
		34195.4,
		34142.6,
		34068.9,
	}, 3, 2)
	go_test_.Equal(t, true, math.IsNaN(u[0]))
	go_test_.Equal(t, true, math.IsNaN(u[1]))
	go_test_.Equal(t, 34430.38561732965, u[2])
	go_test_.Equal(t, 34445.61776482438, u[3])
	go_test_.Equal(t, 34408.46262536705, u[4])
	go_test_.Equal(t, 34239.388986459104, u[5])
	go_test_.Equal(t, true, math.IsNaN(m[0]))
	go_test_.Equal(t, true, math.IsNaN(m[1]))
	go_test_.Equal(t, 34307.36666666667, m[2])
	go_test_.Equal(t, 34299.03333333333, m[3])
	go_test_.Equal(t, 34230.166666666664, m[4])
	go_test_.Equal(t, 34135.63333333333, m[5])
	go_test_.Equal(t, true, math.IsNaN(l[0]))
	go_test_.Equal(t, true, math.IsNaN(l[1]))
	go_test_.Equal(t, 34184.347716003685, l[2])
	go_test_.Equal(t, 34152.44890184228, l[3])
	go_test_.Equal(t, 34051.87070796628, l[4])
	go_test_.Equal(t, 34031.87768020756, l[5])
}

func TestIndicator_KDJ(t *testing.T) {
	k, d, j := NewIndicator().KDJ([]Bar{
		{
			High:  34220.4,
			Open:  34210.4,
			Low:   34220.2,
			Close: 34250.4,
		},
		{
			High:  34420.4,
			Open:  34250.4,
			Low:   34480.2,
			Close: 34320.4,
		},
		{
			High:  34520.4,
			Open:  34320.4,
			Low:   34520.2,
			Close: 34550.4,
		},
		{
			High:  34620.4,
			Open:  34550.4,
			Low:   34620.2,
			Close: 34650.4,
		},
		{
			High:  34320.4,
			Open:  34650.4,
			Low:   34230.2,
			Close: 34150.4,
		},
		{
			High:  34220.4,
			Open:  34150.4,
			Low:   34220.2,
			Close: 34250.4,
		},
	}, 3, 2, 2)
	go_test_.Equal(t, true, math.IsNaN(k[0]))
	go_test_.Equal(t, true, math.IsNaN(k[1]))
	go_test_.Equal(t, 79.99666888740832, k[2])
	go_test_.Equal(t, 100.69733587023735, k[3])
	go_test_.Equal(t, 40.123142563515046, k[4])
	go_test_.Equal(t, 23.83468472503639, k[5])
	go_test_.Equal(t, true, math.IsNaN(d[0]))
	go_test_.Equal(t, true, math.IsNaN(d[1]))
	go_test_.Equal(t, 64.99833444370415, d[2])
	go_test_.Equal(t, 82.84783515697075, d[3])
	go_test_.Equal(t, 61.485488860242896, d[4])
	go_test_.Equal(t, 42.660086792639646, d[5])
	go_test_.Equal(t, true, math.IsNaN(j[0]))
	go_test_.Equal(t, true, math.IsNaN(j[1]))
	go_test_.Equal(t, 109.99333777481665, j[2])
	go_test_.Equal(t, 136.39633729677058, j[3])
	go_test_.Equal(t, -2.601550029940654, j[4])
	go_test_.Equal(t, -13.816119410170124, j[5])
}

func TestIndicator_RSI(t *testing.T) {
	result := NewIndicator().RSI([]float64{
		34220.4,
		34349.2,
		34352.5,
		34195.4,
		34142.6,
		34068.9,
	}, 3)
	fmt.Println(result)
	go_test_.Equal(t, true, math.IsNaN(result[0]))
	go_test_.Equal(t, true, math.IsNaN(result[1]))
	go_test_.Equal(t, true, math.IsNaN(result[2]))
	go_test_.Equal(t, 45.67773167358225, result[3])
	go_test_.Equal(t, 35.8577633007595, result[4])
	go_test_.Equal(t, 24.727408863306607, result[5])
}
