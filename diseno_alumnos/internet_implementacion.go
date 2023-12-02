package internet

import (
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

type paginaPR struct {
	pagina   string
	pageRank float64
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
	return NavegacionPrimerLink(i.grafo, i.primerosLinks, pagina)
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
}
