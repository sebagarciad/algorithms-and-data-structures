package linked_list_test

import (
	"testing"

	ADTList "github.com/sebagarciad/algorithms-and-data-structures/linked_list"

	"github.com/stretchr/testify/require"
)

const (
	_INT_VOL = 100000
)

// Pruebas del TDA Lista
func TestListaVacia(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()

	require.True(t, lista.IsEmpty(), "Verifica que la lista este vacia al ser creada")
	require.Equal(t, 0, lista.Length(), "El Length de una lista recien creada debe ser 0")
	require.Panics(t, func() { lista.DeleteFirst() }, "No se puede borrar el primer elemento de una lista vacia")
	require.Panics(t, func() { lista.SeeFirst() }, "Una lista vacia no tiene primer elemento")
	require.Panics(t, func() { lista.SeeLast() }, "Una lista vacia no tiene ultimo elemento")
}

func TestListaUnElemento(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()

	require.True(t, lista.IsEmpty())
	require.Equal(t, 0, lista.Length())

	lista.InsertFirst(1)
	require.Equal(t, 1, lista.SeeFirst(), "El primer elemento deberia ser 1")
	require.Equal(t, 1, lista.SeeLast(), "El ultimo elemento deberia ser 1")
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.False(t, lista.IsEmpty(), "La lista no esta vacia")
	require.Equal(t, 1, lista.DeleteFirst(), "DeleteFirst() debe devolver 1")
	require.True(t, lista.IsEmpty(), "Una lista con todos los elementos borrados esta vacia")
	require.Equal(t, 0, lista.Length(), "El Length de una lista vacia es 0")
	require.Panics(t, func() { lista.SeeFirst() }, "Una lista vacia no tiene primer elemento")
	require.Panics(t, func() { lista.SeeLast() }, "Una lista vacia no tiene ultimo elemento")

	lista.InsertLast(5)
	require.Equal(t, 5, lista.SeeFirst(), "El primer elemento deberia ser 1")
	require.Equal(t, 5, lista.SeeLast(), "El ultimo elemento deberia ser 1")
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.False(t, lista.IsEmpty(), "Debe devolver false: la lista no esta vacia")
	require.Equal(t, 5, lista.DeleteFirst(), "DeleteFirst() debe devolver 5")
	require.True(t, lista.IsEmpty(), "Debe devolver true: la lista esta vacia")
	require.Equal(t, 0, lista.Length(), "El Length de la lista debe ser 0")
	require.Panics(t, func() { lista.SeeFirst() }, "Una lista vacia no tiene primer elemento")
	require.Panics(t, func() { lista.SeeLast() }, "Una lista vacia no tiene ultimo elemento")
}

func TestInsertarVariosElementosAlPrincipio(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()

	require.True(t, lista.IsEmpty())

	lista.InsertFirst(1)
	lista.InsertFirst(2)
	lista.InsertFirst(3) // Lista: [3, 2, 1]
	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")
	require.Equal(t, 3, lista.SeeFirst(), "El primer elemento deberia ser 3")
	require.Equal(t, 1, lista.SeeLast(), "El ultimo elemento deberia ser 1")
	require.False(t, lista.IsEmpty(), "Debe devolver false: la lista no esta vacia")
	require.Equal(t, 3, lista.DeleteFirst(), "Debe devolver 3") // Lista: [2, 1]
	require.Equal(t, 2, lista.Length(), "El Length de la lista es 2")
	require.Equal(t, 2, lista.SeeFirst(), "El primer elemento debe ser 2")

	lista.InsertFirst(5) // Lista: [5, 2, 1]
	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")
	require.Equal(t, 5, lista.SeeFirst(), "El primer elemento de la lista es 5")
	require.Equal(t, 5, lista.DeleteFirst(), "Debe devolver 5") // Lista: [2, 1]
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.Equal(t, 2, lista.DeleteFirst(), "Debe devolver 2") // Lista: [1]
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.Equal(t, 1, lista.DeleteFirst(), "Debe devolver 1") // Lista: []
	require.True(t, lista.IsEmpty(), "Debe devolver true: la lista esta vacia")
	require.Equal(t, 0, lista.Length(), "El Length de una lista vacia es 0")
}

