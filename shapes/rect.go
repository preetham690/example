package shape

type Rect struct {
	L, B float32
}

//method

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func AreaOf(r *Rect) float32 {
	return r.L * r.B
}
