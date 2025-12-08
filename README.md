# 🏦 Go Loan Simulator

Welcome to **Go Loan Simulator**! This project is designed to help users simulate various loan scenarios, calculate payment schedules, and understand the cost of financing, all built in Go. Whether you're a developer, lender, or borrower, this tool provides valuable insights for loan planning.

---

## 📋 Table of Contents

- [✨ Features](#-features)
- [🛠️ Requirements](#-requirements)
- [🚀 Installation](#-installation)
- [ℹ️ Usage](#-usage)
- [⚙️ Configuration](#-configuration)
- [📖 Example Scenarios](#-example-scenarios)
- [📝 API Reference](#-api-reference)
- [🤝 Contributing](#-contributing)
- [🛡️ License](#-license)

---

## ✨ Features

- 🎯 **Flexible Loan Simulation**  
  Simulate annuity and flat-rate loans for any principal, interest rate, and duration.
- 📆 **Payment Schedule Calculation**  
  Generate detailed amortization tables with principal, interest, and remaining balance breakdowns.
- 📝 **Extensible Design**  
  Easily add more loan types or features without major refactoring.
- 📈 **CLI & API Support**  
  Interact via command-line or integrate into other applications as a library or HTTP API.
- 🔒 **Robust Validation**  
  Handles invalid inputs and edge cases gracefully.

---

## 🛠️ Requirements

- ![Go logo](https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg) **Go 1.19+**  
  Install from the [official Go website](https://go.dev/doc/install).

- 🖥️ **Operating System**  
  Works on Linux, macOS, Windows.

- 🌐 **Network Access (For API Mode)**  
  If using HTTP API, ensure port openness and correct firewall settings.

---

## 🚀 Installation

### 1️⃣ Clone the Repository

```sh
git clone https://github.com/Achintha-999/go-loan-simulator.git
cd go-loan-simulator
```

### 2️⃣ Build the Application

```sh
go build -o go-loan-simulator
```

### 3️⃣ (Optional) Install Globally

```sh
go install
```

---

## ℹ️ Usage

### 💻 Command Line Mode

Simulate a basic annuity loan:

```sh
./go-loan-simulator simulate --type=annuity --principal=50000 --rate=5.5 --years=10
```

Other options:

- `--type`: `annuity` or `flat`
- `--principal`: Loan amount (numeric)
- `--rate`: Annual interest rate (%)
- `--years`: Total loan duration (years)

### 🌐 API Mode

Run as a REST API server:

```sh
./go-loan-simulator serve --port=8080
```

POST requests to `/simulate` endpoint:

```json
{
  "type": "annuity",
  "principal": 60000,
  "rate": 7,
  "years": 7
}
```

---

## ⚙️ Configuration

You can adjust default values using environment variables or config files (see [Config Example](#-example-scenarios)):

- `SIMULATOR_DEFAULT_TYPE`
- `SIMULATOR_DEFAULT_RATE`
- `SIMULATOR_DEFAULT_YEARS`

---

## 📖 Example Scenarios

### 🧑‍💼 Standard Loan Calculation

**CLI:**
```sh
./go-loan-simulator simulate --type=flat --principal=100000 --rate=9 --years=5
```

**API:**
```json
{
  "type": "flat",
  "principal": 100000,
  "rate": 9,
  "years": 5
}
```

### 🏠 Mortgage Planning
Simulate a mortgage with extra payments or irregular schedules (see advanced CLI/API flags).

---

## 📝 API Reference

| Endpoint          | Method | Description                 |
| ------------------|--------|----------------------------|
| `/simulate`       | POST   | Run simulation and get schedule. |

**Request:**
```json
{
  "type": "annuity",
  "principal": 120000,
  "rate": 6.5,
  "years": 15
}
```
**Response:**
```json
{
  "totalPaid": ...,
  "monthlyPayment": ...,
  "schedule": [
     { "month": 1, "principal": ..., "interest": ..., "remaining": ... },
     ...
  ]
}
```

---

## 🤝 Contributing

1. Fork this repo
2. Create a branch (`feature-xyz`)
3. Commit your changes
4. Create a pull request

All contributions are welcome! Please follow the [Go best practices](https://github.com/golang/go/wiki/CodeReviewComments).

---

## 🛡️ License

Released under the [MIT License](LICENSE).

---

## 📫 Contact & Maintainers

- Achintha-999: [GitHub Profile](https://github.com/Achintha-999)

---

## 👏 Acknowledgements

Thanks to all contributors and the Go open source community.

---


