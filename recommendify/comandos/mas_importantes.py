from _tdas.grafo import Grafo
from comandos.recomendacion import es_cancion

ITERACIONES = 100


def mas_importantes(grafo: Grafo) -> list:
    """
    mas_importantes recibe un grafo y devuelve una lista con las canciones del grafo, ordenadas segun
    su centralidad mediante el algoritmo pagerank, de mayor a menor importancia.
    """
    val_pagerank = pagerank(grafo, 0.85)
    importantes_ordenado = []

    for vertice, valor in val_pagerank.items():
        if es_cancion(vertice):
            importantes_ordenado.append((valor, vertice))

    importantes_ordenado.sort(reverse=True)

    mas_importantes = []

    for valor, vertice in importantes_ordenado:
        mas_importantes.append(vertice)

    return mas_importantes


def pagerank(grafo: Grafo, d: float) -> dict:
    vertices = grafo.obtener_vertices()
    n = len(vertices)
    pagerank = {v: 1 / n for v in vertices}
    entradas = {v: set(grafo.adyacentes(v)) for v in vertices}

    pagerank_actualizado = {v: 0 for v in vertices}

    for _ in range(ITERACIONES):
        pagerank_actualizado = pagerank
        suma_dangling = 0

        for v in vertices:
            contribucion = 0
            pagerank_actualizado[v] = (1 - d) / n
            for w in entradas[v]:
                contribucion += pagerank[w] / len(entradas[w])
            pagerank_actualizado[v] += d * contribucion
            if not entradas[v]:
                suma_dangling += pagerank[v]

        distribucion_dangling = d * suma_dangling / n

        for vertice in vertices:
            pagerank_actualizado[vertice] += distribucion_dangling

    return pagerank
