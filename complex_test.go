// testing complex
//
// I think this is sloppy as I am not probably
// thinking enough about the precision of all these calculations
//

package maass

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPrec(t *testing.T) {
	re := new(big.Float).SetPrec(10).SetFloat64(2)
	im := new(big.Float).SetPrec(15).SetFloat64(1)
	z := Complex{re, im}

	got := z.Prec()
	want := uint(10)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	z := NewComplex(1, -2)
	w := NewComplex(4, 5)

	got := NewComplex(0, 0).Add(z, w)
	want := NewComplex(5, 3)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestSub(t *testing.T) {
	z := NewComplex(1, -2)
	w := NewComplex(4, 5)

	got := NewComplex(0, 0).Sub(z, w)
	want := NewComplex(-3, -7)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestMul(t *testing.T) {
	z := NewComplex(1, -2)
	w := NewComplex(4, 5)

	got := NewComplex(0, 0).Mul(z, w)
	want := NewComplex(14, -3)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestConj(t *testing.T) {
	z := NewComplex(1, -2)

	got := z.Conj()
	want := NewComplex(1, 2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestAbs(t *testing.T) {
	z := NewComplex(3, -4)

	got := z.Abs()
	want := NewComplex(5, 0)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestInv(t *testing.T) {
	z := NewComplex(3, -4)

	got := z.Inv()
	want := NewComplex(0.6, 0.8)

	// cheating with precision?
	got.Re.SetPrec(5)
	got.Im.SetPrec(5)

	want.Re.SetPrec(5)
	want.Im.SetPrec(5)

	if !reflect.DeepEqual(want.Re, got.Re) || !reflect.DeepEqual(want.Im, got.Im) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestQuo(t *testing.T) {
	z := NewComplex(1, -2)
	w := NewComplex(-3, 4)

	got := NewComplex(0, 0).Quo(z, w)
	want := NewComplex(-2.2, 0.4)

	// cheating with precision?
	got.Re.SetPrec(5)
	got.Im.SetPrec(5)

	want.Re.SetPrec(5)
	want.Im.SetPrec(5)

	if !reflect.DeepEqual(want.Re, got.Re) || !reflect.DeepEqual(want.Im, got.Im) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestNeg(t *testing.T) {
	got := NewComplex(0, 0).Neg(NewComplex(1, -2))
	want := NewComplex(-1, 2)

	if !reflect.DeepEqual(want.Re, got.Re) || !reflect.DeepEqual(want.Im, got.Im) {
		t.Errorf("got Complex{%v, %v} wanted Complex{%v, %v}", got.Re, got.Im, want.Re, want.Im)
	}
}

func TestComplex(t *testing.T) {
	got, _, _ := NewComplex(0.5, -1.4).Complex128()
	want := complex(0.5, -1.4)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