func TestInsertarVariosElementosAlFinal(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	require.Equal(t, 0, lista.Length())

	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3) // Lista: [1, 2, 3]
	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")
	require.Equal(t, 1, lista.SeeFirst(), "El primero elemento deberia ser 1")
	require.Equal(t, 3, lista.SeeLast(), "El ultimo elemento deberia ser 3")
	require.False(t, lista.IsEmpty(), "La lista no esta vacia")

	require.Equal(t, 1, lista.DeleteFirst(), "Debe devolver 1") // Lista: [2, 3]
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.Equal(t, 2, lista.SeeFirst(), "El primer elemento de la lista es 2")

	require.Equal(t, 2, lista.DeleteFirst(), "Debe devolver 2") // Lista: [3]
	require.Equal(t, 1, lista.Length(), "El Length de la lista es 1")
	require.Equal(t, 3, lista.SeeFirst(), "El primer elemento de la lista es 3")

	require.Equal(t, 3, lista.DeleteFirst(), "Debe devolver 3") // Lista: []
	require.Equal(t, 0, lista.Length(), "El Length de la lista debe ser 0")
	require.True(t, lista.IsEmpty(), "Debe devolver true: la lista esta vacia")
}

func TestBorrarElementosListaNoVacia(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()

	require.True(t, lista.IsEmpty(), "IsEmpty() debe devolver true con una lista recien creada")
	require.Equal(t, 0, lista.Length(), "El Length de una lista recien creada deber ser 0")

	lista.InsertLast(1) // Lista: [1]
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	lista.InsertLast(2) // Lista: [1, 2]
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	lista.InsertLast(3) // Lista: [1, 2, 3]
	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")

	require.Equal(t, 1, lista.SeeFirst(), "El primero elemento deberia ser 1")
	require.Equal(t, 3, lista.SeeLast(), "El ultimo elemento deberia ser 3")
	require.False(t, lista.IsEmpty(), "Debe devolver false: la lista no esta vacia")

	require.Equal(t, 1, lista.DeleteFirst(), "Debe devolver 1") // Lista: [2, 3]
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.Equal(t, 2, lista.SeeFirst(), "El primer elemento de la lista debe ser 2")

	require.Equal(t, 2, lista.DeleteFirst(), "Debe devolver 1") // Lista: [3]
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.Equal(t, 3, lista.SeeFirst(), "El primer elemento de la lista es 3")
	require.Equal(t, 3, lista.SeeLast(), "El ultimo elemento de la lista es 3")

	require.Equal(t, 3, lista.DeleteFirst(), "Debe devolver 3") // Lista: []
	require.Equal(t, 0, lista.Length(), "El Length de una lista vacia es 0")
	require.True(t, lista.IsEmpty(), "Debe devolver true: la lista esta vacia")
}

