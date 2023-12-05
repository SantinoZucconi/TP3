package funciones

import (
	"fmt"
)

func ImprimirLista(lista []string) {
	for i := 0; i < len(lista); i++ {
		fmt.Printf("\n%s", lista[i])
	}
}

func ImprimirPaginas(camino []string, separador string, imprimirCosto bool) {
	largo := len(camino)
	fmt.Printf("\n%s", camino[0])
	for i := 1; i < largo; i++ {
		fmt.Printf(separador+" %s", camino[i])
	}

	if imprimirCosto {
		fmt.Printf("\nCosto: %d", largo-1)
	}

}

func ImprimirValor(cant float64) {
	fmt.Printf("\n%f", cant)
}

func ImprimirResultado(comando string, lista []string, valor float64) {
	switch comando {
	case LISTAR:
		ImprimirLista(lista)
	case CAMINO:
		ImprimirPaginas(lista, " ->", true)
	case PAGERANK:
		ImprimirPaginas(lista, ",", false)
	case CONECTADOS:
		ImprimirLista(lista)
	case CICLO_N:
		ImprimirPaginas(lista, " ->", false)
	case LECTURA:
		ImprimirPaginas(lista, ",", false)
	case DIAMETRO:
		ImprimirPaginas(lista, " ->", true)
	case RANGO:
		ImprimirValor(valor)
	case COMUNIDADES:
		ImprimirPaginas(lista, ",", false)
	case NAVEGACION_1:
		ImprimirPaginas(lista, " ->", false)
	case CLUSTERING:
		ImprimirValor(valor)
	}
}
