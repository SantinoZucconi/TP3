package funciones

import (
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func CaminoMasCorto[K comparable](g TDAGrafo.Grafo[K], origen, destino K) []K {
	if !g.Dirigido() {
		panic("El grafo no es dirigido")
	}
	padres, _ := TDAGrafo.Bfs[K](g, origen, destino, true)
	camino := TDAGrafo.ReconstruirCamino[K](padres, origen, destino)
	return camino
}

func EnRango[K comparable](g TDAGrafo.Grafo[K], origen K, n int) []K {
	if !g.Dirigido() {
		panic("El grafo no es dirigido")
	}
	var none K
	res := []K{}
	_, orden := TDAGrafo.Bfs[K](g, origen, none, false)
	orden.Iterar(func(vertice K, orden int) bool {
		if orden == n {
			res = append(res, vertice)
		} else if orden > n {
			return false
		}
		return true
	})
	return res
}

func Comunidades[K comparable](g TDAGrafo.Grafo[K], pagina K) []K {
	label := TDADicc.CrearHash[K, int]()
	comunidades := TDADicc.CrearHash[int, []K]()
	vEntradas := vertices_entrada[K](g)
	for i, v := range g.ObtenerVertices() {
		label.Guardar(v, i)
	}
	for i := 0; i < 2; i++ {
		for _, v := range g.ObtenerVertices() {
			labelNeighbor := []int{}
			for _, w := range g.Adyacente(v) {
				labelNeighbor = append(labelNeighbor, label.Obtener(w))
			}
			if g.Dirigido() {
				if vEntradas.Pertenece(v) {
					colaAdy := vEntradas.Obtener(v)
					for !colaAdy.EstaVacia() {
						labelNeighbor = append(labelNeighbor, label.Obtener(colaAdy.Desencolar()))
					}
				}
			}
			newLabel := max_freq(labelNeighbor)
			label.Guardar(v, newLabel)
		}
	}
	label.Iterar(func(clave K, dato int) bool {
		if !comunidades.Pertenece(dato) {
			comunidades.Guardar(dato, []K{clave})
		} else {
			comunidad := comunidades.Obtener(dato)
			comunidad = append(comunidad, clave)
			comunidades.Guardar(dato, comunidad)
		}
		return true
	})
	return comunidades.Obtener(label.Obtener(pagina))
}

func max_freq(arr []int) int {
	maxFreqLabel := TDADicc.CrearHash[int, int]()
	var max int
	var maxLabel int
	for _, label := range arr {
		if !maxFreqLabel.Pertenece(label) {
			maxFreqLabel.Guardar(label, 1)
		} else {
			maxFreqLabel.Guardar(label, maxFreqLabel.Obtener(label)+1)
		}
	}
	maxFreqLabel.Iterar(func(clave, dato int) bool {
		if dato > max {
			max = dato
			maxLabel = clave
		}
		return true
	})
	return maxLabel
}

func vertices_entrada[K comparable](g TDAGrafo.Grafo[K]) TDADicc.Diccionario[K, TDACola.Cola[K]] {
	verticesEntrada := TDADicc.CrearHash[K, TDACola.Cola[K]]()
	for _, v := range g.ObtenerVertices() {
		for _, w := range g.Adyacente(v) {
			if !verticesEntrada.Pertenece(w) {
				verticesEntrada.Guardar(w, TDACola.CrearColaEnlazada[K]())
				verticesEntrada.Obtener(w).Encolar(v)
			} else {
				verticesEntrada.Obtener(w).Encolar(v)
			}
		}
	}
	return verticesEntrada
}
