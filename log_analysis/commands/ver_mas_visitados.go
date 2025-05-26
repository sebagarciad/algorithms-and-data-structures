package comandos

import (
	ADTHeap "data_structures/priority_queue"
	"fmt"
)

const (
	_MAS_VISITADOS = "Sitios más visitados:"
)

// ================== ACTUALIZAR HEAP ==================

// actualizarHeapVisitados actualiza el heap de sitios más visitados
// El heap queda ordenado por cantidad de visitas de mayor a menor
func actualizarHeapVisitados(analyzer *dataAnalyzer) {
	arr := make([]recurso, 0, analyzer.recursos.Count())
	for iter := analyzer.recursos.Iterator(); iter.HasNext(); iter.Next() {
		rec, visitas := iter.Current()
		arr = append(arr, crearRecurso(rec, visitas))
	}
	heap := ADTHeap.NewHeapFromArray(arr, cmpVisitas)
	analyzer.visitados = heap
}

// ================== VER MAS VISITADOS ==================

// VerMasVisitados imprime los n sitios más visitados
func (analyzer *dataAnalyzer) VerMasVisitados(n int) error {
	fmt.Println(_MAS_VISITADOS)
	masVisitados := make([]recurso, n)
	i := 0

	for i < n && !analyzer.visitados.IsEmpty() {
		masVisitados[i] = analyzer.visitados.Dequeue()
		i++
	}

	for j := 0; j < i; j++ {
		fmt.Printf("\t%s - %d\n", masVisitados[j].url, masVisitados[j].visitas)
		analyzer.visitados.Enqueue(masVisitados[j])
	}

	return nil
}
