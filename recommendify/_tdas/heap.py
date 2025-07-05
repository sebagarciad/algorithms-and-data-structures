import heapq

class Heap:
    _MENSAJE_HEAP_VACIO = "El heap esta vacio"

    def __init__(self, es_minimo=True):
        self.heap = []
        self.contador = 0
        self.mult = 1 if es_minimo else -1

    def encolar(self, elemento, prioridad):
        heapq.heappush(self.heap, (self.mult * prioridad, self.contador, elemento))
        self.contador += 1

    def desencolar(self):
        if self.esta_vacia():
            raise AssertionError(self._MENSAJE_HEAP_VACIO)
        return heapq.heappop(self.heap)[2]

    def esta_vacia(self):
        return len(self.heap) == 0

    def ver_maximo(self):
        if self.esta_vacia():
            raise AssertionError(self._MENSAJE_HEAP_VACIO)
        return self.heap[0]

    def heapify(self, diccionario):
        self.heap = [(self.multiplier * valor, self.contador, clave) for clave, valor in diccionario.items()]
        self.contador = len(diccionario)
        heapq.heapify(self.heap)
