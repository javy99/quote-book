# Quote Book REST API (Go)

A simple REST API to manage quotes using Go standard library only.

---

## 📦 How to Run

```bash
# Install dependencies
go mod tidy

# Start the server
go run main.go
```

Server runs at: `http://localhost:8080`

---

## 📑 Swagger UI (OpenAPI Docs)

This project supports auto-generated Swagger documentation using [swaggo/swag](https://github.com/swaggo/swag).

### 🛠️ Setup (One-Time)

Install `swag` CLI tool if not already:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate docs:

```bash
swag init --generalInfo cmd/main.go --output docs
```

### 🧼 Clean & Re-generate Docs

If you need to regenerate Swagger docs from scratch:

```bash
rm -rf docs
swag init --generalInfo cmd/main.go --output docs
```

### 🌐 View in Browser

After running the server, open:

```
http://localhost:8080/swagger/index.html
```

Interact with the API and explore endpoints.

---

## 🔌 API Endpoints

### ➕ Add a New Quote

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

- **Method**: `POST`
- **URL**: `/quotes`
- **Body**:
  ```json
  {
    "author": "Confucius",
    "quote": "Life is simple, but we insist on making it complicated."
  }
  ```
- **Response**: `201 Created` + JSON of saved quote

---

### 📃 Get All Quotes

```bash
curl http://localhost:8080/quotes
```

- **Method**: `GET`
- **URL**: `/quotes`
- **Response**: List of all quotes

---

### 🎲 Get a Random Quote

```bash
curl http://localhost:8080/quotes/random
```

- **Method**: `GET`
- **URL**: `/quotes/random`
- **Response**: A single random quote

---

### 🔍 Filter Quotes by Author

```bash
curl http://localhost:8080/quotes?author=Confucius
```

- **Method**: `GET`
- **URL**: `/quotes?author=Confucius`
- **Response**: List of quotes by the given author

---

### ❌ Delete Quote by ID

```bash
curl -X DELETE http://localhost:8080/quotes/1
```

- **Method**: `DELETE`
- **URL**: `/quotes/{id}`
- **Response**: `204 No Content` if deleted, `404 Not Found` if not found

---

## 🧪 Running Tests

Unit tests are placed next to their corresponding source files.

### ▶️ Run All Tests Verbosely

```bash
go test ./... -v
```

### ✅ Run with Coverage Report

```bash
go test ./... -cover
```

### 📊 Generate HTML Coverage Report

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

This opens an HTML page in your browser with detailed line-by-line coverage.


---

## 🧪 VS Code REST Client – `requests.http`

This project includes a `requests.http` file for use with the [REST Client extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

### 📥 Install Extension

- Open VS Code
- Go to Extensions (`Ctrl+Shift+X`)
- Search for: `REST Client`
- Install by `humao`

### 📄 requests.http

```http
### Add a new quote
POST http://localhost:8080/quotes
Content-Type: application/json

{
  "author": "Confucius",
  "quote": "Life is simple, but we insist on making it complicated."
}

### Get all quotes
GET http://localhost:8080/quotes

### Get quotes by author
GET http://localhost:8080/quotes?author=Confucius

### Get a random quote
GET http://localhost:8080/quotes/random

### Delete a quote by ID
DELETE http://localhost:8080/quotes/1
```

Open this file in VS Code and click **"Send Request"** above each block to test.

---

## 📂 Project Structure

```
quote-book/
├── cmd/
│   └── main.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── handler/
│   │   ├── handler_test.go
│   │   └── handler.go
│   ├── model/
│   │   └── quote.go
│   └── storage/
│       ├── memory_test.go
│       └── memory.go
├── coverage.out
├── go.mod
├── go.sum
├── README.md
└── requests.http
```

---

## 🧰 Tech Stack

- 🧠 Go (Standard Library)
- 📚 Swagger + Swaggo for API Docs
- 🧪 REST Client (VS Code extension)
- 🧵 In-memory store using sync.Mutex

---

## ✨ Future Ideas

- 💾 Persistent database support (e.g. PostgreSQL, SQLite)
- 🧩 Categorization/tagging of quotes
- 🧹 Sorting & pagination
- 🐳 Docker support

---

## 👨‍💻 Author

Made with ❤️ by [javy99](https://github.com/javy99)

