package main

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) Rotate90X() Point3D {
	return Point3D{
		X: p.X,
		Y: -p.Z,
		Z: p.Y,
	}
}

func (p Point3D) Rotate90Y() Point3D {
	return Point3D{
		X: p.Z,
		Y: p.Y,
		Z: -p.X,
	}
}

func (p Point3D) Rotate90Z() Point3D {
	return Point3D{
		X: -p.Y,
		Y: p.X,
		Z: p.Z,
	}
}

func (p Point3D) Sub(o Point3D) Point3D {
	return Point3D{
		X: p.X - o.X,
		Y: p.Y - o.Y,
		Z: p.Z - o.Z,
	}
}