func TestOperacionesIntercaladas(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()

	require.True(t, lista.IsEmpty(), "IsEmpty() debe devolver true en una lista recien creada")
	require.Equal(t, 0, lista.Length(), "El Length de una lista recien creada debe ser 0")

	lista.InsertFirst(1) // Lista: [1]
	require.Equal(t, 1, lista.SeeFirst(), "El primer elemento de la lista es 1")
	lista.InsertLast(2) // Lista: [1, 2]
	require.Equal(t, 2, lista.SeeLast(), "El ultimo elemento deberia ser 2")
	lista.InsertFirst(3) // Lista: [3, 1, 2]

	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")
	require.Equal(t, 3, lista.SeeFirst(), "El primer elemento deberia ser 3")
	require.Equal(t, 2, lista.SeeLast(), "El ultimo elemento deberia ser 2")
	require.False(t, lista.IsEmpty(), "IsEmpty() debe devolver false: la lista tiene elementos")

	// Intercalando inserciones al principio y al final
	lista.InsertLast(4)   // Lista: [3, 1, 2, 4]
	lista.InsertFirst(10) // Lista: [10, 3, 1, 2, 4]
	require.Equal(t, 5, lista.Length(), "El Length de la lista debe ser 5")

	lista.InsertLast(5)   // Lista: [10, 3, 1, 2, 4, 5]
	lista.InsertFirst(20) // Lista: [20, 10, 3, 1, 2, 4, 5]
	require.Equal(t, 7, lista.Length(), "El Length de la lista debe ser 7")

	// Borrando elementos intercaladamente
	require.Equal(t, 20, lista.DeleteFirst(), "Debe devolver 20") // Lista: [10, 3, 1, 2, 4, 5]
	require.Equal(t, 10, lista.DeleteFirst(), "Debe devolver 10") // Lista: [3, 1, 2, 4, 5]
	require.Equal(t, 3, lista.SeeFirst(), "El primer elemento debe ser 3")
	require.Equal(t, 5, lista.SeeLast(), "El ultimo elemento debe ser 5")
	require.Equal(t, 5, lista.Length(), "El Length de la lista debe ser 5")

	lista.InsertFirst(15) // Lista: [15, 3, 1, 2, 4, 5]
	lista.InsertLast(25)  // Lista: [15, 3, 1, 2, 4, 5, 25]
	require.Equal(t, 7, lista.Length(), "El Length de la lista debe ser 7")
	require.Equal(t, 15, lista.SeeFirst(), "El primer elemento debe ser 15")
	require.Equal(t, 25, lista.SeeLast(), "El ultimo elemento debe ser 25")

	// Eliminación de todos los elementos
	require.Equal(t, 15, lista.DeleteFirst(), "Debe devolver 15") // Lista: [3, 1, 2, 4, 5, 25]
	require.Equal(t, 3, lista.DeleteFirst(), "Debe devolver 3")   // Lista: [1, 2, 4, 5, 25]
	require.Equal(t, 1, lista.DeleteFirst(), "Debe devolver 1")   // Lista: [2, 4, 5, 25]
	require.Equal(t, 2, lista.DeleteFirst(), "Debe devolver 2")   // Lista: [4, 5, 25]
	require.Equal(t, 4, lista.DeleteFirst(), "Debe devolver 4")   // Lista: [5, 25]
	require.Equal(t, 5, lista.DeleteFirst(), "Debe devolver 5")   // Lista: [25]
	require.Equal(t, 25, lista.DeleteFirst(), "Debe devolver 25") // Lista: []

	require.Equal(t, 0, lista.Length(), "El Length de una lista sin elementos debe ser 0")
	require.True(t, lista.IsEmpty(), "IsEmpty() debe devolver true: la lista esta vacia")
}

// Pruebas de volumen del TDA Lista
func TestVolumenLista(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()

	for i := 0; i < _INT_VOL; i++ {
		lista.InsertFirst(i)
		require.Equal(t, i, lista.SeeFirst(), "El primer elemento debe ser '%d'", i)
		require.False(t, lista.IsEmpty(), "IsEmpty() debe devolver falso")
	}
	require.Equal(t, _INT_VOL, lista.Length(), "El Length debe ser '%d'", _INT_VOL)

	for i := _INT_VOL - 1; i >= 0; i-- {
		primero := lista.DeleteFirst()
		require.Equal(t, i, primero, "El elemento borrado debe ser '%d'", i)
	}

	require.Equal(t, 0, lista.Length(), "El primer elemento debe ser 0")
	require.True(t, lista.IsEmpty(), "IsEmpty() debe devolver falso")

	for i := 0; i < _INT_VOL; i++ {
		lista.InsertLast(i)
		require.Equal(t, i, lista.SeeLast(), "El ultimo elemento debe ser '%d'", i)
		require.False(t, lista.IsEmpty(), "IsEmpty() debe devolver falso")
	}

	require.Equal(t, _INT_VOL, lista.Length(), "El Length debe ser '%d'", _INT_VOL)

	for i := 0; i < _INT_VOL; i++ {
		primero := lista.DeleteFirst()
		require.Equal(t, i, primero, "El elemento borrado debe ser '%d'", i)
	}

	require.Equal(t, 0, lista.Length(), "El primer elemento debe ser 0")
	require.True(t, lista.IsEmpty(), "IsEmpty() debe devolver falso")
}

