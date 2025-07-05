from _tdas.grafo import Grafo

NO_RECORRIDO = "No se encontro recorrido"
SEPARADOR = " --> "


def ciclo(grafo: Grafo, n: int, cancion: str) -> str:
    """
    ciclo recibe como parametros un grafo, un entero n y una cancion, y devuelve un ciclo de largo n, que
    comienza con la cancion del parametro.
    """
    visitados = set()
    camino = _ciclo(grafo, n, cancion, visitados, cancion)
    return SEPARADOR.join(camino) if camino else NO_RECORRIDO


def _ciclo(grafo: Grafo, n: int, actual: str, visitados: set, origen: str) -> list:
    if n == 0:
        return [actual] if actual == origen else []

    visitados.add(actual)

    for adyacente in grafo.adyacentes(actual):
        if adyacente == origen and n == 1:
            return [actual, origen]
        if adyacente not in visitados:
            camino = _ciclo(grafo, n - 1, adyacente, visitados, origen)
            if camino:
                return [actual] + camino

    visitados.remove(actual)
    return []
