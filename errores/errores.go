package errores

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "Error: Cantidad de archivos invalida"
}

type ErrorNoHayEntrada struct{}

func (e ErrorNoHayEntrada) Error() string {
	return "Error: no hay entrada para realizar"
}

type ErrorComandoInvalido struct{}

func (e ErrorComandoInvalido) Error() string {
	return "Error: el comando es invalido"
}

type ErrorNoExisteRecorrido struct{}

func (e ErrorNoExisteRecorrido) Error() string {
	return "No se encontro recorrido"
}

type ErrorNoExisteOrden struct{}

func (e ErrorNoExisteOrden) Error() string {
	return "No existe forma de leer las paginas en orden"
}
