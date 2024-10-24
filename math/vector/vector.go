package vector

import "math"

type Vec3 [3]float64

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
	}
}

func (v Vec3) Scale(r float64) Vec3 {
	return Vec3{
		v[0] * r,
		v[1] * r,
		v[2] * r,
	}
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
	}
}

func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v Vec3) Norm() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v1 Vec3) Dist(v2 Vec3) float64 {
	return v1.Sub(v2).Norm()
}
