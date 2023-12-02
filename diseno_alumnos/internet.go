package internet

type Internet interface {
	Operaciones() []string

	CaminoMasCorto(string, string) ([]string, error)
	Diametro() []string
	EnRango(string, int) int
	NavPrimerLink(string) []string

	Conectividad(string) []string
	Comunidades()
	OrdenPaginas() []string
	ClusteringRed() float64
	ClusteringIndividual(string) float64

	MasImportantes(int) []string
	CicloArticulos() []string
}

/*
Camino más corto (★) LISTO                               // devuelve []K, error
Artículos más importantes (★★★) LISTO?					// devuelve []K
Conectividad (★★) LISTO?								// devuelve [][]K, Dic
Ciclo de n artículos (★★★)								// devuelve []K
Lectura a las 2 a.m. (★★)								// devuelve []K
Diametro (★) LISTO										// devuelve []K
Todos en Rango (★) LISTO								// devuelve int
Comunidades (★★)										// ?
Navegacion por primer link (★) LISTO?					// devuelve []K
Coeficiente de Clustering (★★)	LISTO					// devuelve float

*/
