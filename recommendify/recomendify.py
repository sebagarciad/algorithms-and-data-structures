#!/usr/bin/python3

from _tdas.grafo import Grafo
from comandos.mas_importantes import mas_importantes
from comandos.camino import camino
from comandos.recomendacion import recomendacion
from comandos.ciclo import ciclo
from comandos.rango import rango

import csv
import sys

ID = "id"
USER_ID = "user_id"
TRACK_NAME = "track_name"
ARTIST = "artist"
PLAYLIST_ID = "playlist_id"
PLAYLIST_NAME = "playlist_name"
GENRES = "genres"
COMA = ","
SEPARADOR_MAS_IMP = "; "
SEPARADOR_CAMINO = " >>>> "
ERROR_COMANDO = "Error en el comando {comando}: {error}"
CAMINO = "camino"
MAS_IMPORTANTES = "mas_importantes"
RECOMENDACION = "recomendacion"
CICLO = "ciclo"
RANGO = "rango"
COMANDO_INVALIDO = "Comando invalido"
ESPACIO = " "
NO_ARCHIVO = "No se encontro el archivo {ruta}"
ERROR_ARCHIVO = "Error al leer el archivo {ruta}: {error}"


# =============================== GRAFOS =============================== #

def grafo_bipartito(datos: list) -> Grafo:
    """
    Crea un grafo no dirigido, en el cual los vertices son los usuarios y las canciones del archivo analizado
    y las aristas indican si una cancion se encuentra en el playlist de un usuario.
    """
    grafo = Grafo(es_dirigido=False)
    for entrada in datos:
        usuario = entrada[USER_ID]
        cancion = f"{entrada[TRACK_NAME]} - {entrada[ARTIST]}"
        playlist = entrada[PLAYLIST_NAME]
        grafo.agregar_vertice(usuario)
        grafo.agregar_vertice(cancion)
        grafo.agregar_arista(usuario, cancion, playlist)
    return grafo


def grafo_canciones(datos: list) -> Grafo:
    """
    Crea un grafo no dirigido en el cual los vertices son canciones y las aristas representan si dos canciones
    se encuentran en las playlists de un mismo usuario.
    """
    grafo = Grafo(es_dirigido=False)
    usuarios = {}

    for entrada in datos:
        usuario = entrada[USER_ID]
        cancion = f"{entrada[TRACK_NAME]} - {entrada[ARTIST]}"
        playlist = entrada[PLAYLIST_NAME]

        if usuario not in usuarios:
            usuarios[usuario] = {}

        if playlist not in usuarios[usuario]:
            usuarios[usuario][playlist] = set()

        usuarios[usuario][playlist].add(cancion)

    for usuario, playlists in usuarios.items():
        canciones_relacionadas = set()
        for playlist, canciones in playlists.items():
            canciones_relacionadas.update(canciones)
        for cancion1 in canciones_relacionadas:
            for cancion2 in canciones_relacionadas:
                if cancion1 != cancion2 and not grafo.estan_unidos(cancion1, cancion2):
                    grafo.agregar_vertice(cancion1)
                    grafo.agregar_vertice(cancion2)
                    grafo.agregar_arista(cancion1, cancion2)

    return grafo


# =============================== FUNCIONES AUX (MAIN) =============================== #

def leer_datos_tsv(ruta: str) -> list:
    """
    Recibe una ruta a un archivo y lo procesa. Devuelve una lista con los datos analizados.
    """
    csv.field_size_limit(sys.maxsize)

    datos = []
    with open(ruta, "r") as archivo:
        lector = csv.DictReader(archivo, delimiter="\t")  # OBS: DictReader interpreta la primera fila como los nombres de las columnas
        for entrada in lector:
            try:
                datos.append({
                    ID: int(entrada[ID.upper()]),
                    USER_ID: entrada[USER_ID.upper()],
                    TRACK_NAME: entrada[TRACK_NAME.upper()],
                    ARTIST: entrada[ARTIST.upper()],
                    PLAYLIST_ID: int(entrada[PLAYLIST_ID.upper()]),
                    PLAYLIST_NAME: entrada[PLAYLIST_NAME.upper()],
                    GENRES: entrada[GENRES.upper()].split(COMA)
                })
            except (ValueError, KeyError) as e:
                continue

    return datos


def cargar_datos(ruta: str) -> list:
    """
    Recibe la ruta de un archivo por parametro y devuelve una lista con los datos del archivo leido.
    """
    try:
        return leer_datos_tsv(ruta)
    except FileNotFoundError:
        print(NO_ARCHIVO.format(ruta=ruta))
    except Exception as e:
        print(ERROR_ARCHIVO.format(ruta=ruta, error=e))
    return None


