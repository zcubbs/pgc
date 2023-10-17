# `pgc` PostgreSQL Initializer

`pgc` is a CLI that automates the initialization process of a PostgreSQL database. It creates users, databases, assigns privileges, updates the `pg_hba.conf` file, and restarts the PostgreSQL service.

[![tag](https://img.shields.io/github/tag/zcubbs/pgc)](https://github.com/zcubbs/pgc/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/zcubbs/pgc?status.svg)](https://pkg.go.dev/github.com/zcubbs/pgc)
[![Lint](https://github.com/zcubbs/pgc/actions/workflows/lint.yaml/badge.svg)](https://github.com/zcubbs/pgc/actions/workflows/lint.yaml)
[![Scan](https://github.com/zcubbs/pgc/actions/workflows/scan.yaml/badge.svg?branch=main)](https://github.com/zcubbs/pgc/actions/workflows/scan.yaml)
![Build Status](https://github.com/zcubbs/pgc/actions/workflows/test.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/zcubbs/pgc)](https://goreportcard.com/report/github.com/zcubbs/pgc)
[![Contributors](https://img.shields.io/github/contributors/zcubbs/pgc)](https://github.com/zcubbs/pgc/graphs/contributors)
[![License](https://img.shields.io/github/license/zcubbs/pgc.svg)](./LICENSE)

## Prerequisites

- A running PostgreSQL instance.
- A user with `CREATEDB` and `CREATEROLE` privileges.
- A user with `LOGIN` privileges.

## Installation

```bash
curl -sfL https://raw.githubusercontent.com/zcubbs/pgc/main/install.sh | bash
```

## Usage

```bash
pgc -c config.yaml
```

## Configuration

Edit the `config.yaml` file to set up the PostgreSQL configurations, users, databases, and privileges. Below are the configurable sections:

- `postgresql`: PostgreSQL connection details such as host, port, user, and password.
- `databases`: List of databases to be created along with their owners.
- `users`: List of users to be created along with their passwords.
- `privileges`: Privileges to be granted to users on specific databases.
- `pg_hba`: pg_hba configuration details.
- `pg_hba_conf_path`: Path to the `pg_hba.conf` file.
- `restart_cmd`: Command to restart the PostgreSQL service.

## Development

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Task](https://taskfile.dev/#/installation)

### Run Test Using Docker

1. Build the Docker image:

```bash
docker build -t pgc . && docker run --rm --name pgc -p 5432:5432 pgc
```

the test config file `config.yaml` will be used.

## License

HuB is licensed under the [MIT](./LICENSE) license.
