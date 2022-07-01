// let's compute mobius transformations defined by
// 2x2 matrices with integer entries

package maass

import (
    "math/big"
)

// the 2x2 matrix we care about
// (a b
//  c d)
type Matrix struct {
    a *big.Float
    b *big.Float
    c *big.Float
    d *big.Float
}

// get a new matrix from floats
// will default to precision 53
func NewMatrix(a, b, c, d float64) *Matrix {
    return &Matrix{a: big.NewFloat(a), b: big.NewFloat(b), c: big.NewFloat(c), d: big.NewFloat(d)}
}

// get a new Matrix from integers
// will default to precision 53
func NewIntMatrix(a, b, c, d int64) *Matrix {
    aFloat := float64(a)
    bFloat := float64(b)
    cFloat := float64(c)
    dFloat := float64(d)

    return NewMatrix(aFloat, bFloat, cFloat, dFloat)
}

// find  determinant
func (g *Matrix) Det() *big.Float {
    part1 := big.NewFloat(0).Mul(g.a, g.d)
    part2 := big.NewFloat(0).Mul(g.b, g.c)
    return big.NewFloat(0).Sub(part1, part2)
}

// compute Mobius transformation
func Mobius (g *Matrix, z *Complex) *Complex {
    topRe := big.NewFloat(0).Add(big.NewFloat(0).Mul(g.a, z.Re), g.b)
    topIm := big.NewFloat(0).Mul(g.a, z.Im)
    top := Complex{Re: topRe, Im: topIm}

    bottomRe := big.NewFloat(0).Add(big.NewFloat(0).Mul(g.c, z.Re), g.d)
    bottomIm := big.NewFloat(0).Mul(g.c, z.Im)
    bottom := Complex{Re: bottomRe, Im: bottomIm}

    return NewComplex(0, 0).Quo(&top, &bottom)
}
