package internet

import (
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	ERROR "tp3/errores"
)

type comunidades[K comparable] struct {
	label     TDADicc.Diccionario[K, int]
	comunidad TDADicc.Diccionario[int, []K]
}

type internet struct {
	grafo                  TDAGrafo.GrafoNoPesado[string]
	operaciones            []string
	primerosLinks          TDADicc.Diccionario[string, string]
	pageRank               []string
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
	if !i.grafo.EsVertice(origen) || !i.grafo.EsVertice(destino) {
		return []string{}, &ERROR.ErrorNoExisteRecorrido{}
	}
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
	if !i.grafo.EsVertice(pagina) {
		return 0
	}

	return TodosEnRango(i.grafo, pagina, rango)
}

func (i *internet) NavPrimerLink(pagina string) []string {
	if !i.grafo.EsVertice(pagina) {
		return []string{pagina}
	}
	return NavegacionPrimerLink(i.grafo, i.primerosLinks, pagina)
}

func (i *internet) Conectividad(pagina string) []string {
	if !i.grafo.EsVertice(pagina) {
		return []string{pagina}
	}

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
	if !i.grafo.EsVertice(pagina) {
		return 0
	}
	return ClusteringIndividual(i.grafo, pagina)
}

func (i *internet) MasImportantes(top int) []string {
	if !i.pageRankCalculado {
		i.pageRank = PageRank(i.grafo)
		i.pageRankCalculado = true
	}

	return i.pageRank[:top]
}

func (i *internet) Lectura2am(paginas []string) ([]string, error) {
	for _, pagina := range paginas {
		if !i.grafo.EsVertice(pagina) {
			return []string{}, &ERROR.ErrorNoExisteOrden{}
		}
	}
	return _Lectura2am(i.grafo, paginas)
}

func (i *internet) Comunidades(pagina string) []string {
	if i.comunidadesCalculado {
		if !i.comunidades.label.Pertenece(pagina) {
			return []string{pagina}
		}
		return i.comunidades.comunidad.Obtener(i.comunidades.label.Obtener(pagina))
	}
	i.comunidadesCalculado = true
	i.comunidades = _Comunidades[string](i.grafo)
	return i.comunidades.comunidad.Obtener(i.comunidades.label.Obtener(pagina))
}

func (i *internet) CicloN(pagina string, n int) ([]string, error) {
	if !i.grafo.EsVertice(pagina) {
		return []string{}, &ERROR.ErrorNoExisteRecorrido{}
	}
	return _CicloN[string](i.grafo, pagina, n)
}
