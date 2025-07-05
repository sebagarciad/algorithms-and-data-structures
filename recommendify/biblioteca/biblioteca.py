from _tdas.grafo import Grafo
from _tdas.heap import Heap
from _tdas.cola import Cola
import random
from biblioteca.union_find import UnionFind


# ****************************************************************************** #
# ******************************** RECORRIDOS ********************************** #
# ****************************************************************************** #

def bfs(grafo, origen):
    '''
    Realiza un recorrido de anchura (Breath First Search) sobre un grafo, y devuelve los diccionarios de
    padres y orden.
    '''
    visitados = {origen}
    padres = {origen: None}
    orden = {origen: 0}
    q = Cola()
    q.encolar(origen)

    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.encolar(w)

    return padres, orden


def dfs(grafo, origen):
    '''
    Realiza un recorrido de profundidad (Depth First Search) sobre el grafo,
    a partir del origen pasado por parametro.
    '''
    padres = {origen: None}
    orden = {origen: 0}
    visitados = {origen}

    _dfs(grafo, origen, visitados, padres, orden)

    return padres, orden


def _dfs(grafo, v, visitados, padres, orden):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            orden[w] = orden[v] + 1
            _dfs(grafo, w, visitados, padres, orden)


# ****************************************************************************** #
# ********************************** CICLOS ************************************ #
# ****************************************************************************** #

def ciclo_grafo_no_dirigido(grafo):
    '''
    Busca un ciclo en un grafo no dirigido y, de haberlo, lo devuelve. Si no encuentra un ciclo,
    devuelve None.
    '''
    visitados = {}
    padre = {}
    for v in grafo.obtener_vertices():
        if v not in visitados:
            ciclo = dfs_ciclo(grafo, v, visitados, padre)
            if ciclo is not None:
                return ciclo
    return None


def dfs_ciclo(grafo, v, visitados, padre):
    visitados[v] = True
    for w in grafo.adyacentes(v):
        if w not in visitados:
            if w != padre[v]:
                return reconstruir_ciclo(padre, w, v)
            else:
                padre[w] = v
                ciclo = dfs_ciclo(grafo, w, visitados, padre)
                if ciclo is not None:
                    return ciclo
    return None


def reconstruir_ciclo(padre, inicio, fin):
    v = fin
    camino = []
    while v != inicio:
        camino.append(v)
        v = padre[v]
    camino.append(inicio)
    return camino[::-1]


# ****************************************************************************** #
# ***************************** CAMINOS MINIMOS ******************************** #
# ****************************************************************************** #

def camino_minimo_dijkstra(grafo, origen, destino):
    '''
    Busca un camino minimo entre los vertices de 'origen' y 'destino', a traves del algoritmo de Dijkstra.
    '''
    distancia = {}
    padre = {}

    for v in grafo.obtener_vertices():
        distancia[v] = float('inf')
    distancia[origen] = 0
    padre[origen] = None

    heap = Heap()
    heap.encolar(origen, 0)

    while not heap.esta_vacia():
        v, _ = heap.desencolar()
        if destino is not None and v == destino:
            return padre, distancia

        for w in grafo.adyacentes(v):
            distancia_nueva = distancia[v] + grafo.peso_arista(v, w)
            if distancia_nueva < distancia[w]:
                distancia[w] = distancia_nueva
                padre[w] = v
                heap.encolar(w, distancia[w])

    return padre, distancia


def camino_minimo_bfs(grafo, origen, destino=None):
    '''
    Busca un camino minimo entre los vertices de 'origen' y 'destino', a traves de un recorrido BFS.
    '''
    distancia = {}
    padre = {origen: None}
    visitados = {origen}
    cola = Cola()
    cola.encolar(origen)

    for v in grafo.obtener_vertices():
        distancia[v] = float('inf')

    distancia[origen] = 0

    while not cola.esta_vacia():
        v = cola.desencolar()

        if destino is not None and v == destino:
            break

        for w in grafo.adyacentes(v):
            if w not in visitados:
                distancia[w] = distancia[v] + 1
                padre[w] = v
                visitados.add(w)
                cola.encolar(w)

    return padre, distancia


def reconstruir_camino(padre, origen, destino):
    '''
    Devuelve una lista con el camino minimo encontrado desde el origen hasta el destino (si no hay un
    camino minimo, devuelve una lista vacia).
    '''
    camino = []
    v = destino

    while v is not None:
        camino.append(v)
        v = padre[v]
    if camino[-1] != origen:
        return []

    return camino[::-1]


