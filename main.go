package main

import (
	"fmt"
	FUNCIONES "netstats/funciones"
)

func main() {
	//g := FUNCIONES.LeerGrafotxt("wiki-reducido-75000.tsv")
	// comunidades := FUNCIONES.Comunidades[string](g, "Chile")
	// fmt.Println(FUNCIONES.ComprobarChile[string](comunidades))
	dict := FUNCIONES.LeerChile()
	dict.Iterar(func(clave string, dato bool) bool {
		fmt.Println(clave)
		return true
	})
}
