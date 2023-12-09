package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDALista "tdas/lista"
	INTERNET "tp3/diseno"
	ERROR "tp3/errores"
	FUNCIONES "tp3/funciones"
)

func main() {

	argumentos := os.Args
	links := argumentos[1:]
	var (
		err   error
		lista TDALista.Lista[string]
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

		lista, valor, err = FUNCIONES.ProcesarComando(internet, entrada)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		} else {
			FUNCIONES.ImprimirResultado(entrada[0], lista, valor)
		}
	}
}
