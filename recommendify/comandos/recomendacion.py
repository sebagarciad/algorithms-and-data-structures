from _tdas.grafo import Grafo
from biblioteca.biblioteca import grados as grados_vertice
import random

INT_RECORRIDOS = 100
INT_CAMINOS = 200
SEPARADOR = "; "
CANCIONES = "canciones"
USUARIOS = "usuarios"
FLECHAS = " >>>> "
GUION = " - "


def recomendacion(grafo: Grafo, condicion: str, canciones: str, n: int) -> list:
    """
    recomendacion recibe como parametros un grafo, una condicion, una serie de canciones y un entero n.
    Segun la condicion ingresada ('canciones' o 'usuarios') se devuelve una lista con una cantidad n de
    recomendaciones de canciones o usuarios, respectivamente, en funcion a las canciones pasadas por parametro.
    """
    lista = canciones.split(FLECHAS)
    pagerank = pagerank_personalizado(grafo, INT_CAMINOS, INT_RECORRIDOS, lista)

    recomendacion_tuplas = []
    for vertice, valor in pagerank.items():
        if es_recomendacion_valida(vertice, condicion, canciones):
            recomendacion_tuplas.append((valor, vertice))

    recomendacion_tuplas.sort(reverse=True)

    return SEPARADOR.join(vertice for _, vertice in recomendacion_tuplas[:n])


def pagerank_personalizado(grafo: Grafo, cant_caminos: int, cant_recorridos: int, vertices_iniciales: list):
    grados = grados_vertice(grafo)
    acumulado = {vertice: 0 for vertice in grafo.obtener_vertices()}

    for _ in range(cant_recorridos):
        vertice_actual = random.choice(vertices_iniciales)
        peso = 1

        for _ in range(cant_caminos):
            acumulado[vertice_actual] += peso
            adyacentes = grafo.adyacentes(vertice_actual)
            if not adyacentes:
                break
            grado = grados[vertice_actual]
            peso *= 1 / grado
            vertice_actual = random.choice(adyacentes)

    return acumulado


def es_recomendacion_valida(vertice: str, condicion: str, canciones: str) -> bool:
    """
    Determina si un vértice es válido como recomendación según la condición y las canciones dadas.
    """
    if condicion == CANCIONES:
        return es_cancion(vertice) and vertice not in canciones
    if condicion == USUARIOS:
        return not es_cancion(vertice) and vertice not in canciones
    return False


def es_cancion(vertice: str) -> bool:
    """
    es_cancion indica si el vertice del grafo es una cancion o un usuario.
    """
    return GUION in vertice
