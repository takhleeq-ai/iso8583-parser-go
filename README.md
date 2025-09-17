# ISO8583 Parser (Go)

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)
<!-- Uncomment the build badge after you add the GitHub Actions workflow
[![Build](https://github.com/takhleeq-ai/iso8583-parser-go/actions/workflows/go.yml/badge.svg)](https://github.com/takhleeq-ai/iso8583-parser-go/actions)
-->

A minimal ISO 8583 message parser written in Go — built step by step to learn Golang fundamentals.

---

## ⚡ Features

- Parse `|`-separated ISO 8583 fields (MTI, PAN, Amount, Transmission DateTime)
- Validate field count
- Map fields into a strongly typed struct
- Command-line interface (CLI) input
- Unit tests included

---

## ⚙️ Usage

Clone the repository:

```bash
git clone https://github.com/takhleeq-ai/iso8583-parser-go.git
cd iso8583-parser-go
```

Run with a sample message:

```bash
go run ./cmd/parser "0200|1234567890123456|000000100000|20250916083000"
```

If no message is provided, a default sample is used.

---

## 📋 Example Output

```text
0200|1234567890123456|000000100000|20250916083000
Field 0: 0200
Field 1: 1234567890123456
Field 2: 000000100000
Field 3: 20250916083000
{MTI:0200 PAN:1234567890123456 Amount:000000100000 TransmissionDateTime:20250916083000}
```

---

## 🧪 Run Tests

```bash
go test ./cmd/parser
```

---

## 🎯 Learning Outcomes

This project was built as a hands-on exercise to practice:

- Go modules and folder structure
- CLI apps
- Error handling
- Unit testing
- Structs, slices, and string parsing

---

## 📜 License

MIT License — see [LICENSE](LICENSE)
