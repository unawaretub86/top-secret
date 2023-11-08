package ports

type TriangulationPort interface {
	GetLocation(string, ...float32) (*float32, *float32, error)
}
