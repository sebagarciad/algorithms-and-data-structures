package comandos

import (
	ADTMap "data_structures/map"
	ADTHeap "data_structures/priority_queue"
	"strconv"
	"strings"
	"time"
)

// ================== ESTRUCTURAS ==================

// recurso es una estructura que contiene la URL de un recurso y la cantidad de visitas
type recurso struct {
	url     string
	visitas int
}

// analizadorDatos es una estructura que contiene los datos necesarios para analizar los datos de un archivo
type dataAnalyzer struct {
	ips       ADTMap.BSTMap[string, []time.Time]
	recursos  ADTMap.Map[string, int]
	visitados ADTHeap.PriorityQueue[recurso]
}

// Analizador es una interfaz que define las operaciones que se pueden realizar sobre un analizador de datos
type Analyzer interface {
	AgregarArchivo(archivo string, dicTemporal ADTMap.BSTMap[string, []time.Time]) error
	VerVisitantes(ipInicio, ipFin string) error
	VerMasVisitados(n int) error
}

// ================== FUNCIONES DE CMP ==================

// cmpVisitas compara dos recursos por cantidad de visitas
// Devuelve número negativo si a es menor que b, número positivo si a es mayor que b, 0 si son iguales
func cmpVisitas(a, b recurso) int {
	return a.visitas - b.visitas
}

// CmpIPStr compara dos direcciones IP
// Devuelve -1 si ip1 es menor que ip2, 1 si ip1 es mayor que ip2, 0 si son iguales
func CmpIPStr(ip1, ip2 string) int {
	grupos1, grupos2 := strings.Split(ip1, "."), strings.Split(ip2, ".")
	for i := 0; i < 4; i++ {
		n1, _ := strconv.Atoi(grupos1[i])
		n2, _ := strconv.Atoi(grupos2[i])
		if n1 < n2 {
			return -1
		}
		if n1 > n2 {
			return 1
		}
	}
	return 0
}

// ================== CREAR ESTRUCTURAS ==================

// CrearAnalizador crea un analizador de datos
func CreateAnalyzer() Analyzer {
	return &dataAnalyzer{
		ips:       ADTMap.CreateBST[string, []time.Time](CmpIPStr),
		recursos:  ADTMap.NewHash[string, int](),
		visitados: ADTHeap.NewHeap(cmpVisitas),
	}
}

// crearRecurso crea un recurso con la URL y la cantidad de visitas especificadas
func crearRecurso(url string, visitas int) recurso {
	return recurso{url, visitas}
}
