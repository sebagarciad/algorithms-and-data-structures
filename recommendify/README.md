# Recomendify – Spotify Recommender Engine (Go)

**Recomendify** is a command-line application built in Go that models a music recommendation system using graph-based structures. Based on real data from a Spotify public dataset, the program analyzes user-song relationships and enables features like shortest paths, recommendations, and cycle detection.

## 🎯 Project Overview

This project was developed as part of the **Algorithms and Data Structures** course at the University of Buenos Aires. It emphasizes data modeling, graph algorithms, and efficiency in handling large-scale input data (~2 million entries).

## 📦 Features

- **Shortest path between songs** using BFS, treating users and playlists as connectors
- **Global song ranking** via PageRank (or HITS)
- **Personalized recommendations** for users or songs using Personalized PageRank
- **Song similarity analysis**: list songs exactly `n` hops away
- **Cycle detection**: find a cycle of length `n` starting from a song

## 📂 Dataset Format

Input is a TSV file with columns:

ID USER_ID TRACK_NAME ARTIST PLAYLIST_ID PLAYLIST_NAME GENRES

Each row maps a song to a user-created playlist. Songs may appear in multiple playlists and users may create multiple playlists.

## 🧠 Graph Models

- **Bipartite Graph (Users–Songs):**  
  Edges between users and songs they added to any playlist.

- **Song Graph:**  
  Undirected edges between songs that appear in playlists from the same user.

These models enable traversal and recommendation logic.

## 💻 Usage

```bash
$ ./recomendify.py spotify-mini.tsv
```

Then, commands are entered via stdin (or redirected from a file):

🔹 **Shortest Path**

```camino <song1> >>>> <song2>```

Finds the shortest connection between two songs via users and playlists.

🔹 **Most Important Songs**

```mas_importantes <n>```

Lists the top n most central songs using PageRank.

🔹 **Personalized Recommendation**

```recomendacion canciones <n> <song1> >>>> <song2> >>>> ...```

```recomendacion usuarios <n> <song1> >>>> <song2> >>>> ...```

Recommends songs or users based on personalized PageRank from a given song list.

🔹 **Cycle Detection**

```ciclo <n> <song>```

Finds a cycle of exactly n songs starting and ending at the given song.

🔹 **Songs at Exact Distance**

```rango <n> <song>```

Counts how many songs are exactly n hops away in the song similarity graph.


## 💻 Examples of usage

**Shortest Path**

Input:

```camino Don't Go Away - Oasis >>>> Quitter - Eminem```

Output:

```Don't Go Away - Oasis --> aparece en playlist --> misturo tudãao ;x --> de --> 8902446 --> tiene una playlist --> sóo nacionais' --> donde aparece --> Ela Vai Voltar (Todos Os Defeitos de Uma Mulher Perfeita) - Charlie Brown Jr --> aparece en playlist --> Playlist da Yara --> de --> yarits --> tiene una playlist --> Playlist da Yara --> donde aparece --> Quitter - Eminem```

**Most important songs**

Input:

```mas_importantes 20```

Output:

```Bad Romance - Lady Gaga; Poker Face - Lady Gaga; Telephone (feat. Beyoncé) - Lady Gaga; Paparazzi - Lady Gaga; Halo - Beyoncé; Viva La Vida - Coldplay; Single Ladies (Put a Ring on It) - Beyoncé; Decode - Paramore; In The End - Linkin Park; Levo Comigo - Restart; Leave Out All The Rest - Linkin Park; Broken-Hearted Girl - Beyoncé; Alejandro - Lady Gaga; If I Were A Boy - Beyoncé; I Gotta Feeling - Black Eyed Peas; Amo Noite E Dia - Jorge e Mateus; Sweet Dreams - Beyoncé; Smells Like Teen Spirit - Nirvana; Wonderwall - Oasis; Just Dance (feat. Colby O'Donis) - Lady Gaga```

**Personalized Recommendation - Songs**

Input:

```recomendacion canciones 10 Love Story - Taylor Swift >>>> Toxic - Britney Spears >>>> I Wanna Be Yours - Arctic Monkeys >>>> Hips Don't Lie (feat. Wyclef Jean) - Shakira >>>> Death Of A Martian - Red Hot Chili Peppers```

Output:

```Butterfly - Grimes; Cola - Lana Del Rey; In Time - FKA Twigs; Touch - Troye Sivan; Hurricane - 30 Seconds To Mars; Boring - The Pierces; Cut Your Teeth - Kyla La Grange; Earned It - The Weeknd; Player (Feat. Chris Brown) - Tinashe; If I Were A Boy - Beyoncé```

**Personalized Recommendation - Users**

Input:

```recomendacion usuarios 5 Love Story - Taylor Swift >>>> Toxic - Britney Spears >>>> I Wanna Be Yours - Arctic Monkeys >>>> Hips Don't Lie (feat. Wyclef Jean) - Shakira >>>> Death Of A Martian - Red Hot Chili Peppers```

Output:

```lorenafazion; naosoumodinha; hlovato906gmail; tiagogabbana19; extralouca```

**Cicle Detection**

Input:

```ciclo 7 By The Way - Red Hot Chili Peppers```

Output:

```By The Way - Red Hot Chili Peppers --> Fairy Tale - Shaman --> I Hate Everything About You - Three Days Grace --> Viva La Vida - Coldplay --> Under The Bridge - Red Hot Chili Peppers --> November Rain - Guns N' Roses --> Cryin' - Aerosmith --> By The Way - Red Hot Chili Peppers```

**Songs at Exact Distance**

Input:

```rango 3 Shots - Imagine Dragons```

Output:

```2171```