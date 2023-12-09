package internet

import TDALista "tdas/lista"

type Internet interface {
	Operaciones() TDALista.Lista[string]

	CaminoMasCorto(string, string) (TDALista.Lista[string], error)
	Diametro() TDALista.Lista[string]
	EnRango(string, int) int
	NavPrimerLink(string) TDALista.Lista[string]

	Conectividad(string) TDALista.Lista[string]
	ComunidadPagina(string) TDALista.Lista[string]
	Lectura2am([]string) (TDALista.Lista[string], error)
	ClusteringRed() float64
	ClusteringIndividual(string) float64

	MasImportantes(int) TDALista.Lista[string]
	CicloPaginas(string, int) (TDALista.Lista[string], error)
}
