package funciones

import "fmt"

func ImprimirLista(lista []string) {
	for i := 0; i < len(lista); i++ {
		fmt.Printf("\n%s", lista[i])
	}
}

func ImprimirCamino(camino []string, imprimirCosto bool) {
	largo := len(camino)
	fmt.Printf("\n%s", camino[0])
	for i := 1; i < largo; i++ {
		fmt.Printf(" -> %s", camino[i])
	}

	if imprimirCosto {
		fmt.Printf("\nCosto: %d", largo)
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
