Here's a basic `README.md` file for your **Go Fiber Books CRUD API project**, with instructions on setting up, running, and testing the backend:

---

```markdown
# 📚 Go Fiber Books API

A simple RESTful API built with [Go Fiber](https://gofiber.io/) and [GORM](https://gorm.io/) for managing a collection of books. Includes CRUD operations and test cases.

---

## 🚀 Features

- Create, Read, Update, Delete books
- SQLite database using GORM
- RESTful API endpoints
- Unit tests for all handlers

---

## 🛠️ Tech Stack

- Go
- Fiber (web framework)
- GORM (ORM)
- SQLite (local DB)

---

## 📦 Project Structure
```

go-fiber-books/
├── database/
│ └── db.go
├── handlers/
│ └── book_handler.go
│ └── book_handler_test.go
├── models/
│ └── book.go
├── services/
│ └── book_service.go
├── main.go
└── go.mod / go.sum

````

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/MatheeshaMihiranga/go-fiber-books.git
cd go-fiber-books
````

### 2. Initialize Go Modules

```bash
go mod tidy
```

### 3. Run the App

```bash
go run main.go
```

Server will run on: `http://localhost:3000`

---

## 📚 API Endpoints

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| POST   | `/books`     | Create a new book |
| GET    | `/books`     | Get all books     |
| GET    | `/books/:id` | Get a book by ID  |
| PUT    | `/books/:id` | Update a book     |
| DELETE | `/books/:id` | Delete a book     |

---

## 🧪 Run Tests

```bash
go test -v ./handlers -count=1
```

---

## 📝 Example Book Payload

```json
{
  "title": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "year": 2015
}
```
