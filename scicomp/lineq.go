package main

import (
	"fmt"
	"math/rand"
	"./matrix"
	"./matrix/sparse"
)

func main() {
    t := matrix.MatAlloc(6, 6)
    t.InitTridiag(1, 4, 1)
    if (t.IsSymmetric()) {
    	fmt.Println("t is symmetric")
    }
 	var y matrix.Vector
 	y = append(y, 5, 6, 6, 6, 6, 5)
 	x := sparse.Jacobi(t, y)
	fmt.Println(x)
 	x = sparse.ConjugateGradient(t, y)
	fmt.Println(x)
}

func make_matrix(m, n int) matrix.Matrix {
	a := make([][]float64, m)
	
	for i := 0; i < m; i++ {
		a[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			a[i][j] =  float64 (rand.Intn(10))
		}
	}
	return a
}

func make_vector(n int) matrix.Vector {
	v := make([]float64, n)
	for i := 0; i < len(v); i++ {
		if i == 0 || i == n - 1 {
			v[i] = 5
		} else {
			v[i] = 6
		}
	}
	return v
}