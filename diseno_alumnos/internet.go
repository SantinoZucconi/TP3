package internet

type Internet interface {
	Operaciones() []string

	CaminoMasCorto(string, string) ([]string, error)
	Diametro() []string
	EnRango(string, int) int
	NavPrimerLink(string) []string

	Conectividad(string) []string
	Comunidades(string) []string
	Lectura2am([]string) ([]string, error)
	ClusteringRed() float64
	ClusteringIndividual(string) float64

	MasImportantes(int) []string
	CicloN(string, int) ([]string, error)
}

/*
Camino más corto (★) LISTO                               // devuelve []K, error
Artículos más importantes (★★★) LISTO					// devuelve []K
Conectividad (★★) LISTO								// devuelve [][]K, Dic
Ciclo de n artículos (★★★) 								// devuelve []K
Lectura a las 2 a.m. (★★) LISTO								// devuelve []K
Diametro (★) LISTO										// devuelve []K
Todos en Rango (★) LISTO								// devuelve int
Comunidades (★★)										// ?
Navegacion por primer link (★) LISTO					// devuelve []K
Coeficiente de Clustering (★★)	NO ME DA PERON					// devuelve float

*/
