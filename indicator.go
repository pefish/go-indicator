package go_indicator

import (
	"math"
)

type Indicator struct {
}

type Bar struct {
	High  float64 `json:"high"`
	Open  float64 `json:"open"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`
}

func NewIndicator() *Indicator {
	return &Indicator{}
}

func (ind *Indicator) skip(records []float64, length int) int {
	result := 0
	for k := 0; result < len(records); result++ {
		if !math.IsNaN(records[result]) {
			k++
		}
		if k == length {
			break
		}
	}
	return result
}

func (ind *Indicator) set(records []float64, start int, end int, value float64) {
	e := len(records)
	if end < e {
		e = end
	}
	for i := start; i < e; i++ {
		records[i] = value
	}
}

func (ind *Indicator) sum(records []float64, num int) float64 {
	result := 0.0
	for i := 0; i < num; i++ {
		if !math.IsNaN(records[i]) {
			result += records[i]
		}
	}
	return result
}

func (ind *Indicator) avg(records []float64, num int) float64 {
	n := 0.0
	sum := 0.0
	for i := 0; i < num; i++ {
		if !math.IsNaN(records[i]) {
			sum += records[i]
			n++
		}
	}
	return sum / n
}

func (ind *Indicator) moveDiff(records []float64) []float64 {
	results := make([]float64, 0)
	for i := 1; i < len(records); i++ {
		results = append(results, records[i]-records[i-1])
	}
	return results
}

func (ind *Indicator) sma(records []float64, length int) []float64 {
	r := make([]float64, len(records))
	j := ind.skip(records, length)
	ind.set(r, 0, j, math.NaN())
	if j < len(records) {
		sum := 0.0
		for i := j; i < len(records); i++ {
			if i == j {
				sum = ind.sum(records, i+1)
			} else {
				sum += records[i] - records[i-length]
			}
			r[i] = sum / float64(length)
		}
	}
	return r
}

func (ind *Indicator) cmp(
	records []float64,
	start int,
	end int,
	comFunc func(a, b float64) float64,
) float64 {
	v := records[start]
	for i := start; i < end; i++ {
		v = comFunc(records[i], v)
	}
	return v
}

func (ind *Indicator) SMA(records []float64, length int) []float64 {
	if length == 0 {
		length = 9
	}
	return ind.sma(records, length)
}

func (ind *Indicator) ema(records []float64, length int) []float64 {
	r := make([]float64, len(records))
	multiplier := 2.0 / float64(length+1)
	j := ind.skip(records, length)
	ind.set(r, 0, j, math.NaN())
	if j < len(records) {
		r[j] = ind.avg(records, j+1)
		for i := j + 1; i < len(records); i++ {
			r[i] = (records[i]-r[i-1])*multiplier + r[i-1]
		}
	}
	return r
}

func (ind *Indicator) diff(records1, records2 []float64) []float64 {
	results := make([]float64, 0)
	for i := 0; i < len(records2); i++ {
		if math.IsNaN(records1[i]) || math.IsNaN(records2[i]) {
			results = append(results, math.NaN())
		} else {
			results = append(results, records1[i]-records2[i])
		}
	}
	return results
}

func (ind *Indicator) EMA(records []float64, length int) []float64 {
	if length == 0 {
		length = 9
	}
	return ind.ema(records, length)
}

func (ind *Indicator) MACD(
	records []float64,
	fastLength int,
	slowLength int,
	signalLength int,
) (diffEma []float64, signalEma []float64, histogram []float64) {
	if fastLength == 0 {
		fastLength = 12
	}

	if slowLength == 0 {
		slowLength = 26
	}

	if signalLength == 0 {
		signalLength = 9
	}

	slowEma := ind.ema(records, slowLength)
	fastEma := ind.ema(records, fastLength)
	diffEma = ind.diff(fastEma, slowEma)

	signalEma = ind.ema(diffEma, signalLength)
	histogram = ind.diff(diffEma, signalEma)
	return diffEma, signalEma, histogram
}

