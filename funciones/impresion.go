package funciones

import (
	"fmt"
)

func ImprimirLista(lista []string) {
	for i := 0; i < len(lista); i++ {
		fmt.Printf("\n%s", lista[i])
	}
}

func ImprimirCamino(camino []string, imprimirCosto bool) {
	largo := len(camino)
	for i := 0; i < largo-1; i++ {
		fmt.Printf("%s -> ", camino[i])
	}
	fmt.Println(camino[largo-1])
	if imprimirCosto {
		fmt.Println("Costo:", largo)
	}

}

func ImprimirMasImportantes(paginas []string, top int) {
	fmt.Printf("\n%s", paginas[0])
	for i := 1; i < top; i++ {
		fmt.Printf(", %s", paginas[i])
	}
}

func ImprimirValor(cant float64) {
	fmt.Printf("\n%f", cant)
}

func ImprimirLectura2am(paginas []string) {
	fmt.Println("Orden:")
	for indice, pagina := range paginas {
		fmt.Println(indice+1, "-", pagina)
	}
}

func ImprimirComunidades(comunidades []string) {
	for _, pagina := range comunidades {
		fmt.Printf("%s, ", pagina)
	}
}

func ImprimirCicloN(ciclo []string) {
	for indice, pagina := range ciclo {
		if indice != len(ciclo)-1 {
			fmt.Printf("%s -> ", pagina)
		} else {
			fmt.Println(pagina)
		}
	}
}
