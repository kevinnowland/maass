// Trying to implement a Complex type with arbitrary precision
// by bootstrapping up from math/big
//
// This has surely been done before and better, but hey
// this is just an exercise for myself
//

package maass

import (
	"math/big"
)

// the type that we really care about
type Complex struct {
	Re *big.Float
	Im *big.Float
}

// get a new complex number
// as with NewFloat this will default to precision 53
func NewComplex(a, b float64) *Complex {
	return &Complex{Re: big.NewFloat(a), Im: big.NewFloat(b)}
}

// get the minimum precision between real and imaginary parts
func (z *Complex) Prec() uint {
	if z.Re.Prec() < z.Im.Prec() {
		return z.Re.Prec()
	} else {
		return z.Im.Prec()
	}
}

// set z to a + b and return z
func (z *Complex) Add(a, b *Complex) *Complex {
	z.Re.Add(a.Re, b.Re)
	z.Im.Add(a.Im, b.Im)
	return z
}

// set z to a - b and return z
func (z *Complex) Sub(a, b *Complex) *Complex {
	z.Re.Sub(a.Re, b.Re)
	z.Im.Sub(a.Im, b.Im)
	return z
}

// set z to a*b and return z
func (z *Complex) Mul(a, b *Complex) *Complex {
	z.Re.Sub(big.NewFloat(0).Mul(a.Re, b.Re), big.NewFloat(0).Mul(a.Im, b.Im))
	z.Im.Add(big.NewFloat(0).Mul(a.Re, b.Im), big.NewFloat(0).Mul(a.Im, b.Re))
	return z
}

// set z = u + iv to u - iv
func (z *Complex) Conj() *Complex {
	z.Re.Add(big.NewFloat(0), z.Re)
	z.Im.Sub(big.NewFloat(0), z.Im)
	return z
}

// calculate |z| and return as *Complex
func (z *Complex) Abs() *Complex {
	realSquared := big.NewFloat(0).Mul(z.Re, z.Re)
	imSquared := big.NewFloat(0).Mul(z.Im, z.Im)
	realAbsSquared := big.NewFloat(0).Add(realSquared, imSquared)
	realAbs := big.NewFloat(0).Sqrt(realAbsSquared)
	abs := Complex{Re: realAbs, Im: big.NewFloat(0)}
	return &abs
}

// set z to 1 / z and return z
func (z *Complex) Inv() *Complex {
	invertedAbs := big.NewFloat(0).Quo(big.NewFloat(1), z.Abs().Re)
	z = z.Conj()
	z.Re.Mul(z.Re, invertedAbs)
	z.Im.Mul(z.Im, invertedAbs)
	return z
}

// set z to a / b and return z
func (z *Complex) Quo(a, b *Complex) *Complex {
	return NewComplex(0, 0).Mul(a, b.Inv())
}

// set z to -w and return z
func (z *Complex) Neg(w *Complex) *Complex {
	z.Re.Neg(w.Re)
	z.Im.Neg(w.Im)
	return z
}

// convert to complex128
func (z *Complex) Complex128() (complex128, big.Accuracy, big.Accuracy) {
	re, re_acc := z.Re.Float64()
	im, im_acc := z.Im.Float64()
	return complex(re, im), re_acc, im_acc
}
