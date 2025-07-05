# Reverse Polish Notation (RPN) Calculator

This project is a command-line calculator that evaluates mathematical expressions written in **reverse Polish notation (RPN)**, also known as **postfix notation**.

It was developed as part of coursework for an Algorithms and Data Structures course at the University of Buenos Aires (UBA), where I study Software Engineering in the Faculty of Engineering (FIUBA).

---

## 📌 Description

The calculator reads input **exclusively from standard input (stdin)** — it does not take command-line arguments. Each line is treated as a standalone RPN expression, and the program prints the result to **standard output (stdout)**, one result per line.

If an error occurs during the evaluation of a line (e.g., malformed expression, division by zero, invalid operator usage), the output for that line will be:

---

## ⚙️ Features

- Supports **integer arithmetic** using Go’s `int64` type for large-range calculations.
- Supported operations:
  - **Binary**: `+`, `-`, `*`, `/` (integer division), `^` (exponentiation), `log` (arbitrary base logarithm)
  - **Unary**: `sqrt`
  - **Ternary**: `?` (C-style ternary: `a b c ?` returns `b` if `a ≠ 0`, else `c`)
- All operations are integer-based; results are **truncated** (e.g., `20 / -3` results in `-6`)
- Uses Go’s `math` standard library for power, logarithm, and square root calculations

---

## 🔣 Syntax and Input Format

- Each line is a full **RPN expression** and is evaluated independently.
- Tokens (numbers or operators) are separated by **one or more spaces**.
- **Leading and trailing spaces** are ignored.
- Errors in one line do **not** stop evaluation of the following lines.

### ✍️ RPN Format Rules

- **Binary**: `a b op` → `3 2 -` = `1`
- **Unary**: `a op` → `5 sqrt` = `2`
- **Ternary**: `a b c ?` → `1 -1 0 ?` = `-1` (like `1 ? -1 : 0` in C)

### 📚 Dependencies

- Go 1.18+
- Uses Go’s math standard library for power, root, and log operations

### 🛠️ Compilation and Usage

- Compile with:
    ```
    go build -o dc
    ```

- Run with:
    ```
    ./dc < expressions.txt
    ```

- Or interactively:
    ```
    ./dc
    3 5 + 2 ^
    ```
    
