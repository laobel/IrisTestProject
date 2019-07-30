package ellipsoid

import "math"

/**
(x / a)^2 + (y / b)^2 + (z / c)^2 = 1
 */
type Ellipsoid struct {
	x float64
	y float64
	z float64
}

func WGS84Ellipsoid() *Ellipsoid {
	return &Ellipsoid{
		x: 6378137.0,
		y: 6378137.0,
		z:6356752.3142451793,
	}
}

func (ellipsoid *Ellipsoid) MaximumRadius() float64{
	return math.Max(ellipsoid.x, math.Max(ellipsoid.y, ellipsoid.z))
}

func (ellipsoid *Ellipsoid) SemimajorAxisTimesPi() float64{
	return ellipsoid.MaximumRadius() * math.Pi
}

