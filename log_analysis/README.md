# Log Analyzer – DoS Detection & Web Traffic Reporting (Go)

This is a command-line tool written in Go to analyze Apache-style server log files and extract useful information, such as potential denial-of-service (DoS) attacks, visitor IPs within a given range, and the most visited resources.

## 🛠️ Features

- **Add and parse log files** to extract IPs, timestamps, and requested resources.
- **Detect DoS attacks**: alerts when 5+ requests from the same IP occur within 2 seconds.
- **Query visitors**: list all unique visitor IPs within a specified IP range.
- **List most visited resources**: show the top `n` requested URLs, sorted by frequency.
- Handles large files with efficient algorithms and adheres to time complexity constraints.

## 📂 Input Format

Each log file contains lines in the following format (tab-separated):

IP_ADDRESS  TIMESTAMP   HTTP_METHOD RESOURCE_PATH

Example:

208.115.111.72	    2015-05-17T11:05:15+00:00	    GET	    /corrector.html

## 💻 Commands (via stdin)

```bash
agregar_archivo <filename>
ver_visitantes <from_ip> <to_ip>
ver_mas_visitados <n>
```

- agregar_archivo: Parses a log file, updates analytics, and checks for DoS.

- ver_visitantes: Lists IPs in the given range that made requests.

- ver_mas_visitados: Lists the top n most requested resources with counts.

## 🧠 DoS Detection Criteria

A client is flagged for a possible DoS attack if it makes 5 or more requests within a 2-second window.

Example output:

```
DoS: 192.168.1.4
DoS: 200.10.4.2
OK
```

## 📈 Input/Output Examples

**Process a file**

Input: 

```agregar_archivo 20171025.log```

Output:

```
DoS: 192.168.1.4
DoS: 200.10.4.2
OK
```

**Visitors in range:**

Input: 

```ver_visitantes 62.0.0.0 62.255.255.255```

Output:

```
Visitantes:
	62.0.0.0
	62.9.128.3
	62.10.128.3
	62.10.129.3
	62.10.129.4
	62.62.62.62
	62.255.255.255
OK
```

**Top visited resources:**

Input: 

```ver_mas_visitados 10```

Output:

```
Sitios más visitados:
	/algoritmos/tps/2024_1/tp1 - 57
	/algoritmos/faq/ - 35
	/algoritmos/guias/grafos - 25
OK

```

## 🧪 Complexity Constraints

    agregar_archivo: O(n), with n the number of lines in the log.

    ver_mas_visitados: O(s + k log s), with s the number of different sites, and k the parameter.

    ver_visitantes: O(log v) average, O(v) worst-case, with v the number of visitors.

    ### 🛠️ Compilation and Usage

## Compile with:
    ```
    go build -o main
    ```

## Run with:
    ```
    ./main

    ```
