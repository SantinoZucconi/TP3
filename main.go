package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	INTERNET "tp3/diseno_alumnos"
	ERROR "tp3/errores"
	FUNCIONES "tp3/funciones"
)

const VACIO int = 0

// DUDA: QUE HACER SI LA PAGINA NO EXISTE?

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
	NAVEGACION_1 string = "navegación"
	CLUSTERING   string = "clustering"
)

func main() {

	argumentos := os.Args
	links := argumentos[1:]
	var (
		err   error
		lista []string
		valor float64
	)

	if len(links) > 1 {
		err = &ERROR.ErrorParametros{}
		fmt.Printf("%s\n", err.Error())
		return
	}

	internet := INTERNET.GenerarInternet(links[0])

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		entrada := strings.Fields(s.Text())
		if len(entrada) == VACIO {
			err = &ERROR.ErrorNoHayEntrada{}
		} else {
			comando := entrada[0]
			switch comando {
			case LISTAR:
				lista, err = FUNCIONES.ListarOperaciones(internet, entrada)
			case CAMINO:
				lista, err = FUNCIONES.EncontrarCaminoMinimo(internet, entrada)
			case PAGERANK:
				lista, err = FUNCIONES.PaginasMasImportantes(internet, entrada)
			case CONECTADOS:
				//
			case CICLO_N:
				//
			case LECTURA:
				//
			case DIAMETRO:
				lista, err = FUNCIONES.CalcularDiametro(internet, entrada)
			case RANGO:
				valor, err = FUNCIONES.PaginasEnRango(internet, entrada)
			case COMUNIDADES:
				//
			case NAVEGACION_1:
				lista, err = FUNCIONES.NavegarPrimerLink(internet, entrada)
			case CLUSTERING:
				valor, err = FUNCIONES.CalcularClustering(internet, entrada)
			default:
				err = &ERROR.ErrorComandoInvalido{}
			}

			if err != nil {
				fmt.Printf("%s\n", err.Error())
			} else {
				imprimir_resultado(comando, lista, valor)
			}
		}
	}

}

func imprimir_resultado(comando string, lista []string, valor float64) {
	switch comando {
	case LISTAR:
		FUNCIONES.ImprimirLista(lista)
	case CAMINO:
		FUNCIONES.ImprimirCamino(lista, true)
	case PAGERANK:
		FUNCIONES.ImprimirMasImportantes(lista, int(valor))
	case CONECTADOS:
		FUNCIONES.ImprimirLista(lista)
	case CICLO_N:
		//
	case LECTURA:
		//
	case DIAMETRO:
		FUNCIONES.ImprimirCamino(lista, true)
	case RANGO:
		FUNCIONES.ImprimirValor(valor)
	case COMUNIDADES:
		//
	case NAVEGACION_1:
		FUNCIONES.ImprimirCamino(lista, false)
	case CLUSTERING:
		FUNCIONES.ImprimirValor(valor)
	}
}