// Pruebas de tipos del TDA Lista
func TestListaCadenas(t *testing.T) {
	lista := ADTList.NewLinkedList[string]()

	require.True(t, lista.IsEmpty(), "IsEmpty debe devolver true con una lista recien creada")
	require.Equal(t, 0, lista.Length(), "El Length de una lista recien creada debe ser 0")

	lista.InsertFirst("Hola")
	lista.InsertLast("世界")

	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.Equal(t, "Hola", lista.SeeFirst(), "La primera cadena deberia ser 'Hola'")
	require.Equal(t, "世界", lista.SeeLast(), "La ultima cadena deberia ser '世界'")
	require.False(t, lista.IsEmpty(), "Debe devolver false: la lista tiene elementos")

	require.Equal(t, "Hola", lista.DeleteFirst(), "Debe devolver 'Hola")
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.Equal(t, "世界", lista.SeeFirst(), "La primera cadena de la lista es '世界'")

	require.Equal(t, "世界", lista.DeleteFirst(), "Debe devolver '世界'")
	require.Equal(t, 0, lista.Length(), "Luego de borrar todos los elementos, el Length de la lista debe ser 0")
	require.True(t, lista.IsEmpty(), "Debe devolver true: se borraron todos los elementos de la lista")
}

func TestListaFloats(t *testing.T) {
	lista := ADTList.NewLinkedList[float64]()

	require.True(t, lista.IsEmpty(), "IsEmpty debe devolver true con una lista recien creada")
	require.Equal(t, 0, lista.Length(), "El Length de una lista recien creada debe ser 0")

	lista.InsertFirst(1.123) // Lista: [1.123]
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	lista.InsertLast(22.456789) // Lista: [1.123, 22.456789]
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	lista.InsertLast(4444.789456123) // Lista: [1.123, 22.456789, 4444.789456123]
	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")
	require.False(t, lista.IsEmpty(), "Debe devolver false: la lista no esta vacia")

	require.Equal(t, 1.123, lista.SeeFirst(), "El primer elemento de la lista es 1.123")
	lista.InsertFirst(5.00) // Lista: [5.00, 1.123, 22.456789, 4444.789456123]
	require.Equal(t, 4, lista.Length(), "El Length de la lista debe ser 4")

	require.Equal(t, 5.00, lista.SeeFirst(), "El primer elemento de la lista debe ser 5.00")
	require.Equal(t, 5.00, lista.DeleteFirst(), "Debe devolver 5.00") // Lista: [1.123, 22.456789, 4444.789456123]
	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")
	require.Equal(t, 1.123, lista.SeeFirst(), "El primer elemento de la lsita es 1.123")
}

func TestListaArreglos(t *testing.T) {
	lista := ADTList.NewLinkedList[[]int]()

	require.True(t, lista.IsEmpty(), "IsEmpty debe devolver true con una lista recien creada")
	require.Equal(t, 0, lista.Length(), "El Length de una lista recien creada debe ser 0")

	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6, 7}

	lista.InsertFirst(arr1)
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	lista.InsertLast(arr2)
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.False(t, lista.IsEmpty(), "Debe devolver false: la lista no esta vacia")

	require.Equal(t, arr1, lista.SeeFirst(), "El primer elemento de la lista es [1, 2, 3]")
	require.Equal(t, arr1, lista.DeleteFirst(), "Debe devolver [1, 2, 3]")
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")

	require.Equal(t, arr2, lista.SeeFirst(), "El primer elemento de la lista debe ser [4, 5, 6, 7]")
	require.Equal(t, arr2, lista.SeeLast(), "El ultimo elemento de la lista debe ser [4, 5, 6, 7]")
	require.Equal(t, arr2, lista.DeleteFirst(), "Debe devolver [4, 5, 6, 7]")
	require.Equal(t, 0, lista.Length(), "El Length de una lista con todos los elementos borrados debe ser 0")
	require.True(t, lista.IsEmpty(), "Debe devolver true: la lista esta vacia")
}

