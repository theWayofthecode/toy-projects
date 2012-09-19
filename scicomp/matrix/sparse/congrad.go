package sparse

import (
	"../../matrix"
)

type idVectors struct {
	id int
	x matrix.Vector
	r matrix.Vector
}

var maxIter = 30
var threads = 2
var er = 0.05

func ConjugateGradient(A matrix.Matrix, b matrix.Vector) matrix.Vector {
	var iv idVectors
	n := len(b)
	x := matrix.VecAlloc(n)
	r := matrix.VecOpVec(b, matrix.MatDotVec(A, x), matrix.Minus)
	p := matrix.VecAlloc(n)
	copy(p, r)
	stride := len(b) / threads
	in := make(chan idVectors)
	
	for i := 0; i < maxIter; i++ {
		rTr := matrix.VecDotVec(r, r)
		Ap := matrix.MatDotVec(A, p)
		a := (matrix.VecDotVec(r, r) / matrix.VecDotVec(p, Ap))
		for	id := 0; id < threads; id++ {		
			start := id * stride
			end := start + stride
			go oneIter(a, r[start:end], x[start:end], p[start:end], Ap[start:end], in, id)
		}
		for	id := 0; id < threads; id++ {
			iv = <- in
			start := iv.id * stride
			end := start + stride
			copy(x[start:end], iv.x)
			copy(r[start:end], iv.r)
		}
		if (matrix.Norm(r) < er) {
			break;
		}

		β := matrix.VecDotVec(r, r) / rTr
		p.Scale(β)
		p.OpVec(r, matrix.Plus)
	}
	return x
}

func oneIter(a float64, r, x, p, Ap matrix.Vector, out chan idVectors, id int) {
	var iv idVectors
	iv.id = id
	ap := matrix.VecScale(p, a)
	iv.x = matrix.VecOpVec(x, ap, matrix.Plus)
	Ap.Scale(a)
	iv.r = matrix.VecOpVec(r, Ap, matrix.Minus)
	out <- iv
}