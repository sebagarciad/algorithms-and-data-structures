import random

DIRIGIDO = "dirigido"
NO_DIRIGIDO = "no dirigido"
GRAFO_DESCRIPCION = "Grafo {} con {} vértices: {}"


class Grafo:
    def __init__(self, es_dirigido=False, vertices_init=[]):
        self.es_dirigido = es_dirigido
        self.vertices = {v: {} for v in vertices_init}

    def __len__(self):
        return len(self.vertices)

    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}

    def borrar_vertice(self, v):
        if v in self.vertices:
            for w in self.adyacentes(v):
                self.borrar_arista(v, w)
            self.vertices.pop(v)

    def agregar_arista(self, v, w, peso=1):
        self.vertices[v][w] = peso
        if not self.es_dirigido:
            self.vertices[w][v] = peso

    def borrar_arista(self, v, w):
        if self.estan_unidos(v, w):
            self.vertices[v].pop(w)
        if not self.es_dirigido:
            self.vertices[w].pop(v)

    def estan_unidos(self, v, w):
        return v in self.vertices and w in self.vertices[v]

    def peso_arista(self, v, w):
        return self.vertices[v][w] if self.estan_unidos(v, w) else None

    def obtener_vertices(self):
        return list(self.vertices)

    def adyacentes(self, v):
        return list(self.vertices[v]) if v in self.vertices else []

    def vertice_aleatorio(self):
        return random.choice(self.obtener_vertices())

    def __str__(self):
        tipo = DIRIGIDO if self.es_dirigido else NO_DIRIGIDO
        return GRAFO_DESCRIPCION.format(tipo, len(self.vertices), self.vertices)