// Pruebas iterador interno
func TestIterarListaUnElemento(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(42)

	suma := 0
	lista.Iterate(func(dato int) bool {
		suma += dato
		return true
	})

	require.Equal(t, 42, suma, "La suma debe ser 42")
}

func TestIterarConVariosElementos(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(-2)
	lista.InsertLast(-3)
	lista.InsertLast(-5)
	lista.InsertLast(8)
	lista.InsertLast(-6)

	suma := 0
	lista.Iterate(func(dato int) bool {
		suma += dato
		return true
	})

	require.Equal(t, -7, suma, "La suma de los elementos debe ser -7")
}

func TestIterarListaStrings(t *testing.T) {
	lista := ADTList.NewLinkedList[string]()
	lista.InsertLast("estamos")
	lista.InsertLast("haciendo")
	lista.InsertLast("la")
	lista.InsertLast("lista")

	concatenado := ""
	lista.Iterate(func(dato string) bool {
		concatenado += dato + " "
		return true
	})

	require.Equal(t, "estamos haciendo la lista ", concatenado, "La concatenacion debe ser 'estamos haciendo la lista '")
}

func TestIterarConCorteEnMedio(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4)

	valores := []int{}
	lista.Iterate(func(dato int) bool {
		valores = append(valores, dato)
		return len(valores) < 2
	})

	require.Equal(t, []int{1, 2}, valores, "La iteracion se debe detener despues del segundo elemento")
}

func TestIterarDetenerseEnPrimerElemento(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(10)
	lista.InsertLast(20)
	lista.InsertLast(30)

	valores := []int{}
	lista.Iterate(func(dato int) bool {
		valores = append(valores, dato)
		return false
	})

	require.Equal(t, []int{10}, valores, "La iteracion se debe detener despues del primer elemento")
}

func TestIterarConCondicion(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(3)
	lista.InsertLast(5)
	lista.InsertLast(9)
	lista.InsertLast(12)
	lista.InsertLast(14)

	multiplosDeTres := []int{}
	lista.Iterate(func(dato int) bool {
		if dato%3 == 0 {
			multiplosDeTres = append(multiplosDeTres, dato)
		}
		return true
	})

	require.Equal(t, []int{3, 9, 12}, multiplosDeTres, "Se deben haber acumulado los elementos multiplos de 3")
}

// Pruebas del iterador externo
func TestIterarUnaListaVacia(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	iter := lista.Iterator()

	require.False(t, iter.HasNext(), "No hay Next en una lista vacia")
	require.Panics(t, func() { iter.Next() }, "Next deberia lanzar panico al iterar sobre una lista vacia")
	require.Panics(t, func() { iter.SeeCurrent() }, "SeeCurrent deberia lanzar panico al iterar sobre una lista vacia")
	require.Panics(t, func() { iter.Delete() }, "Borrar deberia lanzar panico al iterar sobre una lista vacia")
}