def reconstruir_camino_minimo(padre, origen, destino):
    '''
    Devuelve una lista con el camino minimo de un grafo, si lo hay, de lo contrario, devuelve una lista vacia.
    '''
    camino = []
    v = destino
    while v is not None:
        camino.append(v)
        v = padre[v]

    if camino[-1] != origen:
        return []

    return camino[::-1]


# ****************************************************************************** #
# ******************************** CENTRALIDAD ********************************* #
# ****************************************************************************** #

def centralidad(grafo):
    '''
    Devuelve la medida de centralidad, del tipo betweenness, de cada vertice en un grafico.
    '''
    cent = {}
    for v in grafo.obtener_vertices():
        cent[v] = 0
    for v in grafo.obtener_vertices():
        for w in grafo.obtener_vertices():
            if v == w:
                continue
            padre, distancia = camino_minimo_dijkstra(grafo, v, None)
            if padre[w] is None:
                continue
            actual = padre[w]
            while actual != v:
                cent[actual] += 1
                actual = padre[actual]
    return cent


# ****************************************************************************** #
# *************************** ARBOL DE TENIDO MINIMO (MST) ********************* #
# ****************************************************************************** #

def mst_prim(grafo):
    '''
    Devuelve un arbol de tendido minimo (MST), a partir del grafo pasado por parametro.
    Implementacion del algoritmo de Prim.
    '''
    v = random.choice(grafo.obtener_vertices())
    visitados = {v}
    heap = Heap()

    for w in grafo.adyacentes(v):
        heap.encolar(grafo.peso_arista(v, w), (v, w))

    arbol = Grafo(es_dirigido=False, vertices_init=grafo.obtener_vertices())

    while not heap.esta_vacia():
        _, (v, w) = heap.desencolar()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, grafo.peso_arista(v, w))
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if x not in visitados:
                heap.encolar(grafo.peso_arista(w, x), (w, x))

    return arbol


def mst_kruskal(grafo):
    '''
    Devuelve un arbol de tendido minimo (MST), a partir del grafo pasado por parametro.
    Implementacion del algoritmo de Kruskal.
    '''
    conjuntos = UnionFind(grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo))
    arbol = Grafo(False, grafo.obtener_vertices())
    for a in aristas:
        v, w, peso = a
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        arbol.agregar_arista(v, w, peso)
        conjuntos.union(v, w)
    return arbol


def obtener_aristas(grafo):
    '''
    Devuelve una lista que contiene las aristas del grafo pasado por parametro.
    '''
    aristas = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v, w, grafo.peso_arista(v, w)))
        visitados.add(v)
    return aristas


# ****************************************************************************** #
# ****************************** GRADOS DE VÉRTICES **************************** #
# ****************************************************************************** #

def grados(g):
    '''
    Recibe un grafo no dirigido y devuelve un diccionario con el grado de cada vertice.
    '''
    resultado = {}
    vertices = g.obtener_vertices()
    for vertice in vertices:
        adyacentes = g.adyacentes(vertice)
        resultado[vertice] = len(adyacentes)
    return resultado


# ****************************************************************************** #
# ************************* PUNTOS DE ARTICULACION - TARJAN ******************** #
# ****************************************************************************** #

def puntos_articulacion(grafo):
    '''
    Recibe un grafo no dirigido y devuelve un conjunto con los puntos de articulacion de dicho grafo,
    obtenidos mediante la implementacion del algoritmo de Tarjan.
    '''
    origen = grafo.vertice_aleatorio()
    puntos_art = set()
    dfs_puntos_articulacion(grafo, origen, {origen}, {origen: None}, {origen: 0}, {}, puntos_art, True)
    return puntos_art


def dfs_puntos_articulacion(grafo, v, visitados, padre, orden, mas_bajo, ptos, es_raiz):
    hijos = 0
    mas_bajo[v] = orden[v]
    for w in grafo.adyacentes(v):
        if w not in visitados:
            hijos += 1
            orden[w] = orden[v] + 1
            padre[w] = v
            visitados.add(w)
            dfs_puntos_articulacion(grafo, w, visitados, padre, orden, mas_bajo, es_raiz=False)

            if mas_bajo[w] >= orden[v] and not es_raiz:
                ptos.add(v)
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
        elif padre[v] != w:
            mas_bajo[v] = min(mas_bajo[v], orden[w])

    if es_raiz and hijos > 1:
        ptos.add(v)
