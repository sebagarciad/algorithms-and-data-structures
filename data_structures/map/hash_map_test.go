package mymap_test

import (
	"fmt"
	"testing"

	TDADiccionario "github.com/sebagarciad/algorithms-and-data-structures/map"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.NewHash[string, string]()
	require.EqualValues(t, 0, dic.Count())
	require.False(t, dic.Contains("A"))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Get("A") })
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Remove("A") })
}

func TestDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.NewHash[string, string]()
	require.False(t, dic.Contains(""))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Get("") })
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Remove("") })

	dicNum := TDADiccionario.NewHash[int, string]()
	require.False(t, dicNum.Contains(0))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dicNum.Get(0) })
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dicNum.Remove(0) })
}

func TestUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.NewHash[string, int]()
	dic.Save("A", 10)
	require.EqualValues(t, 1, dic.Count())
	require.True(t, dic.Contains("A"))
	require.False(t, dic.Contains("B"))
	require.EqualValues(t, 10, dic.Get("A"))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Get("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.NewHash[string, string]()
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

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.NewHash[string, string]()
	dic.Save(clave, "miau")
	dic.Save(clave2, "guau")
	require.True(t, dic.Contains(clave))
	require.True(t, dic.Contains(clave2))
	require.EqualValues(t, "miau", dic.Get(clave))
	require.EqualValues(t, "guau", dic.Get(clave2))
	require.EqualValues(t, 2, dic.Count())

	dic.Save(clave, "miu")
	dic.Save(clave2, "baubau")
	require.True(t, dic.Contains(clave))
	require.True(t, dic.Contains(clave2))
	require.EqualValues(t, 2, dic.Count())
	require.EqualValues(t, "miu", dic.Get(clave))
	require.EqualValues(t, "baubau", dic.Get(clave2))
}

func TestReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.NewHash[int, int]()
	for i := 0; i < 500; i++ {
		dic.Save(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Save(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Get(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioRemove(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.NewHash[string, string]()

	require.False(t, dic.Contains(claves[0]))
	require.False(t, dic.Contains(claves[0]))
	dic.Save(claves[0], valores[0])
	dic.Save(claves[1], valores[1])
	dic.Save(claves[2], valores[2])

	require.True(t, dic.Contains(claves[2]))
	require.EqualValues(t, valores[2], dic.Remove(claves[2]))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Remove(claves[2]) })
	require.EqualValues(t, 2, dic.Count())
	require.False(t, dic.Contains(claves[2]))

	require.True(t, dic.Contains(claves[0]))
	require.EqualValues(t, valores[0], dic.Remove(claves[0]))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Remove(claves[0]) })
	require.EqualValues(t, 1, dic.Count())
	require.False(t, dic.Contains(claves[0]))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Get(claves[0]) })

	require.True(t, dic.Contains(claves[1]))
	require.EqualValues(t, valores[1], dic.Remove(claves[1]))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Remove(claves[1]) })
	require.EqualValues(t, 0, dic.Count())
	require.False(t, dic.Contains(claves[1]))
	require.PanicsWithValue(t, "La clave no Contains al diccionario", func() { dic.Get(claves[1]) })
}

func TestReutlizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.NewHash[string, string]()
	clave := "hola"
	dic.Save(clave, "mundo!")
	dic.Remove(clave)
	require.EqualValues(t, 0, dic.Count())
	require.False(t, dic.Contains(clave))
	dic.Save(clave, "mundooo!")
	require.True(t, dic.Contains(clave))
	require.EqualValues(t, 1, dic.Count())
	require.EqualValues(t, "mundooo!", dic.Get(clave))
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.NewHash[int, string]()
	clave := 10
	valor := "Gatito"

	dic.Save(clave, valor)
	require.EqualValues(t, 1, dic.Count())
	require.True(t, dic.Contains(clave))
	require.EqualValues(t, valor, dic.Get(clave))
	require.EqualValues(t, valor, dic.Remove(clave))
	require.False(t, dic.Contains(clave))
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionario.NewHash[avanzado, int]()

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Save(a1, 0)
	dic.Save(a2, 1)
	dic.Save(a3, 2)

	require.True(t, dic.Contains(a1))
	require.True(t, dic.Contains(a2))
	require.True(t, dic.Contains(a3))
	require.EqualValues(t, 0, dic.Get(a1))
	require.EqualValues(t, 1, dic.Get(a2))
	require.EqualValues(t, 2, dic.Get(a3))
	dic.Save(a1, 5)
	require.EqualValues(t, 5, dic.Get(a1))
	require.EqualValues(t, 2, dic.Get(a3))
	require.EqualValues(t, 5, dic.Remove(a1))
	require.False(t, dic.Contains(a1))
	require.EqualValues(t, 2, dic.Get(a3))

}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.NewHash[string, string]()
	clave := ""
	dic.Save(clave, clave)
	require.True(t, dic.Contains(clave))
	require.EqualValues(t, 1, dic.Count())
	require.EqualValues(t, clave, dic.Get(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.NewHash[string, *int]()
	clave := "Pez"
	dic.Save(clave, nil)
	require.True(t, dic.Contains(clave))
	require.EqualValues(t, 1, dic.Count())
	require.EqualValues(t, (*int)(nil), dic.Get(clave))
	require.EqualValues(t, (*int)(nil), dic.Remove(clave))
	require.False(t, dic.Contains(clave))
}

func TestCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.NewHash[string, string]()
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Save(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Count())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Get(claves[i]) == valores[i]
	}

	require.True(t, ok, "Get clave larga funciona")
}

func TestGuardarYRemoveRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	dic := TDADiccionario.NewHash[int, int]()
	for i := 0; i < 1000; i++ {
		dic.Save(i, i)
		require.True(t, dic.Contains(i))
		dic.Remove(i)
		require.False(t, dic.Contains(i))
	}
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.NewHash[string, *int]()
	dic.Save(claves[0], nil)
	dic.Save(claves[1], nil)
	dic.Save(claves[2], nil)

	cs := []string{"", "", ""}
	Count := 0
	cantPtr := &Count

	dic.Iterate(func(clave string, dato *int) bool {
		cs[Count] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, Count)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.NewHash[string, int]()
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

func TestIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.NewHash[string, int]()
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

func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccionario.NewHash[string, int]()

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		dic.Save(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Count(), "La Count de elementos es incorrecta")

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

	require.True(b, ok, "Contains y Get con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Count(), "La Count de elementos es incorrecta")

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

	require.True(b, ok, "Remove muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Count())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta Count de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la Count " +
		"sea la adecuada. Luego validamos que podemos Get y ver si Contains cada una de las claves geeneradas, " +
		"y que luego podemos Remove sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.NewHash[string, int]()
	iter := dic.Iterator()
	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Current() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Next() })
}

func TestDiccionarioIterar(t *testing.T) {
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
	dic := TDADiccionario.NewHash[string, string]()
	dic.Save(claves[0], valores[0])
	dic.Save(claves[1], valores[1])
	dic.Save(claves[2], valores[2])
	iter := dic.Iterator()

	require.True(t, iter.HasNext())
	primero, _ := iter.Current()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Next()
	segundo, segundo_valor := iter.Current()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HasNext())

	iter.Next()
	require.True(t, iter.HasNext())
	tercero, _ := iter.Current()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Next()

	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Current() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Next() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.NewHash[string, string]()
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
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.NewHash[string, string]()
	dic.Save(clave1, "")
	dic.Save(clave2, "")
	dic.Save(clave3, "")
	dic.Remove(clave1)
	dic.Remove(clave2)
	dic.Remove(clave3)
	iter := dic.Iterator()

	require.False(t, iter.HasNext())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Current() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Next() })
	dic.Save(clave1, "A")
	iter = dic.Iterator()

	require.True(t, iter.HasNext())
	c1, v1 := iter.Current()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Next()
	require.False(t, iter.HasNext())
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.NewHash[string, *int]()

	claves := make([]string, n)
	valores := make([]int, n)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
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

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta Count de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.NewHash[int, int]()

	/* Inserta 'n' parejas en el hash */
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