func TestIterarUnaListaConUnElemento(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertFirst(1) // Lista: [1]
	iter := lista.Iterator()

	require.True(t, iter.HasNext(), "Debe devolver true cuando hay un elemento en la lista")
	require.Equal(t, 1, iter.SeeCurrent(), "El elemento actual del iterador debe ser 1")

	iter.Next()
	require.False(t, iter.HasNext(), "Al avanzar al Next elemento en una lista de un elemento, HasNext() debe devolver false")
	require.Panics(t, func() { iter.SeeCurrent() }, "SeeCurrent deberia lanzar panico al finalizar la iteracion")
	require.Panics(t, func() { iter.Delete() }, "Borrar deberia lanzar panico al finalizar la iteracion")
	require.Panics(t, func() { iter.Next() }, "No Hay elemento Next")

	iter = lista.Iterator()
	require.Equal(t, 1, iter.Delete(), "Borrar deberia devolver el unico elemento de la lista")
	require.True(t, lista.IsEmpty(), "La lista deberia estar vacia despues de eliminar el unico elemento")
	require.False(t, iter.HasNext(), "El iterador no deberia tener Next despues de borrar el unico elemento")
}

func TestIterarUnaListaConVariosElementos(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4) // Lista: [1, 2, 3, 4]
	iter := lista.Iterator()
	elementos := []int{}

	for iter.HasNext() {
		elementos = append(elementos, iter.SeeCurrent())
		iter.Next()
	}

	require.Equal(t, []int{1, 2, 3, 4}, elementos, "El arreglo 'elementos' debe contener los 4 elementos de la lista en orden")
	require.False(t, iter.HasNext(), "Debe devolver false")
	require.Panics(t, func() { iter.SeeCurrent() }, "SeeCurrent deberia lanzar panico al finalizar la iteracion")
	require.Panics(t, func() { iter.Delete() }, "Borrar deberia lanzar panico al finalizar la iteracion")
}

func TestInsertarenListaVacia(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	iter := lista.Iterator()

	iter.Insert(1) // Lista: [1]

	require.Equal(t, 1, lista.SeeFirst(), "El primer elemento de la lista deberia ser 1")
	require.Equal(t, 1, lista.SeeFirst(), "El ultimo elemento de la lista deberia ser 1")
	require.Equal(t, 1, iter.SeeCurrent(), "El elemento actual del iterador deberia ser 1")
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.False(t, lista.IsEmpty(), "La lista no deberia estar vacia")
}

func TestInsertarAlMedio(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4) // Lista: [1, 2, 3, 4]
	iter := lista.Iterator()

	require.True(t, iter.HasNext(), "HasNext() debe devolver true")
	iter.Next() // Actual: 2
	iter.Next() // Actual: 3

	iter.Insert(50) // Lista: [1, 2, 50, 3, 4]
	require.Equal(t, 50, iter.SeeCurrent(), "El elemento actual del iterador debe ser el insertado (50)")
	require.Equal(t, 5, lista.Length(), "El Length de la lista debe ser 5")
	require.Equal(t, 1, lista.SeeFirst(), "El primer elemento de la lista debe ser 1")
	require.Equal(t, 4, lista.SeeLast(), "El ultimo elemento de la lista debe ser 4")

	require.Equal(t, 1, lista.DeleteFirst(), "Debe devolver 1")
	require.Equal(t, 2, lista.DeleteFirst(), "Debe devolver 2")
	require.Equal(t, 50, lista.DeleteFirst(), "Debe devolver 50")
	require.Equal(t, 3, lista.DeleteFirst(), "Debe devolver 3")
	require.Equal(t, 4, lista.DeleteFirst(), "Debe devolver 4")
	require.True(t, lista.IsEmpty())
}

func TestInsertarAlPrincipio(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4) // Lista: [1, 2, 3, 4]
	iter := lista.Iterator()

	require.True(t, iter.HasNext(), "Debe devolver true")

	iter.Insert(50) // Lista: [50, 1, 2, 3, 4]
	require.Equal(t, 5, lista.Length(), "El Length de la lista debe ser 5")
	require.Equal(t, 50, lista.SeeFirst(), "El primer elemento de la lista debe ser 5")
	require.Equal(t, 50, iter.SeeCurrent(), "El elemento actual del iterador debe ser el insertado (5)")
	require.Equal(t, 4, lista.SeeLast(), "El ultimo elemento de la lista debe ser 4")

	require.Equal(t, 50, lista.DeleteFirst(), "Debe devolver 50")
	require.Equal(t, 1, lista.DeleteFirst(), "Debe devolver 1")
	require.Equal(t, 2, lista.DeleteFirst(), "Debe devolver 2")
	require.Equal(t, 3, lista.DeleteFirst(), "Debe devolver 3")
	require.Equal(t, 4, lista.DeleteFirst(), "Debe devolver 4")
	require.True(t, lista.IsEmpty())
}

