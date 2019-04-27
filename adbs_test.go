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
	res := CalcUpper(-5, "x^2+5*x", 20)
	exp := 3.3588
	if res-exp > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcUpperC(t *testing.T) {
	res := CalcUpper(0, "x^(19)*(1-x)^(19)", 6.8917e-13)
	exp := 0.62864
	if res-exp > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcNewtonA(t *testing.T) {
	res := CalcNewton(2, true, "3*x^2+5*x+1", 20, 20, 0.001)
	exp := 2.68320080516111
	if math.Abs(res-exp) > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcNewtonAAcc(t *testing.T) {
	res := CalcNewton(2, true, "3*x^2+5*x+1", 20, 20, 1e-10)
	exp := 2.68320080516111
	if math.Abs(res-exp) > 1e-10 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcNewtonB(t *testing.T) {
	res := CalcNewton(math.Pi/3, false, "sin(x)", (math.Sqrt(3)-1)/2, 20, 0.001)
	exp := math.Pi / 6
	if math.Abs(res-exp) > 0.001 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}

func TestCalcNewtonBAcc(t *testing.T) {
	res := CalcNewton(math.Pi/3, false, "sin(x)", (math.Sqrt(3)-1)/2, 20, 1e-10)
	exp := math.Pi / 6
	if math.Abs(res-exp) > 1e-10 {
		t.Fatalf("Expected %v but received %v\n", exp, res)
	}
}
func BenchmarkUpperA(b *testing.B) {
	expr := "3*x^2+5*x+1"
	for i := 0; i < b.N; i++ {
		CalcNewton(2, true, expr, 20, 20, 0.000001)
	}
}

func BenchmarkNewtonA(b *testing.B) {
	expr := "3*x^2+5*x+1"
	for i := 0; i < b.N; i++ {
		CalcUpper(1, expr, 20)
	}
}
func BenchmarkLowerB(b *testing.B) {
	expr := "sin(x)"
	val := (math.Sqrt(3) - 1) / 2
	for i := 0; i < b.N; i++ {
		CalcNewton(2, false, expr, val, 20, 0.000001)
	}
}

func BenchmarkNewtonB(b *testing.B) {
	expr := "sin(x)"
	val := (math.Sqrt(3) - 1) / 2
	for i := 0; i < b.N; i++ {
		CalcUpper(1, expr, val)
	}
}
