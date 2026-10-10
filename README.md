# 🛡️ Go Auth API - Clean Architecture

![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=for-the-badge&logo=go)
![Architecture](https://img.shields.io/badge/Architecture-Clean_Architecture-2ea44f?style=for-the-badge)

Proyek ini adalah implementasi RESTful API sederhana untuk sistem Autentikasi (Register & Login) menggunakan **Golang Standard Library (`net/http`)**. Dibangun dengan menerapkan pola **Clean Architecture**.

---

## 🚀 Memulai Proyek (Getting Started)

Ikuti panduan langkah demi langkah berikut untuk menjalankan aplikasi ini di komputer lokalmu.

### 📋 Prasyarat
Pastikan kamu telah menginstal **[Go](https://go.dev/dl/)** di komputermu (disarankan versi 1.18 ke atas).

### ⚙️ Instalasi & Menjalankan Server

1. **Masuk ke direktori proyek melalui terminal:**
   ```bash
   cd go-auth
   ```

2. **Unduh dan sinkronkan dependensi (termasuk bcrypt):**
   ```bash
   go mod tidy
   ```

3. **Jalankan aplikasi backend:**
   ```bash
   go run cmd/server/main.go
   ```

4. **Verifikasi server berjalan:** 
   Jika berhasil, kamu akan melihat log berikut di terminal yang menandakan server siap menerima request di port 8080:
   ```text
   Server starting on :8080
   ```

---

## 📡 Dokumentasi API (Endpoints)

**Base URL:** `http://localhost:8080`

### 1. 📝 Register Akun Baru
Mendaftarkan pengguna baru ke dalam sistem. Sistem secara otomatis akan melakukan *hashing* pada password sebelum disimpan.

*   **URL:** `/api/register`
*   **Method:** `POST`
*   **Headers:** `Content-Type: application/json`

**Body Request:**
```json
{
  "email": "user@example.com",
  "password": "secretpassword"
}
```

**Tabel Response:**

| Status Code | Kondisi | Contoh Response (JSON) |
| :--- | :--- | :--- |
| `201 Created` | Registrasi sukses | `{"message": "Registrasi berhasil"}` |
| `400 Bad Request` | Email terdaftar | `{"message": "Email sudah terdaftar"}` |
| `400 Bad Request` | Format JSON salah | `{"message": "Format JSON tidak valid"}` |

---

### 2. 🔐 Login Pengguna
Melakukan verifikasi kredensial pengguna berdasarkan email dan password yang terdaftar.

*   **URL:** `/api/login`
*   **Method:** `POST`
*   **Headers:** `Content-Type: application/json`

**Body Request:**
```json
{
  "email": "user@example.com",
  "password": "secretpassword"
}
```

**Tabel Response:**

| Status Code | Kondisi | Contoh Response (JSON) |
| :--- | :--- | :--- |
| `200 OK` | Login sukses | `{"message": "Login berhasil"}` |
| `401 Unauthorized` | Kredensial salah | `{"message": "Email atau Password Salah"}` |
| `400 Bad Request` | Format JSON salah | `{"message": "Format JSON tidak valid"}` |