func TestInsertarAlFinal(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4) // Lista: [1, 2, 3, 4]
	iter := lista.Iterator()

	for iter.HasNext() {
		iter.Next()
	}

	iter.Insert(50) // Lista: [1, 2, 3, 4, 50]
	require.Equal(t, 50, iter.SeeCurrent(), "El elemento actual del iterador debe ser el insertado (50)")
	require.Equal(t, 50, lista.SeeLast(), "El ultimo elemento de la lista debe ser 50")
	require.Equal(t, 5, lista.Length(), "El Length de la lista debe ser 5")
	require.Equal(t, 1, lista.SeeFirst(), "El primer elemento de la lista debe ser 1")

	require.Equal(t, 1, lista.DeleteFirst(), "Debe devolver 1")
	require.Equal(t, 2, lista.DeleteFirst(), "Debe devolver 2")
	require.Equal(t, 3, lista.DeleteFirst(), "Debe devolver 3")
	require.Equal(t, 4, lista.DeleteFirst(), "Debe devolver 4")
	require.Equal(t, 50, lista.DeleteFirst(), "Debe devolver 50")
}

func TestBorrarElementoEnListaUnElemento(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1) // Lista: [1]
	iter := lista.Iterator()

	require.False(t, lista.IsEmpty(), "Debe devolver false: la lista no esta vacia")
	require.Equal(t, 1, lista.SeeFirst(), "El primer elemento de la lista debe ser 1")
	require.Equal(t, 1, iter.SeeCurrent(), "El elemento actual del iterador debe ser 1")

	iter.Delete() // Lista: []

	require.Equal(t, 0, lista.Length(), "El Length de la lista debe ser 0")
	require.True(t, lista.IsEmpty(), "Debe devolver true: la lista esta vacia")
}

func TestBorrarElementoEnElMedioDeLista(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4)      // Lista: [1, 2, 3, 4]
	iter := lista.Iterator() // Actual: 1

	require.True(t, iter.HasNext(), "Debe devolver true")
	require.Equal(t, 1, iter.SeeCurrent(), "El elemento actual del iterador debe ser 1")
	iter.Next() // Actual: 2
	require.Equal(t, 2, iter.SeeCurrent(), "El elemento actual del iterador debe ser 2")
	iter.Next() // Actual: 3
	require.Equal(t, 3, iter.SeeCurrent(), "El elemento actual del iterador debe ser 3")

	require.Equal(t, 3, iter.Delete(), "Debe devolver 3")
	require.Equal(t, 4, iter.SeeCurrent(), "El elemento actual del iterador debe ser 4")
	require.Equal(t, 3, lista.Length(), "El Length de la lista debe ser 3")
}

func TestBorrarElementoEnPrincipioDeLista(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4)      // Lista: [1, 2, 3, 4]
	iter := lista.Iterator() // Actual: 1

	require.True(t, iter.HasNext(), "Debe devolver true")
	require.Equal(t, 1, iter.SeeCurrent(), "El elemento actual del iterador debe ser 1")

	require.Equal(t, 1, iter.Delete(), "Debe devolver 1") // Actual: 2
	require.Equal(t, 2, iter.SeeCurrent(), "El elemento actual del iterador debe ser 2")
}