def procesar_comandos(datos: list):
    """
    Crea un grafo a partir de los datos pasados por parametro y llama a la funcion ejecutar_comandos
    para procesar los comandos ingresados por stdin.
    """
    grafo_usuarios_canciones = grafo_bipartito(datos)
    grafo_canciones_ciclo_rango, pagerank = None, None

    for linea in sys.stdin:
        comando, parametros = obtener_comando_y_parametros(linea)
        if not comando:
            continue

        if comando == MAS_IMPORTANTES and pagerank is None:
            pagerank = mas_importantes(grafo_usuarios_canciones)

        if comando in {CICLO, RANGO} and grafo_canciones_ciclo_rango is None:
            grafo_canciones_ciclo_rango = grafo_canciones(datos)

        ejecutar_comando(comando, parametros, grafo_usuarios_canciones, grafo_canciones_ciclo_rango, pagerank)


def obtener_comando_y_parametros(linea: str) -> str:
    """
    Procesa cada linea de stdin y devuelve el comando y los parametros correspondientes ingresados.
    """
    entrada = linea.strip().split(ESPACIO, 1)
    if len(entrada) < 1:
        return None, None
    comando = entrada[0]
    parametros = entrada[1]
    return comando, parametros


def ejecutar_comando(comando: str, parametros: str, grafo_usuarios_canciones: Grafo, grafo_canciones_ciclo_rango: Grafo, pagerank: list):
    """
    Ejecuta la función correspondiente al comando recibido.
    """
    funciones = {
        CAMINO: lambda: ejecutar_camino(parametros, grafo_usuarios_canciones),
        MAS_IMPORTANTES: lambda: ejecutar_mas_importantes(parametros, pagerank),
        RECOMENDACION: lambda: ejecutar_recomendacion(parametros, grafo_usuarios_canciones),
        CICLO: lambda: ejecutar_ciclo(parametros, grafo_canciones_ciclo_rango),
        RANGO: lambda: ejecutar_rango(parametros, grafo_canciones_ciclo_rango),
    }

    if comando in funciones:
        funciones[comando]()
    else:
        print(COMANDO_INVALIDO)


# =============================== FUNCIONES COMANDOS =============================== #

def ejecutar_camino(parametros: str, grafo: Grafo):
    """
    Ejecuta la funcion 'camino' del modulo comandos.
    """
    try:
        origen, destino = parametros.split(SEPARADOR_CAMINO)
        res = camino(grafo, origen, destino)
        print(res)

    except ValueError as e:
        print(ERROR_COMANDO.format(comando=CAMINO, error=e))


def ejecutar_mas_importantes(parametros: str, pagerank: list):
    """
    Ejecuta la funcion 'mas_importantes' del modulo comandos.
    """
    try:
        n = int(parametros)
        if n > len(pagerank):
            print(pagerank)
        else:
            print(SEPARADOR_MAS_IMP.join(pagerank[:n]))
    except ValueError as e:
        print(ERROR_COMANDO.format(comando=MAS_IMPORTANTES, error=e))


def ejecutar_recomendacion(parametros: str, grafo: Grafo):
    """
    Ejecuta la funcion 'recomendacion' del modulo comandos.
    """
    try:
        tipo, n, canciones = parametros.split(ESPACIO, 2)
        res = recomendacion(grafo, tipo, canciones, int(n))
        print(res)
    except ValueError as e:
        print(ERROR_COMANDO.format(comando=RECOMENDACION, error=e))


def ejecutar_ciclo(parametros: str, grafo: Grafo):
    """
    Ejecuta la funcion 'ciclo' del modulo comandos.
    """
    try:
        n, cancion = parametros.split(ESPACIO, 1)
        res = ciclo(grafo, int(n), cancion)
        print(res)
    except ValueError as e:
        print(ERROR_COMANDO.format(comando=CICLO, error=e))


def ejecutar_rango(parametros: str, grafo: Grafo):
    """
    Ejecuta la funcion 'rango' del modulo comandos.
    """
    try:
        n, cancion = parametros.split(ESPACIO, 1)
        res = rango(grafo, int(n), cancion)
        print(res)
    except ValueError as e:
        print(ERROR_COMANDO.format(comando=RANGO, error=e))


# =============================== MAIN =============================== #

def main():
    if len(sys.argv) < 2:
        return

    ruta = sys.argv[1]

    datos = cargar_datos(ruta)
    if datos is None:
        return

    procesar_comandos(datos)


main()
