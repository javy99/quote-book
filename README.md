# Quote Book REST API (Go)

A simple REST API to manage quotes using Go standard library only.

---

## 📦 How to Run

```bash
go run main.go
```

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
