from _tdas.grafo import Grafo
from _tdas.cola import Cola
from comandos.recomendacion import es_cancion

ERROR_CANCIONES = "Tanto el origen como el destino deben ser canciones"
NO_RECORRIDO = "No se encontro recorrido"
SEPARADOR = " --> "
VACIO = ""
MENSAJE = "{cancion} --> aparece en playlist --> {playlist_1} --> de --> {usuario} --> tiene una playlist --> {playlist_2} --> donde aparece --> "


def camino(grafo: Grafo, origen: str, destino: str):
    """
    camino recibe un grafo, una cancion de origen y otra de destino, y devuelve una lista con el camino
    minimo entre ambas canciones dentro del grafo.
    """
    if es_cancion_valida(grafo, origen) or es_cancion_valida(grafo, destino):
        return ERROR_CANCIONES

    visitados = set()
    padre = {}
    playlist = {}
    cola = Cola()
    cola.encolar(origen)
    visitados.add(origen)

    while not cola.esta_vacia():
        actual = cola.desencolar()
        if actual == destino:
            return reconstruir_camino(padre, playlist, origen, destino)

        for adyacente in grafo.adyacentes(actual):
            if adyacente not in visitados:
                visitados.add(adyacente)
                padre[adyacente] = actual
                playlist[f"{(actual, adyacente)}"] = grafo.peso_arista(actual, adyacente)
                cola.encolar(adyacente)

    return NO_RECORRIDO


def reconstruir_camino(padres, playlist, origen, destino):
    camino = []
    actual = destino
    camino.append(actual)
    while actual != origen:
        padre = padres[actual]
        name_playlist = playlist.get(f"{(padre, actual)}")
        camino.append(name_playlist)
        camino.append(padre)
        actual = padre

    camino.reverse()
    grupos_de_4 = obtener_grupos_de_4(camino)
    camino_final = obtener_camino_final(grupos_de_4)

    return VACIO.join(camino_final)


def es_cancion_valida(grafo: Grafo, cancion: str) -> bool:
    """
    Verifica si un vértice es una canción válida dentro del grafo.
    """
    return cancion not in grafo.obtener_vertices() or not es_cancion(cancion)


def obtener_grupos_de_4(camino):
    """
    Recibe una lista con los vértices del camino y devuelve una lista de listas con los vértices agrupados de a 4.
    """
    grupos_de_4 = []
    for i in range(0, len(camino), 4):
        grupo = camino[i:i+4]
        grupos_de_4.append(grupo)
    return grupos_de_4


def obtener_camino_final(grupos_de_4):
    """
    Recibe una lista con los vértices del camino y devuelve una lista con los mensajes finales a mostrar.
    """
    camino_final = []
    for grupo in grupos_de_4:
        if len(grupo) == 1:
            camino_final.append(grupo[0])
            break
        cancion, playlist_1, usuario, playlist_2 = grupo
        camino_final.append(MENSAJE.format(cancion=cancion, playlist_1=playlist_1, usuario=usuario, playlist_2=playlist_2))
    return camino_final
