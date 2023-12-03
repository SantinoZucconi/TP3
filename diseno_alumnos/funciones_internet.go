package internet

import (
	"fmt"
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
	ERROR "tp3/errores"
)

const d float64 = 0.85

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func difAbs(x, y float64) float64 {
	if x > y {
		return x - y
	}

	return y - x
}

func difPorcentual(original, nuevo float64) float64 {
	return difAbs(original, nuevo) / original * 100
}

type verticePR[K comparable] struct {
	vertice  K
	pageRank float64
}

func BFS[K comparable](g TDAGrafo.GrafoNoPesado[K], origen K, cond func(K, TDADicc.Diccionario[K, int]) bool) (TDADicc.Diccionario[K, K], TDADicc.Diccionario[K, int]) {
	var none K
	padres := TDADicc.CrearHash[K, K]()
	visitados := TDADicc.CrearHash[K, bool]()
	orden := TDADicc.CrearHash[K, int]()
	q := TDACola.CrearColaEnlazada[K]()

	q.Encolar(origen)
	padres.Guardar(origen, none)
	orden.Guardar(origen, 0)
	visitados.Guardar(origen, true)

	for !q.EstaVacia() {
		v := q.Desencolar()
		for _, w := range g.Adyacente(v) {
			if !visitados.Pertenece(w) {
				padres.Guardar(w, v)
				orden.Guardar(w, orden.Obtener(v)+1)
				visitados.Guardar(w, true)
				if cond(w, orden) {
					return padres, orden
				}
				q.Encolar(w)
			}
		}
	}
	return padres, orden
}

func _dfs[K comparable](g TDAGrafo.GrafoNoPesado[K], vertice, padre K, padres TDADicc.Diccionario[K, K], orden TDADicc.Diccionario[K, int], visitados TDADicc.Diccionario[K, bool]) {
	var none K
	if padre != none {
		padres.Guardar(vertice, padre)
		orden.Guardar(vertice, orden.Obtener(padre)+1)
	}
	visitados.Guardar(vertice, true)

	adyacentes := g.Adyacente(vertice)
	cantAdyacentes := len(adyacentes)

	for i := 0; i < cantAdyacentes; i++ {
		w := adyacentes[i]
		if !visitados.Obtener(w) {
			_dfs(g, w, vertice, padres, orden, visitados)
		}
	}
}

func DFS[K comparable](g TDAGrafo.GrafoNoPesado[K], origen K) (TDADicc.Diccionario[K, K], TDADicc.Diccionario[K, int]) {
	var none K
	padres := TDADicc.CrearHash[K, K]()
	visitados := TDADicc.CrearHash[K, bool]()
	orden := TDADicc.CrearHash[K, int]()

	padres.Guardar(origen, none)
	orden.Guardar(origen, 0)
	visitados.Guardar(origen, true)

	_dfs(g, origen, none, padres, orden, visitados)

	return padres, orden
}

func GradoDeEntrada[K comparable](g TDAGrafo.GrafoNoPesado[K]) TDADicc.Diccionario[K, int] {
	grEntrada := TDADicc.CrearHash[K, int]()
	for _, v := range g.ObtenerVertices() {
		grEntrada.Guardar(v, 0)
	}
	for _, v := range g.ObtenerVertices() {
		for _, w := range g.Adyacente(v) {
			grado_anterior := grEntrada.Obtener(w)
			grEntrada.Guardar(w, grado_anterior+1)
		}
	}
	return grEntrada
}

func GradoDeSalida[K comparable](g TDAGrafo.GrafoNoPesado[K]) TDADicc.Diccionario[K, int] {
	grSalida := TDADicc.CrearHash[K, int]()
	for _, v := range g.ObtenerVertices() {
		cantAdyacentes := len(g.Adyacente(v))
		grSalida.Guardar(v, cantAdyacentes)
	}

	return grSalida
}