func (ind *Indicator) BOLL(
	records []float64,
	length int,
	multiplier float64,
) (upper []float64, middle []float64, lower []float64) {
	if length == 0 {
		length = 20
	}
	if multiplier == 0 {
		multiplier = 2
	}
	upper = make([]float64, len(records))
	middle = make([]float64, len(records))
	lower = make([]float64, len(records))
	j := 0
	for j = length - 1; j < len(records) && math.IsNaN(records[j]); j++ {
	}
	ind.set(upper, 0, j, math.NaN())
	ind.set(middle, 0, j, math.NaN())
	ind.set(lower, 0, j, math.NaN())
	sum := 0.0
	for i := j; i < len(records); i++ {
		if i == j {
			for k := 0; k < length; k++ {
				sum += records[k]
			}
		} else {
			sum += records[i] - records[i-length]
		}
		ma := sum / float64(length)
		d := 0.0
		for k := i + 1 - length; k <= i; k++ {
			d += (records[k] - ma) * (records[k] - ma)
		}

		stdev := math.Sqrt(d / float64(length))
		tmp := stdev * multiplier
		upper[i] = ma + tmp
		middle[i] = ma
		lower[i] = ma - tmp
	}
	return upper, middle, lower
}

func (ind *Indicator) KDJ(
	records []Bar,
	length int,
	signalKLength int,
	signalDLength int,
) (k []float64, d []float64, j []float64) {
	if length == 0 {
		length = 9
	}
	if signalKLength == 0 {
		signalKLength = 3
	}
	if signalDLength == 0 {
		signalDLength = 3
	}

	k = make([]float64, len(records))
	d = make([]float64, len(records))
	j = make([]float64, len(records))

	hs := make([]float64, len(records))
	ls := make([]float64, len(records))

	for i := 0; i < len(records); i++ {
		hs[i] = records[i].High
		ls[i] = records[i].Low
	}

	for i := 0; i < len(records); i++ {
		if i >= (length - 1) {
			c := records[i].Close
			h := ind.cmp(hs, i-(length-1), i+1, math.Max)
			l := ind.cmp(ls, i-(length-1), i+1, math.Min)
			rsv := 100 * (c - l) / (h - l)
			k[i] = (rsv + float64(signalKLength-1)*k[i-1]) / float64(signalKLength)
			d[i] = (k[i] + float64(signalDLength-1)*d[i-1]) / float64(signalDLength)
		} else {
			k[i] = 50
			d[i] = 50
		}
		j[i] = 3*k[i] - 2*d[i]
	}

	for i := 0; i < length-1; i++ {
		k[i] = math.NaN()
		d[i] = math.NaN()
		j[i] = math.NaN()
	}

	return k, d, j
}

func (ind *Indicator) RSI(records []float64, length int) []float64 {
	if length == 0 {
		length = 14
	}

	rsi := make([]float64, len(records))
	ind.set(rsi, 0, len(rsi), math.NaN())
	if len(records) < length {
		return rsi
	}

	deltas := ind.moveDiff(records)
	seed := deltas[0:length]

	up, down := 0.0, 0.0
	for i := 0; i < len(seed); i++ {
		if seed[i] >= 0 {
			up += seed[i]
		} else {
			down += seed[i]
		}
	}
	up /= float64(length)
	down = -(down / float64(length))
	rs := func() float64 {
		if down != 0 {
			return up / down
		} else {
			return 0.0
		}
	}()
	rsi[length] = 100 - 100/(1+rs)
	delta, upval, downval := 0.0, 0.0, 0.0
	for i := length + 1; i < len(records); i++ {
		delta = deltas[i-1]
		if delta > 0 {
			upval = delta
			downval = 0.0
		} else {
			upval = 0.0
			downval = -delta
		}
		up = (up*float64(length-1) + upval) / float64(length)
		down = (down*float64(length-1) + downval) / float64(length)
		rs = up / down
		rsi[i] = 100 - 100/(1+rs)
	}
	return rsi
}