func TestBorrarElementoAlFinalDeLista(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	lista.InsertLast(1)
	lista.InsertLast(2)
	lista.InsertLast(3)
	lista.InsertLast(4)      // Lista: [1, 2, 3, 4]
	iter := lista.Iterator() // Actual: 1

	for iter.HasNext() {
		iter.Next()
	}

	require.Panics(t, func() { iter.Delete() }, "Se llego al final de la lista, no hay elementos para borrar")

	iter = lista.Iterator()
	for i := 1; i < lista.Length(); i++ {
		iter.Next()
	}

	require.Equal(t, 4, iter.SeeCurrent(), "SeeCurrent() debe ser 4")
	require.Equal(t, 4, iter.Delete(), "Debe devolver 4")
	require.Equal(t, 3, lista.Length(), "Length debe ser 3")
	require.Equal(t, 3, lista.SeeLast(), "El ultimo elemento debe ser 3")
}

func TestIteradorOperacionesIntercaladas(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()
	iter := lista.Iterator()

	iter.Insert(10) // Lista: [10]
	iter.Insert(20) // Lista: [20, 10]

	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.Equal(t, 20, lista.SeeFirst(), "El primer elemento de la lista debe ser 20")
	require.Equal(t, 10, lista.SeeLast(), "El ultimo elemento de la lista debe ser 10")

	iter = lista.Iterator()
	require.Equal(t, 20, iter.Delete(), "Debe devolver 20") // Lista: [10]
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.Equal(t, 10, lista.SeeFirst(), "El primer elemento de la lista debe ser 10")
	require.True(t, iter.HasNext(), "Debe devolver true")

	iter.Insert(30) // Lista: [30, 10]
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.Equal(t, 30, lista.SeeFirst(), "El primer elemento de la lista debe ser 30")

	iter.Next() // Actual: 10

	require.Equal(t, 10, iter.Delete(), "Debe devolver 10") // Lista: [30]
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.Equal(t, 30, lista.SeeFirst(), "El primer elemento de la lista debe ser 30")
	require.False(t, iter.HasNext(), "No debe haber Next")

	iter.Insert(40) // Lista: [30, 40]
	require.Equal(t, 2, lista.Length(), "El Length de la lista debe ser 2")
	require.Equal(t, 30, lista.SeeFirst(), "El primer elemento de la lista debe ser 30")

	iter = lista.Iterator()
	iter.Next() // Actual: 40

	require.Equal(t, 40, iter.Delete(), "Debe devolver 40") // Lista: [30]
	require.Equal(t, 1, lista.Length(), "El Length de la lista debe ser 1")
	require.Equal(t, 30, lista.SeeFirst(), "El primer elemento de la lista debe ser 30")
	require.False(t, iter.HasNext(), "No debe haber Next")

	iter = lista.Iterator()

	require.Equal(t, 30, iter.Delete(), "Debe devolver 30") // Lista: []
	require.True(t, lista.IsEmpty(), "La lista debe estar vacia")
}

func TestVolumenIterador(t *testing.T) {
	lista := ADTList.NewLinkedList[int]()

	iter := lista.Iterator()

	for i := 0; i < _INT_VOL; i++ {
		iter.Insert(i)
		require.Equal(t, i, lista.SeeLast(), "El ultimo elemento de la lista debe ser el '%d'", i)
		require.Equal(t, i+1, lista.Length(), "El Length de la lista debe ser '%d'", i+1)
		iter.Next()
	}

	require.Equal(t, _INT_VOL, lista.Length(), "El Length de la lista debe ser '%d'", _INT_VOL)
	require.Equal(t, 0, lista.SeeFirst(), "El primer elemento de la lista debe ser 0")

	iter = lista.Iterator()

	for i := 0; i < _INT_VOL; i++ {
		require.Equal(t, i, lista.SeeFirst(), "El primer elemento de la lista debe ser '%d'", i)
		require.Equal(t, _INT_VOL-i, lista.Length(), "El Length de la lista debe ser '%d'", _INT_VOL-i)
		iter.Delete()
	}

	require.Equal(t, 0, lista.Length(), "El Length de la lista debe ser 0")
	require.True(t, lista.IsEmpty(), "La lista debe estar vacia")
}