func ReconstruirCamino[K comparable](padres TDADicc.Diccionario[K, K], in, fin K) ([]K, error) {
	var none K
	p := TDAPila.CrearPilaDinamica[K]()
	v := fin
	res := []K{}
	for v != in {
		p.Apilar(v)
		v = padres.Obtener(v)
		if v == none {
			return []K{}, &ERROR.ErrorNoExisteRecorrido{}
		}
	}
	p.Apilar(in)
	for !p.EstaVacia() {
		res = append(res, p.Desapilar())
	}
	return res, nil
}

////

func CaminoMinimo[K comparable](g TDAGrafo.GrafoNoPesado[K], origen, destino K) ([]K, error) {
	padres, _ := BFS(g, origen, func(vertice K, orden TDADicc.Diccionario[K, int]) bool {
		return vertice == destino
	})
	camino, err := ReconstruirCamino(padres, origen, destino)
	return camino, err
}

func TodosEnRango[K comparable](g TDAGrafo.GrafoNoPesado[K], origen K, rango int) int {
	_, orden := BFS(g, origen, func(vertice K, orden TDADicc.Diccionario[K, int]) bool {
		return orden.Obtener(vertice) <= rango
	})

	contador := 0
	for iter := orden.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, n := iter.VerActual()
		if n == rango {
			contador++
		}
	}

	return contador
}

func Diametro[K comparable](g TDAGrafo.GrafoNoPesado[K]) []K {
	diametro := 0
	var verticeInicio, verticeMasLejos K
	var padresDiametro TDADicc.Diccionario[K, K]

	vertices := g.ObtenerVertices()
	for i := 0; i < g.Cantidad(); i++ {
		v := vertices[i]
		padres, orden := BFS(g, v, func(vertice K, orden TDADicc.Diccionario[K, int]) bool {
			return false
		})

		for iter := orden.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			w, n := iter.VerActual()
			if n > diametro {
				verticeInicio = v
				verticeMasLejos = w
				diametro = n
				padresDiametro = padres
			}
		}
	}

	camino, _ := ReconstruirCamino(padresDiametro, verticeInicio, verticeMasLejos)

	return camino
}

////

func _dfs_cfc[K comparable](g TDAGrafo.GrafoNoPesado[K], v K, visitados, apilados TDADicc.Diccionario[K, K], orden, masbajo TDADicc.Diccionario[K, int], pila TDAPila.Pila[K], cfc *[]TDADicc.Diccionario[K, K], contador_global *int) {
	visitados.Guardar(v, v)
	pila.Apilar(v)
	apilados.Guardar(v, v)
	orden.Guardar(v, *contador_global)
	masbajo.Guardar(v, *contador_global)
	*contador_global++

	adyacentes := g.Adyacente(v)
	for i := 0; i < len(adyacentes); i++ {
		w := adyacentes[i]
		if !visitados.Pertenece(w) {
			_dfs_cfc(g, w, visitados, apilados, orden, masbajo, pila, cfc, contador_global)
		}
		if apilados.Pertenece(w) {
			masbajo.Guardar(v, min(masbajo.Obtener(v), masbajo.Obtener(w)))
		}
	}

	if orden.Obtener(v) == masbajo.Obtener(v) {
		fin := false
		nueva_cfc := TDADicc.CrearHash[K, K]()
		for !fin {
			max := pila.VerTope()
			if max == v {
				fin = true
			}
			nueva_cfc.Guardar(max, max)
			apilados.Borrar(max)
			pila.Desapilar()
		}

		*cfc = append(*cfc, nueva_cfc)
	}

}

