package comandos

import (
	"bufio"
	ADTMap "data_structures/map"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	_2_S   = "2s"
	_VACIO = ""
)

// ================== PROCESAMIENTO ARCHIVO ==================

// leerArchivo lee un archivo y procesa cada línea
// Devuelve un error si hay problemas al leer el archivo
// Devuelve nil si no hay errores
// Actualiza el heap de visitados con la información de las visitas
func leerArchivo(ruta string, analyzer *dataAnalyzer, dicTemporal ADTMap.BSTMap[string, []time.Time]) error {
	archivo, err := os.Open(ruta)
	if err != nil {
		return err
	}
	defer archivo.Close()

	escaneo := bufio.NewScanner(archivo)
	for escaneo.Scan() {
		linea := escaneo.Text()
		if err := procesarLinea(linea, analyzer, dicTemporal); err != nil {
			return err
		}
	}

	if err := escaneo.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error al escanear el archivo: %v\n", err)
	}

	actualizarHeapVisitados(analyzer)
	return nil
}

// ================== PROCESAMIENTO DATOS ==================

// leerLinea lee una línea del archivo y devuelve los campos de la misma
// Devuelve la dirección IP, la fecha, el método y el recurso de la línea
func leerLinea(linea string) (string, time.Time, string, string, error) {
	campos := strings.Split(linea, "\t")
	if len(campos) < 4 {
		return _VACIO, time.Time{}, _VACIO, _VACIO, fmt.Errorf("error en el procesamiento del archivo")
	}
	ip := campos[0]
	fecha, err := time.Parse(time.RFC3339, campos[1])
	if err != nil {
		return _VACIO, time.Time{}, _VACIO, _VACIO, err
	}
	metodo := campos[2]
	recurso := campos[3]
	return ip, fecha, metodo, recurso, nil
}

// procesarLinea procesa una línea del archivo
// Actualiza los diccionarios de IPs y recursos con la información de la línea
// Devuelve un error si hay problemas al leer la línea
// Devuelve nil si no hay errores
func procesarLinea(linea string, analyzer *dataAnalyzer, dicTemporal ADTMap.BSTMap[string, []time.Time]) error {
	ip, fecha, _, recurso, err := leerLinea(linea)
	if err != nil {
		return err
	}
	diccionarioIPs := analyzer.ips
	actualizarDiccionarioIPs(ip, fecha, diccionarioIPs)
	actualizarDiccionarioIPs(ip, fecha, dicTemporal)
	actualizarDiccionarioRecursos(recurso, analyzer)
	return nil
}

// ================== ACTUALIZACIÓN DATOS ==================

// actualizarDiccionarioIPs actualiza el diccionario de IPs con la información de la línea
// Si la IP ya está en el diccionario, agrega la fecha a la lista de visitas
// Si la IP no está en el diccionario, crea una nueva entrada con la IP y la fecha
func actualizarDiccionarioIPs(ip string, fecha time.Time, bst ADTMap.BSTMap[string, []time.Time]) {
	if bst.Contains(ip) {
		horariosVisitas := bst.Get(ip)
		horariosVisitas = append(horariosVisitas, fecha)
		bst.Save(ip, horariosVisitas)
	} else {
		bst.Save(ip, []time.Time{fecha})
	}
}

// actualizarDiccionarioRecursos actualiza el diccionario de recursos con la información de la línea
// Si el recurso ya está en el diccionario, incrementa la cantidad de visitas
// Si el recurso no está en el diccionario, guarda el recurso y cantidad de visitas 1
func actualizarDiccionarioRecursos(recurso string, analyzer *dataAnalyzer) {
	if analyzer.recursos.Contains(recurso) {
		analyzer.recursos.Save(recurso, analyzer.recursos.Get(recurso)+1)
	} else {
		analyzer.recursos.Save(recurso, 1)
	}
}

// ================== DETECCIÓN DOS ==================

// detectarIPsSospechosas detecta las IPs sospechosas que realizaron 5 visitas en menos de 2 segundos
// Devuelve un diccionario con las IPs sospechosas
func detectarIPsSospechosas(bst ADTMap.BSTMap[string, []time.Time]) ADTMap.BSTMap[string, bool] {
	suspicious := ADTMap.CreateBST[string, bool](CmpIPStr)
	for iter := bst.Iterator(); iter.HasNext(); iter.Next() {
		ip, visitas := iter.Current()
		for i := 0; i < len(visitas)-4; i++ {
			primero := visitas[i]
			ultimo := visitas[i+4]
			duracion, _ := time.ParseDuration(_2_S)
			if ultimo.Sub(primero) < duracion {
				suspicious.Save(ip, true)
				break
			}
		}
	}
	return suspicious
}

// imprimirIPsSospechosas imprime las IPs sospechosas
func imprimirIPsSospechosas(suspicious ADTMap.Map[string, bool]) {
	suspicious.Iterate(func(clave string, dato bool) bool {
		fmt.Printf("DoS: %s\n", clave)
		return true
	})
}

// ================== AGREGAR ARCHIVO ==================

// AgregarArchivo agrega un archivo al analizador de datos
// Devuelve un error si hay problemas al leer el archivo
// Devuelve nil si no hay errores
// Detecta las IPs sospechosas y las imprime
func (analyzer *dataAnalyzer) AgregarArchivo(ruta string, dicTemporal ADTMap.BSTMap[string, []time.Time]) error {
	if err := leerArchivo(ruta, analyzer, dicTemporal); err != nil {
		return err
	}

	sospechosos := detectarIPsSospechosas(dicTemporal)
	imprimirIPsSospechosas(sospechosos)
	return nil
}
