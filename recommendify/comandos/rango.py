from _tdas.grafo import Grafo
from biblioteca.biblioteca import bfs


def rango(grafo: Grafo, n: int, cancion: str) -> int:
    """
    rango recibe un grafo, un entero n y una cancion como parametros, y devuelve la cantidad de canciones
    que se encuentran a 'n' saltos de la cancion pasada por parametro.
    """
    if cancion not in grafo.obtener_vertices():
        return 0

    canciones_n_saltos = 0

    _, dic_orden = bfs(grafo, cancion)

    for orden in dic_orden.values():
        if orden == n:
            canciones_n_saltos += 1

    return canciones_n_saltos
