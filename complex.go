// Trying to implement a Complex type with arbitrary precision
// by bootstrapping up from math/big
//
// This has surely been done before and better, but hey
// this is just an exercise for myself
//

package complex

import (
	"math/big"
)

// the type that we really care about
type Complex struct {
	re *big.Float
	im *big.Float
}

// get a new complex number
// as with NewFloat this will default to precision 53
func NewComplex(a, b float64) *Complex {
	return &Complex{re: big.NewFloat(a), im: big.NewFloat(b)}
}

// get the minimum precision between real and imaginary parts
func (z *Complex) Prec() uint {
	if z.re.Prec() < z.im.Prec() {
		return z.re.Prec()
	} else {
		return z.im.Prec()
	}
}

// set z to a + b and return z
func (z *Complex) Add(a, b *Complex) *Complex {
	z.re.Add(a.re, b.re)
	z.im.Add(a.im, b.im)
	return z
}

// set z to a - b and return z
func (z *Complex) Sub(a, b *Complex) *Complex {
	z.re.Sub(a.re, b.re)
	z.im.Sub(a.im, b.im)
	return z
}

// set z to a*b and return z
func (z *Complex) Mul(a, b *Complex) *Complex {
	z.re.Sub(big.NewFloat(0).Mul(a.re, b.re), big.NewFloat(0).Mul(a.im, b.im))
	z.im.Add(big.NewFloat(0).Mul(a.re, b.im), big.NewFloat(0).Mul(a.im, b.re))
	return z
}

// set z = u + iv to u - iv
func (z *Complex) Conj() *Complex {
	z.re.Add(big.NewFloat(0), z.re)
	z.im.Sub(big.NewFloat(0), z.im)
	return z
}

// calculate |z| and return as *Complex
func (z *Complex) Abs() *Complex {
	realSquared := big.NewFloat(0).Mul(z.re, z.re)
	imSquared := big.NewFloat(0).Mul(z.im, z.im)
	realAbsSquared := big.NewFloat(0).Add(realSquared, imSquared)
	realAbs := big.NewFloat(0).Sqrt(realAbsSquared)
	abs := Complex{re: realAbs, im: big.NewFloat(0)}
	return &abs
}

// set z to 1 / z and return z
func (z *Complex) Inv() *Complex {
	invertedAbs := big.NewFloat(0).Quo(big.NewFloat(1), z.Abs().re)
	z = z.Conj()
	z.re.Mul(z.re, invertedAbs)
	z.im.Mul(z.im, invertedAbs)
	return z
}

// set z to a / b and return z
func (z *Complex) Quo(a, b *Complex) *Complex {
	return NewComplex(0, 0).Mul(a, b.Inv())
}

// set z to -w and return z
func (z *Complex) Neg(w *Complex) *Complex {
	z.re.Neg(w.re)
	z.im.Neg(w.im)
	return z
}

// convert to complex128
func (z *Complex) Complex128() (complex128, big.Accuracy, big.Accuracy) {
    re, re_acc := z.re.Float64()
    im, im_acc := z.im.Float64()
    return complex(re, im), re_acc, im_acc
}
