package usecase

import (
	"fmt"
	"math"

	"github.com/unawaretub86/top-secret/internal/config/errors"
)

type TriangulationUseCase struct{}

func NewTriangulationUseCase() *TriangulationUseCase {
	return &TriangulationUseCase{}
}

// Esta funcion se encarga de calcular la ubicacion x y y basandose en la distancia de los 3 satellites entregados
func (usecase *TriangulationUseCase) GetLocation(requestID string, distances ...float32) (*float32, *float32, error) {
	if len(distances) != 3 {
		fmt.Printf("[RequestId: %s][%v]", requestID, errors.ErrInvalidSatellites)
		return nil, nil, errors.ErrInvalidSatellites
	}

	// Coordernadas de los puntos de referencia
	var kenobi = []float32{-500, -200}
	var skywalker = []float32{100, -100}
	var sato = []float32{500, 100}

	x, y := triangulation(kenobi, skywalker, sato, distances)

	return &x, &y, nil
}

// Triangulation calcula las coordenadas (x, y) de un punto mediante el metodo de triangulación
// utilizando las distancias a tres puntos de referencia: kenobi, skywalker y sato.
func triangulation(kenobi, skywalker, sato, distances []float32) (x, y float32) {

	// Variables A, B, C, D, E, F se usan para los cálculos de la triangulación
	A := float64(2*skywalker[0] - 2*kenobi[0])
	B := float64(2*skywalker[1] - 2*kenobi[1])
	C := math.Pow(float64(distances[0]), 2) - math.Pow(float64(distances[1]), 2) - math.Pow(float64(kenobi[0]), 2) + math.Pow(float64(skywalker[0]), 2) - math.Pow(float64(kenobi[1]), 2) + math.Pow(float64(skywalker[1]), 2)
	D := float64(2*sato[0] - 2*skywalker[0])
	E := float64(2*sato[1] - 2*skywalker[1])
	F := math.Pow(float64(distances[1]), 2) - math.Pow(float64(distances[2]), 2) - math.Pow(float64(skywalker[0]), 2) + math.Pow(float64(sato[0]), 2) - math.Pow(float64(skywalker[1]), 2) + math.Pow(float64(sato[1]), 2)

	// Fórmula para calcular las coordenadas (x, y) utilizando los valores calculados anteriormente
	x = float32((C*E - F*B) / (E*A - B*D))
	y = float32((C*D - A*F) / (B*D - A*E))

	return x, y
}
