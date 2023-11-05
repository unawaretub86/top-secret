package request

// Contains es una funcion que verifica si un elemento esta presente en un slice.
// Devuelve un booleano dependiendo si se encuentra el elmento a buscar.
func Contains(slice []string, item string) bool {
	// map para realizar búsquedas eficientes de elementos en el slice.
	itemMap := make(map[string]bool)

	// Itera a traves de cada elemento del slice y agrega cada elemento al mapa.
	for _, s := range slice {
		itemMap[s] = true
	}

	// Comprueba si el elemento item se encuentra en el map.
	return itemMap[item]
}
