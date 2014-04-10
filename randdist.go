package simulator

import (
	"math/rand"
	"math"
)

func Poisson(min, max, mean int) (func() int) {
	L := math.Exp(float64(-mean))
	return func() int {
		p := rand.Float64()
		var k int
		for k = 0; p > L; k++ {
			p *= rand.Float64()
		}
		return k
	}
}

func Exponential(min,max,mean int) (func() int) {
	return func() int {
		lamb := 1 / mean
		var samp int
		for samp = int(rand.ExpFloat64() / float64(lamb));
			samp < min || samp > max;
			samp = int(rand.ExpFloat64() / float64(lamb)) {
			}
		return samp
	}
}

func Normal(min, max, mean, dev int) (func() int) {
	return func() int {
		var f int
		for f = int(rand.NormFloat64() * float64(dev) + float64(mean));
			f < min || f > max;
			f = int(rand.NormFloat64() * float64(dev) + float64(mean)) {
			}
		return f

	}
}

func Uniform(min,max int) (func () int) {
	return func() int {
		return rand.Int() % (max - min) + min
	}
}
