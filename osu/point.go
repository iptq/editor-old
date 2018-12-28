package osu

type IntPoint struct {
	x, y int
}

func (p IntPoint) ToFloat() FloatPoint {
	return FloatPoint{
		x: float64(p.x),
		y: float64(p.y),
	}
}

type FloatPoint struct {
	x, y float64
}

func (p FloatPoint) ToInt() IntPoint {
	return IntPoint{
		x: int(p.x),
		y: int(p.y),
	}
}
