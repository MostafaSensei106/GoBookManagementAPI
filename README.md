<h1 align="center">GoBookManagementAPI</h1>

<p align="center">
  <strong>A RESTful API for managing books, authors, and categories, built with Go.</strong>
</p>

<p align="center">
  <a href="#about">About</a> •
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#api-endpoints">API Endpoints</a> •
  <a href="#technologies">Technologies</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

---

## About

Welcome to **GoBookManagementAPI** — a RESTful API built with Go for managing a book store.
This API provides endpoints for performing CRUD operations on books.

---

## Features

### 🌟 Core Functionality

- **Book Management**: CRUD operations for books.
  - Get all books.
  - Get a book by its ID.
  - Create a new book.
  - Update an existing book.
  - Delete a book.

---

## Installation

### 🏗️ Build from Source

> ![Note]
> Make sure you have `Go` and `git` installed on your system.

### ⚙️ Step 1: Clone the Repository

```bash
git clone https://github.com/MostafaSensei106/GoBookManagementAPI.git
cd GoBookManagementAPI
```

### ✅ Step 2: Build and Run

#### Using `go run`

```bash
go run main.go
```

#### Using `Makefile`

The project includes a `Makefile` with commands to build and run the application.

```bash
make build
```

you will get a binary in the `bin` directory , run it
server will start on port `8080`.

#### Using Docker

The project includes a `Dockerfile` to build a container image.

```bash
make docker-build
make docker-run
```

The server will start on port `8080`.

---

## 🚀 Quick Start

Once the server is running, you can use `curl` or any API client to interact with the API.

### Get all books

```bash
curl http://localhost:8080/books
```

### Get a book by ID

```bash
curl http://localhost:8080/book/1
```

---

## API Endpoints

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| GET    | `/books`     | Get all books     |
| GET    | `/book/{id}` | Get a book by ID  |
| POST   | `/book/`     | Create a new book |
| PUT    | `/book/{id}` | Update a book     |
| DELETE | `/book/{id}` | Delete a book     |

---

## Technologies

| Technology         | Description                                                                                  |
| ------------------ | -------------------------------------------------------------------------------------------- |
| 🧠 **Golang**      | [go.dev](https://go.dev) — The core language.                                                |
| 🌐 **Gorilla/Mux** | [gorilla/mux](https://github.com/gorilla/mux) — A powerful URL router and dispatcher for Go. |
| 🗃️ **GORM**        | [gorm.io](https://gorm.io) — The fantastic ORM library for Go.                               |
| 🐳 **Docker**      | [docker.com](https://www.docker.com/) — For containerization.                                |

---

## Contributing

Contributions are welcome! Here’s how to get started:

1. Fork the repository
2. Create a new branch:
   `git checkout -b feature/YourFeature`
3. Commit your changes:
   `git commit -m "Add amazing feature"`
4. Push to your branch:
   `git push origin feature/YourFeature`
5. Open a pull request

> 💡 Please open an issue first for major feature ideas or changes.

---

## License

This project is licensed under the **MIT License**.
See the [LICENSE](LICENSE) file for full details.

<p align="center">
  Made with ❤️ by <a href="https://github.com/MostafaSensei106">MostafaSensei106</a>
</p>
