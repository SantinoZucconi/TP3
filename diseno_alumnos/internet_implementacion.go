package internet

import (
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

type paginaPR struct {
	pagina   string
	pageRank float64
}

type comunidades[K comparable] struct {
	label     TDADicc.Diccionario[K, int]
	comunidad TDADicc.Diccionario[int, []K]
}

type internet struct {
	grafo                  TDAGrafo.GrafoNoPesado[string]
	operaciones            []string
	primerosLinks          TDADicc.Diccionario[string, string]
	pageRank               TDADicc.Diccionario[string, float64] // o pageRank []string
	pageRankCalculado      bool
	cfcs                   [][]string
	pertenencia            TDADicc.Diccionario[string, int]
	cfcsCalculadas         bool
	clusteringRed          float64
	clusteringRedCalculado bool
	diametro               []string
	diametroCalculado      bool
	comunidades            comunidades[string]
	comunidadesCalculado   bool
}

func GenerarInternet(archivo string) Internet {
	operaciones := []string{"camino", "mas_importantes", "conectados", "ciclo", "lectura", "diametro", "rango", "comunidad", "navegacion", "clustering"}
	grafo, primerosLinks := LeerGrafotxt(archivo)
	return &internet{grafo: grafo, operaciones: operaciones, primerosLinks: primerosLinks}
}

func (i *internet) Operaciones() []string {
	return i.operaciones
}

func (i *internet) CaminoMasCorto(origen, destino string) ([]string, error) {
	camino, err := CaminoMinimo(i.grafo, origen, destino)
	return camino, err
}

func (i *internet) Diametro() []string {
	if !i.diametroCalculado {
		i.diametro = Diametro(i.grafo)
		i.diametroCalculado = true
	}
	return i.diametro
}

func (i *internet) EnRango(pagina string, rango int) int {
	return TodosEnRango(i.grafo, pagina, rango)
}

func (i *internet) NavPrimerLink(pagina string) []string {
	v := pagina
	contador := 0
	hayLinks := true
	camino := []string{v}
	contador++
	for contador < 20 && hayLinks {
		if i.primerosLinks.Pertenece(v) {
			v = i.primerosLinks.Obtener(v)
			camino = append(camino, v)
			contador++
		} else {
			hayLinks = false
		}
	}

	return camino
}

func (i *internet) Conectividad(pagina string) []string {
	if !i.cfcsCalculadas {
		i.cfcs, i.pertenencia = CFC(i.grafo)
		i.cfcsCalculadas = true
	}

	return i.cfcs[i.pertenencia.Obtener(pagina)]

}

func (i *internet) ClusteringRed() float64 {
	if !i.clusteringRedCalculado {
		i.clusteringRed = ClusteringRed(i.grafo)
		i.clusteringRedCalculado = true
	}

	return i.clusteringRed
}

func (i *internet) ClusteringIndividual(pagina string) float64 {
	return ClusteringIndividual(i.grafo, pagina)
}

func (i *internet) MasImportantes(top int) []string {
	// definir que hacer (ordenar o top K a partir de diccionario?)
	return []string{}
}

func (i *internet) Lectura2am(paginas []string) ([]string, error) {
	dicc := TDADicc.CrearHash[string, bool]()
	subgrafo := TDAGrafo.CrearGrafoNoPesado[string](true)
	for _, pagina := range paginas {
		dicc.Guardar(pagina, true)
		subgrafo.AgregarVertice(pagina)
	}
	for _, pagina := range paginas {
		lista := i.grafo.Adyacente(pagina)
		for _, link := range lista {
			if dicc.Pertenece(link) {
				subgrafo.AgregarAristaNP(pagina, link)
			}
		}
	}
	recorrido, err := OrdenTopologico[string](subgrafo)
	return invertir[string](recorrido), err
}

func (i *internet) Comunidades(s string) []string {
	if i.comunidadesCalculado {
		if !i.comunidades.label.Pertenece(s) {
			return []string{}
		}
		return i.comunidades.comunidad.Obtener(i.comunidades.label.Obtener(s))
	}
	label := TDADicc.CrearHash[string, int]()
	_comunidades := TDADicc.CrearHash[int, []string]()
	vEntradas := vertices_entrada[string](i.grafo)
	for i, v := range i.grafo.ObtenerVertices() {
		label.Guardar(v, i)
	}
	for j := 0; j < 4; j++ {
		for _, v := range i.grafo.ObtenerVertices() {
			labelNeighbor := []int{}
			for _, w := range i.grafo.Adyacente(v) {
				labelNeighbor = append(labelNeighbor, label.Obtener(w))
			}
			if i.grafo.Dirigido() {
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
	label.Iterar(func(clave string, dato int) bool {
		if !_comunidades.Pertenece(dato) {
			_comunidades.Guardar(dato, []string{clave})
		} else {
			comunidad := _comunidades.Obtener(dato)
			comunidad = append(comunidad, clave)
			_comunidades.Guardar(dato, comunidad)
		}
		return true
	})
	i.comunidades = comunidades[string]{label: label, comunidad: _comunidades}
	i.comunidadesCalculado = true
	return i.comunidades.comunidad.Obtener(i.comunidades.label.Obtener(s))
}

func (i *internet) CicloN(p string, n int) []string {
	distancia := TDADicc.CrearHash[string, int]()
	visitados := TDADicc.CrearHash[string, bool]()
	distancia.Guardar(p, 0)
	subgrafo := TDAGrafo.CrearGrafoNoPesado[string](true)
	q := TDACola.CrearColaEnlazada[string]()
	q.Encolar(p)
	for !q.EstaVacia() {
		v := q.Desencolar()
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, true)
			subgrafo.AgregarVertice(v)
		}
		for _, w := range i.grafo.Adyacente(v) {
			distancia_w := distancia.Obtener(v) + 1
			if !visitados.Pertenece(w) && distancia_w < n {
				visitados.Guardar(w, true)
				subgrafo.AgregarVertice(w)
			}
			if distancia_w == n && w == p || distancia_w < n {
				subgrafo.AgregarAristaNP(v, w)
			}
			if distancia_w <= n {
				distancia.Guardar(w, distancia_w)
				q.Encolar(w)
			}
		}
	}
	return dfs_cicloN[string](subgrafo, p, p, n)
}
