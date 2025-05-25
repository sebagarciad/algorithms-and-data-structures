package mymap_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	ADTMap "github.com/sebagarciad/algorithms-and-data-structures/map"

	"github.com/stretchr/testify/require"
)

var _TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

// FUNCIONES COMPARACION
func cmpInt(a, b int) int {
	return a - b
}

func cmpStr(a, b string) int {
	return strings.Compare(a, b)
}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := ADTMap.CreateBST[int, string](cmpInt)
	require.EqualValues(t, 0, dic.Count(), "La cantidad de un diccionario vacio debe ser 0")
	require.False(t, dic.Contains(1), "Un diccionario vacio no tiene claves guardadas")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Get(1) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Remove(1) })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Árbol vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := ADTMap.CreateBST[string, string](cmpStr)
	require.False(t, dic.Contains(""), "Debe ser false")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Get("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Remove("") })

	dicNum := ADTMap.CreateBST[int, string](cmpInt)
	require.False(t, dicNum.Contains(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Get(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Remove(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := ADTMap.CreateBST[string, int](cmpStr)
	dic.Save("A", 10)
	require.EqualValues(t, 1, dic.Count(), "La cantidad debe ser 1")
	require.True(t, dic.Contains("A"), "Debe devolver true")
	require.False(t, dic.Contains("B"), "Debe devolver false")
	require.EqualValues(t, 10, dic.Get("A"), "Debe devolver 10")
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Get("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := ADTMap.CreateBST[string, string](cmpStr)
	require.False(t, dic.Contains(claves[0]))
	require.False(t, dic.Contains(claves[0]))
	dic.Save(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Count())
	require.True(t, dic.Contains(claves[0]))
	require.True(t, dic.Contains(claves[0]))
	require.EqualValues(t, valores[0], dic.Get(claves[0]))
	require.EqualValues(t, valores[0], dic.Get(claves[0]))

	require.False(t, dic.Contains(claves[1]))
	require.False(t, dic.Contains(claves[2]))
	dic.Save(claves[1], valores[1])
	require.True(t, dic.Contains(claves[0]))
	require.True(t, dic.Contains(claves[1]))
	require.EqualValues(t, 2, dic.Count())
	require.EqualValues(t, valores[0], dic.Get(claves[0]))
	require.EqualValues(t, valores[1], dic.Get(claves[1]))

	require.False(t, dic.Contains(claves[2]))
	dic.Save(claves[2], valores[2])
	require.True(t, dic.Contains(claves[0]))
	require.True(t, dic.Contains(claves[1]))
	require.True(t, dic.Contains(claves[2]))
	require.EqualValues(t, 3, dic.Count())
	require.EqualValues(t, valores[0], dic.Get(claves[0]))
	require.EqualValues(t, valores[1], dic.Get(claves[1]))
	require.EqualValues(t, valores[2], dic.Get(claves[2]))
}

func TestDiccOrdReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Perro"
	clave2 := "Gato"
	dic := ADTMap.CreateBST[string, string](cmpStr)
	dic.Save(clave, "guau")
	dic.Save(clave2, "miau")
	require.True(t, dic.Contains(clave))
	require.True(t, dic.Contains(clave2))
	require.EqualValues(t, "guau", dic.Get(clave))
	require.EqualValues(t, "miau", dic.Get(clave2))
	require.EqualValues(t, 2, dic.Count())

	dic.Save(clave, "baubau")
	dic.Save(clave2, "miu")
	require.True(t, dic.Contains(clave))
	require.True(t, dic.Contains(clave2))
	require.EqualValues(t, 2, dic.Count())
	require.EqualValues(t, "baubau", dic.Get(clave))
	require.EqualValues(t, "miu", dic.Get(clave2))
}

func TestDicOrdReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := ADTMap.CreateBST[int, int](cmpInt)
	keys := rand.Perm(1000) // Genera una permutación aleatoria de 0 a 499

	for _, key := range keys {
		dic.Save(key, key) // Guarda los datos originales
	}

	for _, key := range keys {
		dic.Save(key, 2*key) // Reemplaza con el doble
	}

	// Se verifica que los datos hayan sido reemplazados
	ok := true
	for _, key := range keys {
		ok = dic.Get(key) == 2*key
		if !ok {
			break
		}
	}

	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := 3
	clave2 := 1
	clave3 := 0
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []int{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := ADTMap.CreateBST[int, string](cmpInt)

	require.False(t, dic.Contains(claves[0]))
	require.False(t, dic.Contains(claves[0]))
	dic.Save(claves[0], valores[0])
	dic.Save(claves[1], valores[1])
	dic.Save(claves[2], valores[2])

	require.True(t, dic.Contains(claves[2]))
	require.EqualValues(t, valores[2], dic.Remove(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Remove(claves[2]) })
	require.EqualValues(t, 2, dic.Count())
	require.False(t, dic.Contains(claves[2]))

	require.True(t, dic.Contains(claves[0]))
	require.EqualValues(t, valores[0], dic.Remove(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Remove(claves[0]) })
	require.EqualValues(t, 1, dic.Count())
	require.False(t, dic.Contains(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Get(claves[0]) })

	require.True(t, dic.Contains(claves[1]))
	require.EqualValues(t, valores[1], dic.Remove(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Remove(claves[1]) })
	require.EqualValues(t, 0, dic.Count())
	require.False(t, dic.Contains(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Get(claves[1]) })
}

func TestDicOrdConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := ADTMap.CreateBST[int, string](cmpInt)
	clave := 10
	valor := "Gatito"

	dic.Save(clave, valor)
	require.EqualValues(t, 1, dic.Count())
	require.True(t, dic.Contains(clave))
	require.EqualValues(t, valor, dic.Get(clave))
	require.EqualValues(t, valor, dic.Remove(clave))
	require.False(t, dic.Contains(clave))
}

func TestDicOrdClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := ADTMap.CreateBST[string, string](cmpStr)
	clave := ""
	dic.Save(clave, clave)
	require.True(t, dic.Contains(clave))
	require.EqualValues(t, 1, dic.Count())
	require.EqualValues(t, clave, dic.Get(clave))
}

func TestDicOrdValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := ADTMap.CreateBST[string, *int](cmpStr)
	clave := "Pez"
	dic.Save(clave, nil)
	require.True(t, dic.Contains(clave))
	require.EqualValues(t, 1, dic.Count())
	require.EqualValues(t, (*int)(nil), dic.Get(clave))
	require.EqualValues(t, (*int)(nil), dic.Remove(clave))
	require.False(t, dic.Contains(clave))
}

func TestDicOrdGuardarYBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces.")

	dic := ADTMap.CreateBST[int, int](cmpInt)
	keys := rand.Perm(1000) // Genera una permutación aleatoria de 0 a 999

	for _, key := range keys {
		dic.Save(key, key) // Guarda el elemento
		require.True(t, dic.Contains(key), "El elemento debería pertenecer al diccionario después de guardarlo")
		dic.Remove(key) // Borra el elemento
		require.False(t, dic.Contains(key), "El elemento no debería pertenecer al diccionario después de borrarlo")
	}
}

func TestDicOrdIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := ADTMap.CreateBST[string, *int](cmpStr)
	dic.Save(claves[0], nil)
	dic.Save(claves[1], nil)
	dic.Save(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterate(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestDicOrdIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := ADTMap.CreateBST[string, int](cmpStr)
	dic.Save(clave1, 6)
	dic.Save(clave2, 2)
	dic.Save(clave3, 3)
	dic.Save(clave4, 4)
	dic.Save(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterate(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestDicOrdIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := ADTMap.CreateBST[string, int](cmpStr)
	dic.Save(clave0, 7)
	dic.Save(clave1, 6)
	dic.Save(clave2, 2)
	dic.Save(clave3, 3)
	dic.Save(clave4, 4)
	dic.Save(clave5, 5)

	dic.Remove(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterate(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumenDicOrd(b *testing.B, n int) {
	dic := ADTMap.CreateBST[string, int](cmpStr)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el árbol */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
	}

	/* Mezcla las claves para que no se guarden en orden en el arbol */
	rand.Shuffle(len(claves), func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
	})

	for i := 0; i < n; i++ {
		dic.Save(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Count(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Contains(claves[i])
		if !ok {
			break
		}
		ok = dic.Get(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Count(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Remove(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !dic.Contains(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Count())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range _TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenDicOrd(b, n)
			}
		})
	}
}

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := ADTMap.CreateBST[string, int](cmpStr)
	iter := dic.Iterator()
	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Current() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Next() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := ADTMap.CreateBST[string, string](cmpStr)

	// Guardar las claves y valores en el diccionario
	dic.Save(claves[0], valores[0])
	dic.Save(claves[1], valores[1])
	dic.Save(claves[2], valores[2])

	iter := dic.Iterator()

	require.True(t, iter.HasNext())
	primero, _ := iter.Current()
	require.True(t, dic.Contains(primero)) // Verifica si la clave pertenece al diccionario

	iter.Next()
	segundo, segundo_valor := iter.Current()
	require.True(t, dic.Contains(segundo))
	require.EqualValues(t, valores[1], segundo_valor) // Verifica el valor correspondiente
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HasNext())

	iter.Next()
	require.True(t, iter.HasNext())
	tercero, _ := iter.Current()
	require.True(t, dic.Contains(tercero))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)

	iter.Next()
	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Current() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Next() })
}

func TestDicOrdIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := ADTMap.CreateBST[string, string](cmpStr)
	claves := []string{"A", "B", "C"}
	dic.Save(claves[0], "")
	dic.Save(claves[1], "")
	dic.Save(claves[2], "")

	dic.Iterator()
	iter2 := dic.Iterator()
	iter2.Next()
	iter3 := dic.Iterator()
	primero, _ := iter3.Current()
	iter3.Next()
	segundo, _ := iter3.Current()
	iter3.Next()
	tercero, _ := iter3.Current()
	iter3.Next()
	require.False(t, iter3.HasNext())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
}

