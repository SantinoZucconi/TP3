package internet

import (
	"bufio"
	"os"
	"strings"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

func LeerGrafotxt(archivo string) (TDAGrafo.GrafoNoPesado[string], TDADicc.Diccionario[string, string]) {
	g := TDAGrafo.CrearGrafoNoPesado[string](true)
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
		if !g.EsVertice(vertices[0]) {
			g.AgregarVertice(vertices[0])
		}
		if len(vertices) > 1 {
			primerosLinks.Guardar(vertices[0], vertices[1])
		}
		for i := 1; i < len(vertices); i++ {
			if !g.EsVertice(vertices[i]) {
				g.AgregarVertice(vertices[i])
			}
			g.AgregarAristaNP(vertices[0], vertices[i])
		}
	}

	return g, primerosLinks
}
