package adbs

import "testing"
import "math"

func TestCalcUpperA(t *testing.T) {
	res := CalcUpper(1, "x", 20)
	exp := math.Sqrt(41.0)
	if res-exp > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcUpperB(t *testing.T) {
	res := CalcUpper(-5, "x^2+5x", 20)
	exp := 3.3588
	if res-exp > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcUpperC(t *testing.T) {
	res := CalcUpper(0, "x^(19)*(1-x)^(19)", 6.8917E-13)
	exp := 0.62864
	if res-exp > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcLowerA(t *testing.T) {
	res := CalcLower(math.Sqrt(41)-5, "3x^2+30x48", -164*math.Sqrt(41))
	exp := -5 - math.Sqrt(41)
	if res-exp > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}
