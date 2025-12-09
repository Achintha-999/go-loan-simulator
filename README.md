# 🏦 Go Loan Simulator &nbsp; ![Go](https://img.shields.io/badge/Go-1.19%2B-00ADD8?logo=go&logoColor=white)

Welcome to **Go Loan Simulator**!  
A lightweight, interactive command-line tool written in pure [Go](https://go.dev/) for quickly estimating monthly loan installments.

---

## 📋 Table of Contents

- [✨ Features](#-features)
- [🛠️ Requirements](#-requirements)
- [🚀 Installation](#-installation)
- [💻 Usage](#-usage)
- [📖 Example](#-example)
- [🛡️ License](#-license)
- [📫 Contact](#-contact)

---

## ✨ Features

- 📦 **Simple CLI app** — no setup or dependencies required.
- 🔢 Interactive prompts for:
  - Loan amount
  - Interest rate (annual, %)
  - Loan duration (years)
- 🧮 Calculates and displays monthly installment amount (rounded to the nearest integer)
- 🟦 Written in Go: Fast, cross-platform, and easy to use!

---

## 🛠️ Requirements

- ![Go](https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg) **Go 1.19+**  
  [Get Go here](https://go.dev/doc/install)
- 🖥️ Works on Linux, macOS, and Windows

---

## 🚀 Installation

1. **Clone the repo:**
    ```sh
    git clone https://github.com/Achintha-999/go-loan-simulator.git
    cd go-loan-simulator
    ```

2. **Build it:**
    ```sh
    go build -o go-loan-simulator
    ```

---

## 💻 Usage

Just launch the compiled binary:

```sh
./go-loan-simulator
```

You'll see prompts in your terminal for the three inputs:
- **Loan amount** (integer)
- **Interest rate** (annual % integer)
- **Time period** (years, integer)

The program then prints your monthly installment.

---

## 📖 Example

```sh
$ ./go-loan-simulator
Enter loan amount: 120000
Enter interest rate: 7
Enter time period(years): 10
Installment amount is: 1375
```

> 💡 *The formula used is illustrative and doesn't provide advanced features like amortization schedules or different loan types.*

---

## 🛡️ License

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

---

## 📫 Contact

👤 **Maintainer:** [Achintha-999](https://github.com/Achintha-999)

---  
![Go Logo](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png)




