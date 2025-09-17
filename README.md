# ISO8583 Parser (Go)

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)](https://go.dev/)
[![Build](https://github.com/takhleeq-ai/iso8583-parser-go/actions/workflows/go.yml/badge.svg)](https://github.com/takhleeq-ai/iso8583-parser-go/actions)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

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