func ejecutarDicOrdPruebasVolumenIterador(b *testing.B, n int) {
	dic := ADTMap.CreateBST[string, *int](cmpStr)

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el Árbol */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
	}

	/* Mezcla las claves para que no se guarden en orden en el arbol */
	rand.Shuffle(len(claves), func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
	})

	for i := 0; i < n; i++ {
		dic.Save(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterator()
	require.True(b, iter.HasNext())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HasNext() {
			ok = false
			break
		}
		c1, v1 := iter.Current()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Next()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HasNext(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorDicOrd(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range _TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarDicOrdPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestDicOrdVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := ADTMap.CreateBST[int, int](cmpInt)

	/* Inserta 'n' parejas en el Árbol */
	for i := 0; i < 10000; i++ {
		dic.Save(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterate(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

// Pruebas iterador interno
func TestIteradorInternoRangoAbbNulo(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	desde := 5
	hasta := 15
	suma := 0

	visitar := func(clave int, valor string) bool {
		suma += clave
		return true
	}

	dic.IterateRange(&desde, &hasta, visitar)
	require.Equal(t, 0, suma, "La suma de todos los elementos de un arbol nulo debe ser 0")
}

func TestIteradorInternoRangoSinCorte(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	// Iterar entre 5 y 15 sin condicion de corte
	desde := 5
	hasta := 15
	suma := 0
	res := 5 + 7 + 10 + 12 + 15

	visitar := func(clave int, valor string) bool {
		suma += clave
		return true
	}

	dic.IterateRange(&desde, &hasta, visitar)
	require.Equal(t, res, suma, "La suma debe ser igual a la variable res")
}

func TestIteradorInternoConCorte(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	suma := 0
	desde := 5
	hasta := 15

	// Iterar entre 5 y 15 con condicion de corte (la iteracion corta cuando encuentra la primera clave par)
	visitar1 := func(clave int, valor string) bool {
		suma += clave
		return clave%2 != 0
	}
	res := 5 + 7 + 10
	dic.IterateRange(&desde, &hasta, visitar1)
	require.Equal(t, res, suma, "La suma debe ser igual a la variable res")
}

func TestIteradorInternoSinRango(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	suma := 0

	// Iterar sin rango
	visitar2 := func(clave int, valor string) bool {
		suma += clave
		return true
	}

	res := 3 + 5 + 7 + 10 + 12 + 15 + 20
	dic.IterateRange(nil, nil, visitar2)
	require.Equal(t, res, suma, "La suma debe ser igual a la variable res")
}

func TestIteradorInternoEnRangoInexistente(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	suma := 0

	// Iterar en un rango sin claves (25-30)
	desde := 25
	hasta := 30

	visitar3 := func(clave int, valor string) bool {
		suma += clave
		return true
	}

	dic.IterateRange(&desde, &hasta, visitar3)
	require.Equal(t, 0, suma, "La suma de los elementos de un rango inexistente debe ser 0")
}

func TestIteradorInternoRecorridoInOrder(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	// Guardar elementos en el diccionario
	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	// Prueba de recorrido in-order (desde = 5, hasta = nil)
	desde := 5
	arr := []int{}

	visitar4 := func(clave int, valor string) bool {
		arr = append(arr, clave)
		return true
	}

	dic.IterateRange(&desde, nil, visitar4)
	require.Equal(t, []int{5, 7, 10, 12, 15, 20}, arr, "El recorrido in-order debe ser [5, 7, 10, 12, 15, 20]")
}

// Pruebas del iterador externo
func TestIteradorExternoRango(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	// Guardar elementos en el diccionario
	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	// Iterar desde 5 hasta 12 in-order
	desde := 5
	hasta := 12

	iter := dic.IteratorRange(&desde, &hasta)
	require.True(t, iter.HasNext(), "Debe devolver true")

	clave, valor := iter.Current()
	require.Equal(t, 5, clave, "La clave debe ser 5")
	require.Equal(t, "cinco", valor, "El valor debe ser 'cinco'")
	iter.Next()

	clave, valor = iter.Current()
	require.Equal(t, 7, clave, "La clave debe ser 7")
	require.Equal(t, "siete", valor, "El valor debe ser 'siete'")
	iter.Next()

	clave, valor = iter.Current()
	require.Equal(t, 10, clave, "La clave debe ser 10")
	require.Equal(t, "diez", valor, "El valor debe ser 'diez'")
	iter.Next()

	clave, valor = iter.Current()
	require.Equal(t, 12, clave, "La clave debe ser 12")
	require.Equal(t, "doce", valor, "El valor debe ser 'doce'")
	iter.Next()
	require.False(t, iter.HasNext(), "Debe devolver false")
	require.Panics(t, func() { iter.Next() })
	require.Panics(t, func() { iter.Current() })
}

func TestIterarRangoSinElementos(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	// Iterar en un rango sin elementos (25-30)

	desde2 := 25
	hasta2 := 30

	iter2 := dic.IteratorRange(&desde2, &hasta2)
	require.False(t, iter2.HasNext(), "Debe devolver false")
	require.Panics(t, func() { iter2.Next() })
	require.Panics(t, func() { iter2.Current() })
}

func TestIterarRangoSinRango(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	iter3 := dic.IteratorRange(nil, nil)
	require.True(t, iter3.HasNext(), "Debe devolver true")

	clave, valor := iter3.Current()
	require.Equal(t, 3, clave, "La clave debe ser 3")
	require.Equal(t, "tres", valor, "El valor debe ser 'tres'")
	iter3.Next()

	clave, valor = iter3.Current()
	require.Equal(t, 5, clave, "La clave debe ser 5")
	require.Equal(t, "cinco", valor, "El valor debe ser 'cinco'")
	iter3.Next()

	clave, valor = iter3.Current()
	require.Equal(t, 7, clave, "La clave debe ser 7")
	require.Equal(t, "siete", valor, "El valor debe ser 'siete'")
	iter3.Next()

	clave, valor = iter3.Current()
	require.Equal(t, 10, clave, "La clave debe ser 10")
	require.Equal(t, "diez", valor, "El valor debe ser 'diez'")
	iter3.Next()

	clave, valor = iter3.Current()
	require.Equal(t, 12, clave, "La clave debe ser 12")
	require.Equal(t, "doce", valor, "El valor debe ser 'doce'")
	iter3.Next()

	clave, valor = iter3.Current()
	require.Equal(t, 15, clave, "La clave debe ser 15")
	require.Equal(t, "quince", valor, "El valor debe ser 'quince'")
	iter3.Next()

	clave, valor = iter3.Current()
	require.Equal(t, 20, clave, "La clave debe ser 20")
	require.Equal(t, "veinte", valor, "El valor debe ser 'veinte'")
	iter3.Next()
	require.False(t, iter3.HasNext(), "Debe devolver false")
	require.Panics(t, func() { iter3.Next() })
	require.Panics(t, func() { iter3.Current() })
}

func TestIterarRangoHastaNil(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	// Iterar de 10 hasta nil

	desde4 := 10

	iter4 := dic.IteratorRange(&desde4, nil)
	require.True(t, iter4.HasNext())

	clave, valor := iter4.Current()
	require.Equal(t, 10, clave, "La clave debe ser 10")
	require.Equal(t, "diez", valor, "El valor debe ser 'diez'")
	iter4.Next()

	clave, valor = iter4.Current()
	require.Equal(t, 12, clave, "La clave debe ser 12")
	require.Equal(t, "doce", valor, "El valor debe ser 'doce'")
	iter4.Next()

	clave, valor = iter4.Current()
	require.Equal(t, 15, clave, "La clave debe ser 15")
	require.Equal(t, "quince", valor, "El valor debe ser 'quince'")
	iter4.Next()

	clave, valor = iter4.Current()
	require.Equal(t, 20, clave, "La clave debe ser 20")
	require.Equal(t, "veinte", valor, "El valor debe ser 'veinte'")
	iter4.Next()
	require.False(t, iter4.HasNext(), "Debe devolver false")
	require.Panics(t, func() { iter4.Next() })
	require.Panics(t, func() { iter4.Current() })
}

func TestIterarRangoDesdeNil(t *testing.T) {
	dic := ADTMap.CreateBST[int, string](cmpInt)

	dic.Save(10, "diez")
	dic.Save(5, "cinco")
	dic.Save(15, "quince")
	dic.Save(7, "siete")
	dic.Save(12, "doce")
	dic.Save(20, "veinte")
	dic.Save(3, "tres")

	// Iterar de nil hasta 12
	hasta5 := 12

	iter5 := dic.IteratorRange(nil, &hasta5)
	require.True(t, iter5.HasNext(), "Debe devolver true")

	clave, valor := iter5.Current()
	require.Equal(t, 3, clave, "La clave debe ser 3")
	require.Equal(t, "tres", valor, "El valor debe ser 'tres'")
	iter5.Next()

	clave, valor = iter5.Current()
	require.Equal(t, 5, clave, "La clave debe ser 5")
	require.Equal(t, "cinco", valor, "El valor debe ser 'cinco'")
	iter5.Next()

	clave, valor = iter5.Current()
	require.Equal(t, 7, clave, "La clave debe ser 7")
	require.Equal(t, "siete", valor, "El valor debe ser 'siete'")
	iter5.Next()

	clave, valor = iter5.Current()
	require.Equal(t, 10, clave, "La clave debe ser 10")
	require.Equal(t, "diez", valor, "El valor debe ser 'diez'")
	iter5.Next()

	clave, valor = iter5.Current()
	require.Equal(t, 12, clave, "La clave debe ser 12")
	require.Equal(t, "doce", valor, "El valor debe ser 'doce'")
	iter5.Next()
	require.False(t, iter5.HasNext(), "Debe devolver false")
	require.Panics(t, func() { iter5.Next() })
	require.Panics(t, func() { iter5.Current() })
}
