package funciones

import (
	"fmt"
	TDALista "tdas/lista"
)

func ImprimirLista(lista TDALista.Lista[string]) {
	lista.Iterar(func(s string) bool {
		fmt.Printf("%s\n", s)
		return true
	})
}

func ImprimirPaginas(camino TDALista.Lista[string], separador string, imprimirCosto bool) {
	var contador int
	camino.Iterar(func(s string) bool {
		if contador == 0 {
			fmt.Printf("%s", s)
			contador++
		} else {
			fmt.Printf(separador+" %s", s)
		}
		return true
	})
	if imprimirCosto {
		fmt.Printf("\nCosto: %d", camino.Largo()-1)
	}
	fmt.Print("\n")
}

func ImprimirValor(cant float64, entero bool) {
	if entero {
		ent := int(cant)
		fmt.Printf("%d\n", ent)
		return
	}

	fmt.Printf("%s\n", fmt.Sprintf("%.3f", cant))
}

func ImprimirResultado(comando string, lista TDALista.Lista[string], valor float64) {
	switch comando {
	case LISTAR:
		ImprimirLista(lista)
	case CAMINO:
		ImprimirPaginas(lista, " ->", true)
	case PAGERANK:
		ImprimirPaginas(lista, ",", false)
	case CONECTADOS:
		ImprimirPaginas(lista, ",", false)
	case CICLO_N:
		ImprimirPaginas(lista, " ->", false)
	case LECTURA:
		ImprimirPaginas(lista, ",", false)
	case DIAMETRO:
		ImprimirPaginas(lista, " ->", true)
	case RANGO:
		ImprimirValor(valor, true)
	case COMUNIDADES:
		ImprimirPaginas(lista, ",", false)
	case NAVEGACION_1:
		ImprimirPaginas(lista, " ->", false)
	case CLUSTERING:
		ImprimirValor(valor, false)
	}
}
