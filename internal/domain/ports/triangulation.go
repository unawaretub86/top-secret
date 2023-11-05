package ports

type TriangulationPort interface {
	GetLocation(...float32) (float32, float32)
}
