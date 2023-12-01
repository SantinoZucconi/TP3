package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	ERROR "tp3/errores"
	FUNCIONES "tp3/funciones"
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
	NAVEGACION_1 string = "navegación"
	CLUSTERING   string = "clustering"
)

func main() {

	argumentos := os.Args
	links := argumentos[1:]
	var (
		err                    error
		lista                  []string
		valor                  float64
		pageRanks              []string
		pRcalculado            bool
		diametro               []string
		diametroCalculado      bool
		clusteringRed          float64
		clusteringRedCalculado bool
	)

	operaciones := []string{"camino", "mas_importantes", "conectados", "ciclo", "lectura", "diametro", "rango", "comunidad", "navegacion", "clustering"}

	if len(links) > 1 {
		err = &ERROR.ErrorParametros{}
		fmt.Printf("%s\n", err.Error())
		return
	}

	internet := FUNCIONES.LeerGrafotxt(links[0])

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		entrada := strings.Fields(s.Text())
		if len(entrada) == VACIO {
			err = &ERROR.ErrorNoHayEntrada{}
		} else {
			comando := entrada[0]
			switch comando {
			case LISTAR:
				if len(entrada) != 1 {
					err = &ERROR.ErrorComandoInvalido{}
				}
				lista = operaciones
			case CAMINO:
				if len(entrada) != 2 {
					err = &ERROR.ErrorComandoInvalido{}
				} else {
					extremos := strings.Split(entrada[1], ",")
					origen := extremos[0]
					destino := extremos[1]
					lista, err = FUNCIONES.CaminoMinimo(internet, origen, destino)
				}
			case PAGERANK:
				if len(entrada) != 2 {
					err = &ERROR.ErrorComandoInvalido{}
				}
				if !pRcalculado {
					pageRanks = FUNCIONES.PageRank(internet)
					pRcalculado = true
				}
				lista = pageRanks
				pr, err := strconv.Atoi(entrada[1])
				valor = float64(pr)
				if err != nil {
					err = &ERROR.ErrorComandoInvalido{}
				}
			case CONECTADOS:
				//
			case CICLO_N:
				//
			case LECTURA:
				//
			case DIAMETRO:
				if len(entrada) != 1 {
					err = &ERROR.ErrorComandoInvalido{}
				} else if !diametroCalculado {
					diametro = FUNCIONES.Diametro(internet)
				}
				lista = diametro

			case RANGO:
				if len(entrada) != 2 {
					err = &ERROR.ErrorComandoInvalido{}
				}
				calculo := strings.Split(entrada[1], ",")
				origen := calculo[0]
				rango, err := strconv.Atoi(calculo[1])
				if err != nil {
					err = &ERROR.ErrorComandoInvalido{}
				} else {
					valor = float64(FUNCIONES.TodosEnRango(internet, origen, rango))
				}

			case COMUNIDADES:
				//
			case NAVEGACION_1:
				//
			case CLUSTERING:
				if len(entrada) > 1 {
					pagina := strings.Join(entrada[1:], " ")
					valor = FUNCIONES.CLUSTERING_INDIVIDUAL(internet, pagina)
				} else if !clusteringRedCalculado {
					clusteringRed = FUNCIONES.CLUSTERING_RED(internet)
				} else {
					valor = clusteringRed
				}

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
		FUNCIONES.ImprimirOperaciones(lista)
	case CAMINO:
		FUNCIONES.ImprimirCamino(lista, true)
	case PAGERANK:
		FUNCIONES.ImprimirMasImportantes(lista, int(valor))
	case CONECTADOS:
		//
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
		//
	case CLUSTERING:
		FUNCIONES.ImprimirValor(valor)
	}
}
