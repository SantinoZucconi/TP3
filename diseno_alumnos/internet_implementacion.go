package internet

import (
	FSI "tdas/FSI"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	ERROR "tp3/errores"
)

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
	comunidades            TDADicc.Diccionario[int, []string]
	labels                 TDADicc.Diccionario[string, int]
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
	camino, err := FSI.CaminoMinimo(i.grafo, origen, destino)
	return camino, err
}

func (i *internet) Diametro() []string {
	if !i.diametroCalculado {
		i.diametro = FSI.Diametro(i.grafo)
		i.diametroCalculado = true
	}
	return i.diametro
}

func (i *internet) EnRango(pagina string, rango int) int {
	if !i.grafo.EsVertice(pagina) {
		return 0
	}

	return FSI.TodosEnRango(i.grafo, pagina, rango)
}

func (i *internet) NavPrimerLink(pagina string) []string {
	if !i.grafo.EsVertice(pagina) {
		return []string{pagina}
	}
	return FSI.NavegacionPrimerLink(i.grafo, i.primerosLinks, pagina)
}

func (i *internet) Conectividad(pagina string) []string {
	if !i.grafo.EsVertice(pagina) {
		return []string{pagina}
	}

	if !i.cfcsCalculadas {
		i.cfcs, i.pertenencia = FSI.CFC(i.grafo)
		i.cfcsCalculadas = true
	}
	return i.cfcs[i.pertenencia.Obtener(pagina)]

}

func (i *internet) ClusteringRed() float64 {

	if !i.clusteringRedCalculado {
		i.clusteringRed = FSI.ClusteringRed(i.grafo)
		i.clusteringRedCalculado = true
	}

	return i.clusteringRed
}

func (i *internet) ClusteringIndividual(pagina string) float64 {
	if !i.grafo.EsVertice(pagina) {
		return 0
	}
	return FSI.ClusteringIndividual(i.grafo, pagina)
}

func (i *internet) MasImportantes(top int) []string {
	if !i.pageRankCalculado {
		i.pageRank = FSI.PageRank(i.grafo)
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
	return FSI.OrdenLectura(i.grafo, paginas)
}

func (i *internet) ComunidadPagina(pagina string) []string {
	if i.comunidadesCalculado {
		if !i.labels.Pertenece(pagina) {
			return []string{pagina}
		}
		return i.comunidades.Obtener(i.labels.Obtener(pagina))
	}
	i.comunidadesCalculado = true
	i.labels, i.comunidades = FSI.Comunidades(i.grafo)
	return i.comunidades.Obtener(i.labels.Obtener(pagina))
}

func (i *internet) CicloPaginas(pagina string, n int) ([]string, error) {
	if !i.grafo.EsVertice(pagina) {
		return []string{}, &ERROR.ErrorNoExisteRecorrido{}
	}
	return FSI.CicloN(i.grafo, pagina, n)
}
