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

// DUDA: QUE HACER SI LA PAGINA NO EXISTE? dejas la carrera!!

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

	fmt.Print("\nBienvenido a NetStats\n\n")

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		entrada := strings.Fields(s.Text())

		lista, valor, err = FUNCIONES.ProcesarComando(internet, entrada)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		} else {
			FUNCIONES.ImprimirResultado(entrada[0], lista, valor)
		}
		fmt.Print("\n\n")
	}
}
