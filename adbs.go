package adbs

import (
	calculus "github.com/TheDemx27/calculus"
	"time"
	//"math"
)

func CalcUpper(lower float64, expr string, k float64) float64 {
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
	case <-time.After(time.Second * 1):
		return lower
	}
}

func CalcLower(upper float64, expr string, k float64) float64 {
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
	case <-time.After(time.Second * 1):
		return upper
	}
}
