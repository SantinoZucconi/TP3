package funciones

import (
	"strconv"
	"strings"
	INTERNET "tp3/diseno_alumnos"
	ERROR "tp3/errores"
)

const VACIO int = 0

const (
	LISTAR       string = "listar_operaciones"
	CAMINO       string = "camino"
	PAGERANK     string = "mas_importantes"
	CONECTADOS   string = "conectados"
	CICLO_N      string = "ciclo"
	LECTURA      string = "lectura"
	DIAMETRO     string = "diametro"
	RANGO        string = "rango"
	COMUNIDADES  string = "comunidad"
	NAVEGACION_1 string = "navegacion"
	CLUSTERING   string = "clustering"
)

const NULO float64 = -1

func recuperarPagina(pagina []string) string {
	return strings.Join(pagina, " ")
}

func separarPorComas(paginas []string) []string {
	return strings.Split(strings.Join(paginas, " "), ",")
}

func ListarOperaciones(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 1 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	return internet.Operaciones(), nil
}

func EncontrarCaminoMinimo(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) < 2 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	extremos := separarPorComas(entrada[1:])
	if len(extremos) != 2 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	origen := extremos[0]
	destino := extremos[1]
	camino, err := internet.CaminoMasCorto(origen, destino)
	return camino, err
}

func CalcularDiametro(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 1 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	return internet.Diametro(), nil
}

func PaginasEnRango(internet INTERNET.Internet, entrada []string) (float64, error) {
	if len(entrada) < 2 {
		return NULO, &ERROR.ErrorComandoInvalido{}
	}
	calculo := separarPorComas(entrada[1:])
	if len(calculo) > 2 {
		return NULO, &ERROR.ErrorComandoInvalido{}
	}
	origen := calculo[0]
	rango, err := strconv.Atoi(calculo[1])
	if err != nil || rango < 0 {
		return NULO, &ERROR.ErrorComandoInvalido{}
	}
	return float64(internet.EnRango(origen, rango)), nil
}

func NavegarPrimerLink(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) < 2 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}

	origen := recuperarPagina(entrada[1:])
	camino := internet.NavPrimerLink(origen)
	return camino, nil
}

func CalcularClustering(internet INTERNET.Internet, entrada []string) (float64, error) {
	if len(entrada) > 1 {
		pagina := recuperarPagina(entrada[1:])
		return internet.ClusteringIndividual(pagina), nil
	}
	return internet.ClusteringRed(), nil
}

func ListaConectados(internet INTERNET.Internet, entrada []string) ([]string, error) {
	pagina := recuperarPagina(entrada[1:])
	return internet.Conectividad(pagina), nil
}

func PaginasMasImportantes(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 2 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	top, err := strconv.Atoi(entrada[1])
	if err != nil || top < 1 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}

	return internet.MasImportantes(top), nil

}

func Lectura2am(internet INTERNET.Internet, entrada []string) ([]string, error) {
	paginas := separarPorComas(entrada[1:])
	return internet.Lectura2am(paginas)
}

func Comunidades(internet INTERNET.Internet, entrada []string) ([]string, error) {
	pagina := recuperarPagina(entrada[1:])
	return internet.ComunidadPagina(pagina), nil
}

func CicloNesimo(internet INTERNET.Internet, entrada []string) ([]string, error) {
	ciclo := separarPorComas(entrada[1:])
	if len(ciclo) != 2 {
		return []string{}, ERROR.ErrorComandoInvalido{}
	}
	cantidad, err := strconv.Atoi(ciclo[1])
	if err != nil || cantidad < 1 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}

	return internet.CicloPaginas(ciclo[0], cantidad)
}

func ProcesarComando(internet INTERNET.Internet, entrada []string) ([]string, float64, error) {
	var err error
	var valor float64
	var lista []string

	if len(entrada) == VACIO {
		err = &ERROR.ErrorNoHayEntrada{}
		return lista, valor, err
	}
	comando := entrada[0]
	switch comando {
	case LISTAR:
		lista, err = ListarOperaciones(internet, entrada)
	case CAMINO:
		lista, err = EncontrarCaminoMinimo(internet, entrada)
	case PAGERANK:
		lista, err = PaginasMasImportantes(internet, entrada)
	case CONECTADOS:
		lista, err = ListaConectados(internet, entrada)
	case CICLO_N:
		lista, err = CicloNesimo(internet, entrada)
	case LECTURA:
		lista, err = Lectura2am(internet, entrada)
	case DIAMETRO:
		lista, err = CalcularDiametro(internet, entrada)
	case RANGO:
		valor, err = PaginasEnRango(internet, entrada)
	case COMUNIDADES:
		lista, err = Comunidades(internet, entrada)
	case NAVEGACION_1:
		lista, err = NavegarPrimerLink(internet, entrada)
	case CLUSTERING:
		valor, err = CalcularClustering(internet, entrada)
	default:
		err = &ERROR.ErrorComandoInvalido{}
	}
	return lista, valor, err
}
