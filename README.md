# Todo List API (Gin + PostgreSQL)

Aplikasi REST API sederhana untuk manajemen todo list menggunakan **Go**, **Gin Framework**, dan **PostgreSQL**. Dibuat sebagai latihan/pelatihan BLKPP.

## Fitur

- 📋 **GET /api/v1/todos** — Melihat semua todo (opsional filter `?done=true` / `?done=false`)
- 🔍 **GET /api/v1/todos/:id** — Melihat detail todo berdasarkan ID
- ➕ **POST /api/v1/todos** — Membuat todo baru
- ✏️ **PUT /api/v1/todos/:id** — Mengupdate todo yang sudah ada
- 🗑️ **DELETE /api/v1/todos/:id** — Menghapus todo
- 🌐 **CORS** — Diaktifkan untuk semua origin (siap diakses dari frontend)
- 📝 **Validasi input** — Title wajib diisi minimal 3 karakter
- 🔧 **Environment variables** — Konfigurasi database via `.env`

## Teknologi

| Komponen | Library / Versi |
|----------|-----------------|
| Bahasa | Go 1.21+ |
| Web Framework | [Gin](https://github.com/gin-gonic/gin) v1.12.0 |
| Database Driver | [pgx](https://github.com/jackc/pgx) v5 (stdlib) |
| Database | PostgreSQL |
| Env Loader | [godotenv](https://github.com/joho/godotenv) v1.5.1 |
| Validator | [go-playground/validator](https://github.com/go-playground/validator) v10 |

## Struktur Project

```
todolist_gin/
├── main.go                 # Entry point aplikasi
├── go.mod                  # Dependensi Go
├── go.sum                  # Checksum dependensi
├── .env                    # Variabel environment (buat sendiri, tidak di-commit)
├── .gitignore              # Mengabaikan .env
├── README.md               # Dokumentasi ini
├── database/
│   └── database.go         # Koneksi PostgreSQL
├── router/
│   └── router.go           # Setup router Gin + CORS
└── todo/
    ├── model.go            # Struct model (Todo, CreateTodoRequest, UpdateTodoRequest)
    ├── handler.go          # Handler CRUD (List, GetByID, Create, Update, Delete)
    └── routes.go           # Pendaftaran route /todos
```

## Persyaratan

- [Go](https://golang.org/dl/) terinstall (versi 1.21 atau lebih baru)
- [PostgreSQL](https://www.postgresql.org/download/) terinstall dan berjalan
- Database sudah dibuat untuk aplikasi ini

## Instalasi

1. **Clone / buka folder project**
   ```bash
   cd "c:\Users\User\OneDrive\Documents\Pelatihan BLKPP\todolist_gin"
   ```

2. **Unduh dependensi**
   ```bash
   go mod tidy
   ```

3. **Buat file `.env`**
   Buat file `.env` di root project dengan isi berikut (sesuaikan dengan konfigurasi PostgreSQL Anda):

   ```env
   DATABASE_URL=postgresql://postgres:password@localhost:5432/todolist?sslmode=disable
   ```

   > ⚠️ Jangan commit file `.env` ke git. File ini sudah diabaikan melalui `.gitignore`.

4. **Buat tabel di database**
   Jalankan SQL berikut di PostgreSQL Anda:

   ```sql
   CREATE TABLE IF NOT EXISTS todos (
       id SERIAL PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       is_done BOOLEAN DEFAULT FALSE,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

5. **Jalankan aplikasi**
   ```bash
   go run main.go
   ```

   Aplikasi akan berjalan di `http://localhost:8080`

## Contoh Penggunaan API

### 1. Lihat semua todo
```bash
curl http://localhost:8080/api/v1/todos
```

### 2. Lihat todo yang sudah selesai saja
```bash
curl "http://localhost:8080/api/v1/todos?done=true"
```

### 3. Lihat detail todo
```bash
curl http://localhost:8080/api/v1/todos/1
```

### 4. Buat todo baru
```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go Gin", "is_done": false}'
```

### 5. Update todo
```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Belajar Go Gin Lanjutan", "is_done": true}'
```

### 6. Hapus todo
```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

## Respons Error

Semua error dikembalikan dalam format JSON seragam:

```json
{
  "error": "Pesan deskripsi error"
}
```

Kode status HTTP yang digunakan:
- `200 OK` — Berhasil
- `201 Created` — Berhasil membuat data baru
- `400 Bad Request` — Input tidak valid / ID harus angka
- `404 Not Found` — Data tidak ditemukan
- `500 Internal Server Error` — Kesalahan server

## Lisensi

Projek ini dibuat untuk tujuan pelatihan/pembelajaran di BLKPP.