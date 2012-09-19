package matrix

import (
	"math"
)

type Matrix [][]float64
type Vector []float64

func (m Matrix) InitTridiag(b float64, d float64, a float64) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if i - 1 == j {
				m[i][j] = b
			} else if i == j {
				m[i][j] = d
			} else if i + 1 == j {
				m[i][j] = a
			}
		}
	}
}

func MatAlloc(m int, n int) Matrix {
	mat := make(Matrix, m)
	for i := 0; i < m; i++ {
		mat[i] = make(Vector, n)
	}
	return mat
}

func VecAlloc(n int) Vector {
	return make(Vector, n)
}

func VecDotVec(v1, v2 Vector) float64 {
	var dot float64
	for i := 0; i < len(v1); i++ {
		dot += v1[i] * v2[i]
	}
	return dot
}

func MatDotVec(m Matrix, v Vector) Vector {
	dot := make([]float64, len(m))
	
	for i := 0; i < len(m); i++ {
		dot[i] = VecDotVec(m[i], v)
	}
	return dot
}

func VecOpVec(v1, v2 Vector, op func(float64, float64) float64) Vector {
	v := make([]float64, len(v1))
	
	for i := 0; i < len(v1); i++ {
		v[i] = op(v1[i], v2[i])
	}
	return v
}

func LDU(m Matrix) (d Vector, lu Matrix) {
	lu = make([][]float64, len(m))
	d = make([]float64, len(m))

	for i := 0; i < len(m); i++ {
		lu[i] = make([]float64, len(m))
		copy(lu[i], m[i])
		d[i] = lu[i][i]
		lu[i][i] = 0
	}
	return
}
	
func Norm(v Vector) float64 {
	var sum float64
	sum = 0
	for i := 0; i < len(v); i++ {
		sum += v[i] * v[i]
	}
	return math.Sqrt(sum)
}
	