## Belajar Golang Restfulll API

# 📝 Todo API with JWT Auth

Todo API ini dibuat menggunakan **Golang (Gin + GORM)** dengan fitur:
- ✅ CRUD Todo
- 🔐 JWT Authentication (Login, Register, Refresh Token)
- 👤 Role-based Access Control (RBAC)
- 📖 Swagger Documentation
- 🛡 Middleware Auth & Role Check

---

## ⚙️ Tech Stack
- [Gin](https://gin-gonic.com/) – HTTP Web Framework
- [GORM](https://gorm.io/) – ORM untuk database
- [Swagger](https://swagger.io/) – API Documentation
- [JWT](https://jwt.io/) – Authentication

---

## 🚀 Installation & Run

1. **Clone repository**
   ``` bash
   git clone https://github.com/IkhsanDS/golang-api.git/cd golang-api
   ```
2. **Setup Database**
   Edit file database/connection.go sesuai konfigurasi database kamu (contoh: MySQL/MariaDB).
   Jalankan migrasi:
   ``` bash
   go run main.go
   ```
3. **Run API**
   ``` bash
   go run main.go
   ```
## 📌 API Endpoints
  **Auth**
  - POST /api/v1/auth/register → Register user baru
  - POST /api/v1/auth/login → Login user, dapatkan JWT token
  - GET /api/v1/auth/me → Lihat user yang sedang login (pakai Bearer token)
 **To do**
  - GET /api/v1/todos → List semua todos (public)
  - POST /api/v1/todos → Tambah todo (harus login JWT)
  - PUT /api/v1/todos/:id → Update todo (harus login JWT)
  - DELETE /api/v1/todos/:id → Hapus todo (harus login JWT + role admin)

## 📖 Swagger Docs
   Setelah menjalankan API, buka di browser:
   ``` bash
   👉 http://localhost:8080/swagger/index.html
   ```
## 🔑 Authentication
   Gunakan JWT token di header Authorization:
   ``` makefile
   Authorization: Bearer <access_token>
  ```

  

   
