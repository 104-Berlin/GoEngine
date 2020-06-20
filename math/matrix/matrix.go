package matrix

import "github.com/104-Berlin/GoEngine/math/vector"

type Mat4 [4][4]float64

func Id4() Mat4 {
	return Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func (m1 *Mat4) Add(m2 *Mat4) (r Mat4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return
}

func (m1 *Mat4) Mult(m2 *Mat4) (r Mat4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := 0.0
			for k := 0; k < 4; k++ {
				sum += m1[i][k] * m2[k][j]
			}
			r[i][j] = sum
		}
	}
	return
}

func (m *Mat4) Trans(v *vector.Vec3) (r vector.Vec3) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r[i] += m[i][j] * v[j]
		}
		r[i] += m[i][3]
	}
	return
}
