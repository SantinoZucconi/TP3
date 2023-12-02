package internet

import (
	"bufio"
	"os"
	"strings"
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func LeerGrafotxt(archivo string) (TDAGrafo.GrafoNoPesado[string], TDADicc.Diccionario[string, string]) {
	g := TDAGrafo.CrearGrafoNoPesado[string](true)
	dicc := TDADicc.CrearHash[string, TDACola.Cola[string]]()
	primerosLinks := TDADicc.CrearHash[string, string]()
	datos, err := os.Open(archivo)
	if err != nil {
		panic("No se pudo abrir el archivo")
	}

	defer datos.Close()
	scaner := bufio.NewScanner(datos)
	for scaner.Scan() {
		texto := scaner.Text()
		vertices := strings.Split(texto, "\t")
		g.AgregarVertice(vertices[0])
		primerosLinks.Guardar(vertices[0], vertices[1])
		dicc.Guardar(vertices[0], TDACola.CrearColaEnlazada[string]())
		for i := 1; i < len(vertices); i++ {
			dicc.Obtener(vertices[0]).Encolar(vertices[i])
		}
	}

	dicc.Iterar(func(v string, adyacentes TDACola.Cola[string]) bool {
		for !adyacentes.EstaVacia() {
			g.AgregarAristaNP(v, adyacentes.Desencolar())
		}
		return true
	})

	return g, primerosLinks
}

func LeerChile() TDADicc.Diccionario[string, bool] {
	dicc := TDADicc.CrearHash[string, bool]()
	datos, err := os.Open("comunidad-chile.txt")
	if err != nil {
		panic("No se pudo abrir el archivo")
	}
	defer datos.Close()
	scaner := bufio.NewScanner(datos)
	for scaner.Scan() {
		texto := scaner.Text()
		vertices := strings.Split(texto, ", ")
		for _, v := range vertices {
			dicc.Guardar(v, true)
		}
	}
	return dicc
}

func ComprobarChile[K comparable](arr []string) []string {
	comChile := LeerChile()
	faltantes := []string{}
	for _, v := range arr {
		if !comChile.Pertenece(v) {
			faltantes = append(faltantes, v)
		}
	}
	return faltantes
}