func CFC[K comparable](g TDAGrafo.GrafoNoPesado[K]) ([][]K, TDADicc.Diccionario[K, int]) {
	var v K

	orden := TDADicc.CrearHash[K, int]()
	masbajo := TDADicc.CrearHash[K, int]()
	visitados := TDADicc.CrearHash[K, K]()
	pila := TDAPila.CrearPilaDinamica[K]()
	apilados := TDADicc.CrearHash[K, K]()

	cfc := make([]TDADicc.Diccionario[K, K], 0)
	pertenencia := TDADicc.CrearHash[K, int]()

	listas := make([][]K, 0)

	cont := 0
	contador_global := &cont

	vertices := g.ObtenerVertices()
	for i := 0; i < g.Cantidad(); i++ {
		v = vertices[i]
		if !visitados.Pertenece(v) {
			_dfs_cfc(g, v, visitados, apilados, orden, masbajo, pila, &cfc, contador_global)
		}
	}

	for i := 0; i < g.Cantidad(); i++ {
		v = vertices[i]
		for d := 0; d < len(cfc); d++ {
			comp := cfc[d]
			if comp.Pertenece(vertices[i]) {
				pertenencia.Guardar(v, d)
			}
		}
	}

	for d := 0; d < len(cfc); d++ {
		lista := make([]K, 0)
		comp := cfc[d]
		comp.Iterar(func(clave, dato K) bool {
			lista = append(lista, clave)
			return true
		})
	}

	return listas, pertenencia
}

// Comunidades

func max_freq(arr []int) int {
	maxFreqLabel := TDADicc.CrearHash[int, int]()
	var max int
	var maxLabel int
	for _, label := range arr {
		if !maxFreqLabel.Pertenece(label) {
			maxFreqLabel.Guardar(label, 0)
		}
		maxFreqLabel.Guardar(label, maxFreqLabel.Obtener(label)+1)
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
			}
			verticesEntrada.Obtener(w).Encolar(v)
		}
	}
	return verticesEntrada
}

// Cluserting

func ClusteringIndividual[K comparable](g TDAGrafo.GrafoNoPesado[K], vertice K) float64 {
	adyacentes := g.Adyacente(vertice)
	cantAdyacentes := len(adyacentes)
	unionAdyacentes := 0

	for j := 0; j < cantAdyacentes; j++ {
		w1 := adyacentes[j]
		for k := 0; k < cantAdyacentes; k++ {
			w2 := adyacentes[k]
			if g.HayArista(w1, w2) {
				unionAdyacentes++
			}
		}
	}

	return float64(unionAdyacentes) / float64(cantAdyacentes*(cantAdyacentes-1))
}

func ClusteringRed[K comparable](g TDAGrafo.GrafoNoPesado[K]) float64 {
	vertices := g.ObtenerVertices()
	cantidadVertices := g.Cantidad()
	var total float64

	for i := 0; i < cantidadVertices; i++ {
		v := vertices[i]
		total += ClusteringIndividual(g, v)
	}

	return total / float64(cantidadVertices)
}

// Page Rank

func PageRank[K comparable](g TDAGrafo.GrafoNoPesado[K]) []K {
	vertices := g.ObtenerVertices()
	cantidadVertices := g.Cantidad()
	pageRanks := TDADicc.CrearHash[K, float64]()
	pageRanksActualizados := TDADicc.CrearHash[K, float64]()
	pageOrdenadas := make([]K, 0)

	heap := TDAHeap.CrearHeap[verticePR[K]](func(v1, v2 verticePR[K]) int {
		comp := v2.pageRank*float64(cantidadVertices) - v1.pageRank*float64(cantidadVertices)
		if comp < 0 {
			return -1
		} else if comp > 0 {
			return 1
		}
		return 0
	})

	for i := 0; i < cantidadVertices; i++ {
		v := vertices[i]
		pageRanks.Guardar(v, float64(1/cantidadVertices))
	}

	seguirIterando := true

	for seguirIterando {
		seguirIterando = false
		for k := 0; k < cantidadVertices; k++ {
			v := vertices[k]
			prNuevo := _arrastrePR(g, v, vertices, cantidadVertices, pageRanks)
			pageRanksActualizados.Guardar(v, prNuevo)
			if difPorcentual(pageRanks.Obtener(v), prNuevo) > 5 {
				seguirIterando = true
			}
		}

		pageRanksActualizados.Iterar(func(clave K, dato float64) bool {
			pageRanks.Guardar(clave, dato)
			return true
		})
	}

	pageRanks.Iterar(func(clave K, dato float64) bool {
		heap.Encolar(verticePR[K]{clave, dato})
		return true
	})

	pageOrdenadas = append(pageOrdenadas, heap.Desencolar().vertice)

	return pageOrdenadas
}

