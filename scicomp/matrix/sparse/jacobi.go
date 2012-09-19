package sparse

import (
	"../../matrix"
)

var MaxIter = 100
var Threads = 2
var Er = 0.05

type IdVector struct {
	id int
	v matrix.Vector
}

func Jacobi(A matrix.Matrix, b matrix.Vector) matrix.Vector {
	var iv IdVector
	x := make(matrix.Vector, len(A))
	diff := make(matrix.Vector, len(A))
	in := make(chan IdVector)
	stride := len(b) / Threads
	d, l_u := matrix.LDU(A)
	for it := 0; it < MaxIter; it++ {
		for	id := 0; id < Threads; id++ {
			start := id * stride
			end := start + stride
			go oneIteration(l_u[start:end], b[start:end], d[start:end], x, in, id)
		}
		for	id := 0; id < Threads; id++ {
			iv = <- in
			start := iv.id * stride
			end := start + stride
			copy(diff[start:end], 
				matrix.VecOpVec(iv.v, x[start:end], matrix.Minus))
			copy(x[start:end], iv.v)
		}
		if (matrix.Norm(diff) < Er) {
			break;
		}
	}
	return x
}

func oneIteration(l_u matrix.Matrix, b, d, x matrix.Vector, out chan IdVector, id int) {
		right := matrix.MatDotVec(l_u, x)
		right = matrix.VecOpVec(b, right, func(v1 float64, v2 float64) float64 {return v1 - v2})
		var iv IdVector
		iv.id = id
		iv.v = matrix.VecOpVec(right, d, func(v1 float64, v2 float64) float64 {return v1 / v2})
		out <- iv
}