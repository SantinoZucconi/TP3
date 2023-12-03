package funciones

import (
	"strconv"
	"strings"
	INTERNET "tp3/diseno_alumnos"
	ERROR "tp3/errores"
)

const NULO float64 = -1

func ListarOperaciones(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 1 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	return internet.Operaciones(), nil
}

func EncontrarCaminoMinimo(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 2 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	extremos := strings.Split(entrada[1], ",")
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
	if len(entrada) != 2 {
		return NULO, &ERROR.ErrorComandoInvalido{}
	}
	calculo := strings.Split(entrada[1], ",")
	origen := calculo[0]
	rango, err := strconv.Atoi(calculo[1])
	if err != nil {
		return NULO, &ERROR.ErrorComandoInvalido{}
	}
	return float64(internet.EnRango(origen, rango)), nil

}

func NavegarPrimerLink(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 2 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	origen := entrada[1]
	camino := internet.NavPrimerLink(origen)
	return camino, nil
}

func CalcularClustering(internet INTERNET.Internet, entrada []string) (float64, error) {
	if len(entrada) > 1 {
		pagina := strings.Join(entrada[1:], " ")
		return internet.ClusteringIndividual(pagina), nil
	}
	return internet.ClusteringRed(), nil
}

func ListaConectados(internet INTERNET.Internet, entrada []string) ([]string, error) {
	pagina := strings.Join(entrada[1:], " ")
	return internet.Conectividad(pagina), nil
}

/*

func --- (internet INTERNET.Internet, entrada []string) ([]string, error) {

}

func --- (internet INTERNET.Internet, entrada []string) ([]string, error) {

}
*/

func PaginasMasImportantes(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 2 {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}
	top, err := strconv.Atoi(entrada[1])
	if err != nil {
		return []string{}, &ERROR.ErrorComandoInvalido{}
	}

	return internet.MasImportantes(top), nil

}

func Lectura2am(internet INTERNET.Internet, entrada []string) ([]string, error) {
	return internet.Lectura2am(entrada)
}

func Comunidades(internet INTERNET.Internet, pagina []string) ([]string, error) {
	if len(pagina) != 1 {
		return []string{}, ERROR.ErrorComandoInvalido{}
	}
	return internet.Comunidades(pagina[0]), nil
}

func CicloNesimo(internet INTERNET.Internet, entrada []string) ([]string, error) {
	if len(entrada) != 2 {
		return []string{}, ERROR.ErrorComandoInvalido{}
	}
	cantidad, err := strconv.Atoi(entrada[1])
	return internet.CicloN(entrada[0], cantidad), err
}
