package internet

import (
	FSI "tdas/FSI"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDALista "tdas/lista"
	ERROR "tp3/errores"
)

type internet struct {
	grafo                  TDAGrafo.GrafoNoPesado[string]
	operaciones            TDALista.Lista[string]
	primerosLinks          TDADicc.Diccionario[string, string]
	pageRank               TDALista.Lista[string]
	pageRankCalculado      bool
	cfcs                   []TDALista.Lista[string]
	pertenencia            TDADicc.Diccionario[string, int]
	cfcsCalculadas         bool
	clusteringRed          float64
	clusteringRedCalculado bool
	diametro               TDALista.Lista[string]
	diametroCalculado      bool
	comunidades            TDADicc.Diccionario[int, TDALista.Lista[string]]
	labels                 TDADicc.Diccionario[string, int]
	comunidadesCalculado   bool
}

func GenerarInternet(archivo string) Internet {
	comandos := []string{"camino", "mas_importantes", "conectados", "ciclo", "lectura", "diametro", "rango", "comunidad", "navegacion", "clustering"}
	operaciones := TDALista.CrearListaEnlazada[string]()
	for _, operacion := range comandos {
		operaciones.InsertarUltimo(operacion)
	}
	grafo, primerosLinks := LeerGrafotxt(archivo)
	return &internet{grafo: grafo, operaciones: operaciones, primerosLinks: primerosLinks}
}

func (i *internet) Operaciones() TDALista.Lista[string] {
	return i.operaciones
}

func (i *internet) CaminoMasCorto(origen, destino string) (TDALista.Lista[string], error) {
	if !i.grafo.EsVertice(origen) || !i.grafo.EsVertice(destino) {
		return TDALista.CrearListaEnlazada[string](), &ERROR.ErrorNoExisteRecorrido{}
	}
	return FSI.CaminoMinimo(i.grafo, origen, destino)
}

func (i *internet) Diametro() TDALista.Lista[string] {
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

func (i *internet) NavPrimerLink(pagina string) TDALista.Lista[string] {
	if !i.grafo.EsVertice(pagina) {
		lista := TDALista.CrearListaEnlazada[string]()
		lista.InsertarPrimero(pagina)
		return lista
	}
	return FSI.NavegacionPrimerLink(i.grafo, i.primerosLinks, pagina)
}

func (i *internet) Conectividad(pagina string) TDALista.Lista[string] {
	if !i.grafo.EsVertice(pagina) {
		lista := TDALista.CrearListaEnlazada[string]()
		lista.InsertarPrimero(pagina)
		return lista
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

func (i *internet) MasImportantes(top int) TDALista.Lista[string] {
	if !i.pageRankCalculado {
		i.pageRank = FSI.PageRank(i.grafo)
		i.pageRankCalculado = true
	}
	res := TDALista.CrearListaEnlazada[string]()
	var contador int
	i.pageRank.Iterar(func(s string) bool {
		if contador == top {
			return false
		}
		res.InsertarUltimo(s)
		contador++
		return true
	})
	return res
}

func (i *internet) Lectura2am(paginas []string) (TDALista.Lista[string], error) {
	for _, pagina := range paginas {
		if !i.grafo.EsVertice(pagina) {
			return TDALista.CrearListaEnlazada[string](), &ERROR.ErrorNoExisteOrden{}
		}
	}
	return FSI.OrdenLectura(i.grafo, paginas)
}

func (i *internet) ComunidadPagina(pagina string) TDALista.Lista[string] {
	if i.comunidadesCalculado {
		if !i.labels.Pertenece(pagina) {
			lista := TDALista.CrearListaEnlazada[string]()
			lista.InsertarPrimero(pagina)
			return lista
		}
		return i.comunidades.Obtener(i.labels.Obtener(pagina))
	}
	i.comunidadesCalculado = true
	i.labels, i.comunidades = FSI.Comunidades(i.grafo)
	return i.comunidades.Obtener(i.labels.Obtener(pagina))
}

func (i *internet) CicloPaginas(pagina string, n int) (TDALista.Lista[string], error) {
	if !i.grafo.EsVertice(pagina) {
		return TDALista.CrearListaEnlazada[string](), &ERROR.ErrorNoExisteRecorrido{}
	}
	return FSI.CicloN(i.grafo, pagina, n)
}
