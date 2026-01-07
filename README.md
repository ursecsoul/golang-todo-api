# Todo List REST API (Golang)

REST API **Todo List sederhana** yang dibuat menggunakan **Golang (net/http)**.
Project ini cocok untuk **belajar CRUD**, **konsep HTTP API**, dan **penulisan dokumentasi REST API**.

---

## Fitur

- CRUD Todo (Create, Read, Update, Delete)
- Endpoint RESTful
- Request & response menggunakan JSON
- Penggunaan HTTP status code yang sesuai
- Struktur project rapi dan sederhana
- In-memory storage (tanpa database)

---

## Teknologi

- Golang `1.22+`
- net/http (standard library)

---

## Struktur Project

- todo-api/
- ├── go.mod
- ├── main.go
- ├── handlers/
- │   └── todo_handler.go
- ├── models/
- │   └── todo.go
- ├── routes/
- │   └── routes.go
- └── storage/
- │   └── memory.go

---

## Cara Menjalankan Project

### 1. Clone Repository

git clone https://github.com/ursecsoul/golang-todo-api.git
cd todo-api

### 2. Jalankan Server

go run main.go

### 3. Informasi Server

Base URL: http://localhost:8080

---

## Model Data

### Object Todo

{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Membuat REST API",
  "status": "pending",
  "created_at": "2026-01-07T11:15:30+07:00"
}

| Field       | Tipe   | Deskripsi       |
| ----------- | ------ | --------------- |
| id          | int    | ID unik todo    |
| title       | string | Judul todo      |
| description | string | Deskripsi todo  |
| status      | string | pending / done  |
| created_at  | time   | Waktu pembuatan |

## Dokumentasi API

### Create Todo

**POST** `/todos`

#### Request Body

{
  "title": "Belajar Golang",
  "description": "Membuat REST API"
}

#### Response (201 Created)

{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Membuat REST API",
  "status": "pending",
  "created_at": "2026-01-07T11:15:30+07:00"
}

---

### Get List Todo

**GET** `/todos`

#### Response (200 OK)

[
  {
    "id": 1,
    "title": "Belajar Golang",
    "description": "Membuat REST API",
    "status": "pending",
    "created_at": "2026-01-07T11:15:30+07:00"
  }
]

---

### Get Detail Todo

**GET** `/todos/{id}`

#### Contoh

GET /todos/1

#### Response (200 OK)

{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Membuat REST API",
  "status": "pending",
  "created_at": "2026-01-07T11:15:30+07:00"
}


#### Response (404 Not Found)

Todo tidak ditemukan

---

### Update Todo

**PUT** `/todos/{id}`

#### Request Body

{
  "title": "Belajar Golang Lanjutan",
  "description": "CRUD & Dokumentasi",
  "status": "done"
}

#### Response (200 OK)

{
  "id": 1,
  "title": "Belajar Golang Lanjutan",
  "description": "CRUD & Dokumentasi",
  "status": "done",
  "created_at": "2026-01-07T11:15:30+07:00"
}

---

### Delete Todo

**DELETE** `/todos/{id}`

#### Contoh

DELETE /todos/1

#### Response

204 No Content

---

## Response Error

### 400 Bad Request

Title wajib diisi

### 404 Not Found

Todo tidak ditemukan

---

## Pengujian API

### Menggunakan Postman

1. Buat request baru
2. Pilih HTTP method
3. Masukkan URL endpoint
4. Body → raw → JSON
5. Klik Send

### Menggunakan curl

curl -X GET http://localhost:8080/todos

curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-d '{"title":"Belajar API","description":"CRUD Golang"}'

---

## HTTP Status Code

| Kode | Keterangan  |
| ---- | ----------- |
| 200  | OK          |
| 201  | Created     |
| 204  | No Content  |
| 400  | Bad Request |
| 404  | Not Found   |

---

Penulis
[Adinda Rachmania]

