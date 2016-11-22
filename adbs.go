package adbs

import (
	calculus "github.com/TheDemx27/calculus"
	"time"
	//"math"
)

func CalcUpperTimeout(lower float64, expr string, k float64, timeout float64) float64 {
	timeout = timeout * float64(time.Second)
	c := make(chan float64, 1)
	//ret := make(chan float64)
	go func() {
		f := calculus.NewFunc(expr)
		v := float64(0)
		for uppr := lower; v != k; uppr += float64(0.000001) {
			v = f.AntiDiff(lower, uppr)
			if v-k < 0.000001 {
				//ret <- uppr
				c <- uppr
				return
			} else if uppr == uppr+float64(0.000001) {
				break
			}
		}
		//ret <- lower
		c <- lower
		return
	}()
	select {
	case ret := <-c:
		return ret
		//return <-ret
	case <-time.After(time.Duration(timeout)):
		return lower
	}
}

func CalcUpper(lower float64, expr string, k float64) float64 {
	return CalcUpperTimeout(lower, expr, k, 1)
}

func CalcLowerTimeout(upper float64, expr string, k float64, timeout float64) float64 {
	timeout = timeout * float64(time.Second)
	c := make(chan float64, 1)
	go func() {
		f := calculus.NewFunc(expr)
		v := float64(0)
		for lowr := upper; v != k; lowr -= float64(0.000001) {
			v = f.AntiDiff(lowr, upper)
			if v-k < 0.000001 {
				c <- lowr
				return
			} else if lowr/(lowr+float64(000001)) < 0.0001 {
				break
			}
		}
		c <- upper
		return
	}()
	select {
	case ret := <-c:
		return ret
		//return <-ret
	case <-time.After(time.Duration(timeout)):
		return upper
	}
}

func CalcLower(upper float64, expr string, k float64) float64 {
	return CalcLowerTimeout(upper, expr, k, 1)
}
