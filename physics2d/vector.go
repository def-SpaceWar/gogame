package physics2d

import "math"

type Vector struct {
	X float32
	Y float32
}

func Vec(x, y float32) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

func (v *Vector) Clone() Vector {
	n := *v
	return n
}

func (v *Vector) Add(o *Vector) *Vector {
	v.X += o.X
	v.Y += o.Y
	return v
}

func (v *Vector) Subtract(o *Vector) *Vector {
	v.X -= o.X
	v.Y -= o.Y
	return v
}

func (v *Vector) Scale(n float32) *Vector {
	v.X *= n
	v.Y *= n
	return v
}

func (v *Vector) ScaleVec(o *Vector) *Vector {
	v.X *= o.X
	v.Y *= o.Y
	return v
}

func (v *Vector) MagnitudeSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector) Magnitude() float32 {
	return float32(math.Sqrt(float64(v.MagnitudeSquared())))
}

func (v *Vector) Normalize() *Vector {
	magnitude := v.Magnitude()
	v.X /= magnitude
	v.Y /= magnitude
	return v
}

func (v *Vector) Normal() *Vector {
	v.X, v.Y = -v.Y, v.X
	return v
}

func (v *Vector) Dot(o *Vector) float32 {
	return v.X*o.X + v.Y*o.Y
}

func (v *Vector) Cross(o *Vector) float32 {
	return v.X*o.Y - v.Y*o.X
}

func (v *Vector) Rotate(angle float32) *Vector {
	v.X, v.Y = v.X*float32(math.Cos(float64(angle)))-v.Y*float32(math.Sin(float64(angle))), v.Y*float32(math.Cos(float64(angle)))+v.X*float32(math.Sin(float64(angle)))
    return v
}