func _arrastrePR[K comparable](g TDAGrafo.GrafoNoPesado[K], vertice K, vertices []K, N int, pageRanks TDADicc.Diccionario[K, float64]) float64 {
	var primerTermino, segundoTermino float64
	primerTermino = (1 - d) / float64(N)

	for j := 0; j < N; j++ {
		w := vertices[j]
		if g.HayArista(w, vertice) {
			cantAdyacentes := len(g.Adyacente(w))
			segundoTermino += pageRanks.Obtener(w) / float64(cantAdyacentes)
		}
	}

	segundoTermino *= d

	return primerTermino + segundoTermino
}

// Lectura 2 AM

func OrdenTopologico[K comparable](g TDAGrafo.GrafoNoPesado[K]) ([]K, error) {
	if !g.Dirigido() || existeCiclo[K](g) {
		return []K{}, ERROR.ErrorNoExisteOrden{}
	}
	res := []K{}
	grados := GradoDeEntrada[K](g)
	q := TDACola.CrearColaEnlazada[K]()
	for i := grados.Iterador(); i.HaySiguiente(); i.Siguiente() {
		vertice, grado := i.VerActual()
		if grado == 0 {
			q.Encolar(vertice)
			res = append(res, vertice)
		}
	}
	for !q.EstaVacia() {
		v := q.Desencolar()
		for _, w := range g.Adyacente(v) {
			grado_anterior := grados.Obtener(w)
			grados.Guardar(w, grado_anterior-1)
			if grado_anterior-1 == 0 {
				res = append(res, w)
				q.Encolar(w)
			}
		}
	}
	return res, nil
}

func existeCiclo[K comparable](g TDAGrafo.Grafo[K]) bool {
	var none K
	padres := TDADicc.CrearHash[K, K]()
	q := TDACola.CrearColaEnlazada[K]()
	vertice_random := g.VerticeAleatorio()
	padres.Guardar(vertice_random, none)
	q.Encolar(vertice_random)
	for !q.EstaVacia() {
		v := q.Desencolar()
		for _, w := range g.Adyacente(v) {
			if padres.Pertenece(w) {
				if padres.Obtener(w) != v {
					return true
				}
			}
			padres.Guardar(w, v)
			q.Encolar(w)
		}
	}
	return false
}

func invertir[K comparable](arr []K) []K {
	res := []K{}
	p := TDAPila.CrearPilaDinamica[K]()
	for _, elemento := range arr {
		p.Apilar(elemento)
	}
	for !p.EstaVacia() {
		res = append(res, p.Desapilar())
	}
	return res
}

// Ciclo N Artículos

func dfs_cicloN[K comparable](g TDAGrafo.Grafo[K], origen, destino K, dist int) []K {
	var none K
	for _, w := range g.Adyacente(origen) {
		padres := TDADicc.CrearHash[K, K]()
		padres.Guardar(origen, none)
		padres.Guardar(w, origen)
		var hayCamino bool
		var camino []K
		_dfs_cicloN_aux[K](g, 1, dist, w, destino, padres, &camino, &hayCamino)
		if hayCamino {
			return camino
		}
	}
	fmt.Print("error")
	return []K{}
}

func _dfs_cicloN_aux[K comparable](g TDAGrafo.Grafo[K], contador, n int, v, destino K, padres TDADicc.Diccionario[K, K], camino *[]K, hayCamino *bool) {
	if contador == n && v == destino {
		cam, _ := ReconstruirCamino[K](padres, destino, padres.Obtener(v))
		cam = append(cam, v)
		*hayCamino = true
		*camino = cam
	} else if contador < n {
		lista := g.Adyacente(v)
		for _, w := range lista {
			padres.Guardar(w, v)
			_dfs_cicloN_aux[K](g, contador+1, n, w, destino, padres, camino, hayCamino)
		}
	}
}
