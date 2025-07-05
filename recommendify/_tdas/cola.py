class NodoCola:
    def __init__(self, dato, proximo=None):
        self.dato = dato
        self.proximo = proximo


class Cola:
    _MENSAJE_COLA_VACIA = "La cola está vacía"

    def __init__(self):
        self.primero = None
        self.ultimo = None

    def esta_vacia(self) -> bool:
        return self.primero is None and self.ultimo is None

    def ver_primero(self):
        if self.esta_vacia():
            raise ValueError(self._MENSAJE_COLA_VACIA)
        return self.primero.dato

    def encolar(self, elemento) -> None:
        nuevo_nodo = NodoCola(elemento)
        if self.esta_vacia():
            self.primero = nuevo_nodo
        else:
            self.ultimo.proximo = nuevo_nodo
        self.ultimo = nuevo_nodo

    def desencolar(self):
        if self.esta_vacia():
            raise ValueError(self._MENSAJE_COLA_VACIA)
        elemento = self.primero.dato
        self.primero = self.primero.proximo
        if self.primero is None:
            self.ultimo = None
        return elemento
