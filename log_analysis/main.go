package main

import (
	"bufio"
	ADTMap "data_structures/map"
	"fmt"
	commands "log_analysis/commands"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	_ADD_FILE                = "agregar_archivo"
	_VIEW_VISITORS           = "ver_visitantes"
	_VIEW_MOST_VISITED       = "ver_mas_visitados"
	_OK                      = "OK"
	_ERROR_ADD_FILE          = "Error en comando agregar_archivo"
	_ERROR_VIEW_VISITORS     = "Error en comando ver_visitantes"
	_ERROR_VIEW_MOST_VISITED = "Error en comando ver_mas_visitados"
)

// =============== EJECUTAR COMANDOS ===============

// ejecutarVerMasVisitados ejecuta el comando ver_mas_visitados
// Imprime los n sitios más visitados
// Si la cantidad de argumentos no es la esperada, imprime un mensaje de error y finaliza la ejecución
// Si hay un error al ejecutar el comando, imprime un mensaje de error y finaliza la ejecución
// Si no hay errores, imprime "OK"
func ejecutarVerMasVisitados(args []string, analyzer commands.Analyzer) {
	chequeoArgs(args, 2, _ERROR_VIEW_MOST_VISITED)
	n, err := strconv.Atoi(args[1])
	chequeoError(err, _ERROR_VIEW_MOST_VISITED)
	err = analyzer.VerMasVisitados(n)
	chequeoError(err, _ERROR_VIEW_MOST_VISITED)
	fmt.Println(_OK)
}

// ejecutarVerVisitantes ejecuta el comando ver_visitantes
// Imprime todas las IPs que solicitaron algun recurso en el servidor, dentro del rango de IPs pasado por parametro
// Si la cantidad de argumentos no es la esperada, imprime un mensaje de error y finaliza la ejecución
// Si hay un error al ejecutar el comando, imprime un mensaje de error y finaliza la ejecución
// Si no hay errores, imprime "OK"
func ejecutarVerVisitantes(args []string, analyzer commands.Analyzer) {
	chequeoArgs(args, 3, _ERROR_VIEW_VISITORS)
	err := analyzer.VerVisitantes(args[1], args[2])
	chequeoError(err, _ERROR_VIEW_VISITORS)
	fmt.Println(_OK)
}

// ejecutarAgregarArchivo ejecuta el comando agregar_archivo
// Agrega un archivo al analizador
// Si la cantidad de argumentos no es la esperada, imprime un mensaje de error y finaliza la ejecución
// Si hay un error al agregar el archivo, imprime un mensaje de error y finaliza la ejecución
// Si no hay errores, imprime "OK"
func ejecutarAgregarArchivo(args []string, analyzer commands.Analyzer) {
	chequeoArgs(args, 2, _ERROR_ADD_FILE)
	dicTemporal := ADTMap.CreateBST[string, []time.Time](commands.CmpIPStr)
	err := analyzer.AgregarArchivo(args[1], dicTemporal)
	chequeoError(err, _ERROR_ADD_FILE)
	fmt.Println(_OK)
}

// =============== FUNCIONES AUXILIARES ===============

// errorSalida imprime un mensaje de error por stderr y finaliza la ejecución
func errorSalida(mensaje string) {
	fmt.Fprintln(os.Stderr, mensaje)
	os.Exit(0)
}

// chequeoError imprime un mensaje de error y finaliza la ejecución si hay un error
func chequeoError(err error, mensaje string) {
	if err != nil {
		errorSalida(mensaje)
	}
}

// chequeoArgs imprime un mensaje de error y finaliza la ejecución si la cantidad de argumentos no es la esperada
func chequeoArgs(args []string, cantidad int, mensaje string) {
	if len(args) != cantidad {
		errorSalida(mensaje)
	}
}

// =============== PROCESAMIENTOS ===============

// procesarComando procesa los comandos ingresados
// args: lista de argumentos del comando
// analizador: instancia del analizador de comandos
func procesarComando(args []string, analyzer commands.Analyzer) {
	if len(args) == 0 {
		os.Exit(0)
	}

	switch args[0] {
	case _ADD_FILE:
		ejecutarAgregarArchivo(args, analyzer)

	case _VIEW_VISITORS:
		ejecutarVerVisitantes(args, analyzer)

	case _VIEW_MOST_VISITED:
		ejecutarVerMasVisitados(args, analyzer)

	default:
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", args[0])
		os.Exit(0)
	}
}

// =============== MAIN ===============

// main es la función principal
// Inicializa el analizador de comandos y procesa la entrada
func main() {
	analyzer := commands.CreateAnalyzer()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		comando := scanner.Text()
		args := strings.Fields(comando)
		procesarComando(args, analyzer)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error leyendo entrada: %v\n", err)
		os.Exit(0)
	}
}
