# TRPL-Web Backend

Backend API untuk TRPL Project Showcase dengan Golang (Gin Framework). 

**⚠️ Authentication sudah dipindahkan ke service terpisah:** `sso-auth` (lihat `/home/potydev/SSO-PNC-TA/sso-auth`)

## 🚀 Tech Stack

- **Language**: Go 1.25+
- **Framework**: Gin
- **Database**: SQLite (shared dengan sso-auth)
- **ORM**: GORM
- **Authentication**: JWT validation only (SSO service handles OAuth)

## 📁 Project Structure

```
backend/
├── config/          # Database configuration
├── controllers/     # Request handlers
├── middleware/      # Auth & role middleware
├── models/          # Database models
├── routes/          # API routes
├── utils/           # JWT & helpers
├── migrations/      # SQL migrations (optional)
├── main.go          # Entry point
├── .env.example     # Environment template
└── go.mod           # Dependencies
```

## ⚙️ Setup Instructions

### 1. Install Dependencies

```bash
cd TRPL-Web/backend
go mod download
```

### 2. Configure Environment

```bash
cp .env.example .env
```

Edit `.env` dan sesuaikan:
- `DB_DATABASE`: Path ke SQLite database (shared dengan sso-auth)
- `JWT_SECRET`: Harus sama dengan sso-auth untuk validasi token
- `PORT`: Default 8080

**Penting**: JWT_SECRET harus identik dengan sso-auth agar token validation berfungsi!

### 3. Setup Database

Database SQLite otomatis dibuat saat pertama kali run. Pastikan path `DB_DATABASE` benar dan writable.

### 4. Run Application

```bash
go run .
```

Server akan jalan di `http://localhost:8080`

## 📡 API Endpoints

### Health Check
- `GET /health` - Server health status

### Projects (Public)
- `GET /api/projects` - List all published projects
- `GET /api/projects/:id` - Get project detail
- `GET /api/projects/stats` - Get statistics

### Projects (Protected - requires JWT token)
- `POST /api/projects` - Create project (Mahasiswa only)
- `PUT /api/projects/:id` - Update project (Owner only)
- `DELETE /api/projects/:id` - Delete project (Owner only)
- `GET /api/my-projects` - Get user's projects

**Note**: Authentication endpoints (`/api/auth/*`) sudah dipindahkan ke `sso-auth` service (port 9090)

## 🔒 Authentication Flow

**Service ini TIDAK handle OAuth**, hanya validasi JWT token dari `sso-auth` service.

1. Frontend redirect user ke `sso-auth` service (port 9090)
2. SSO service handle Google OAuth dan generate JWT token
3. Frontend menyimpan token
4. Request ke backend ini harus include header:
   ```
   Authorization: Bearer <jwt_token>
   ```
5. Middleware `AuthMiddleware` validasi token menggunakan JWT_SECRET yang sama

**Architecture**:
```
Frontend (5173) → SSO-Auth (9090) [OAuth + JWT] 
                ↓
                → TRPL Backend (8080) [JWT Validation only]
```

## 🗄️ Database Models

**Shared SQLite database dengan sso-auth service:**

- **users**: User accounts (dikelola oleh sso-auth)
- **mahasiswa**: Student profiles (dikelola oleh sso-auth)
- **dosen**: Lecturer profiles (dikelola oleh sso-auth)
- **program_studi**: Study programs (dikelola oleh sso-auth)
- **tahun_ajaran**: Academic years (dikelola oleh sso-auth)
- **projects**: Project showcase (dikelola oleh service ini)

## 🏗️ Service Architecture

```
┌─────────────────────────────────────────────────────────┐
│                      Frontend (React)                    │
│                    http://localhost:5173                 │
└────────────────┬────────────────────────┬────────────────┘
                 │                        │
        Auth API │                        │ Project API
                 ▼                        ▼
    ┌────────────────────┐   ┌──────────────────────┐
    │   SSO-Auth Service │   │   TRPL-Web Backend   │
    │   Port: 9090       │   │   Port: 8080         │
    │                    │   │                      │
    │  - Google OAuth    │   │  - Project CRUD      │
    │  - JWT Generate    │   │  - JWT Validate      │
    │  - User Management │   │  - Public API        │
    └─────────┬──────────┘   └──────────┬───────────┘
              │                         │
              └─────────┬───────────────┘
                        │
                        ▼
              ┌──────────────────┐
              │  SQLite Database │
              │  (Shared)        │
              └──────────────────┘
```

## 🔧 Development

### Run dengan hot reload

```bash
go install github.com/air-verse/air@latest
air
```

### Build for production

```bash
go build -o trpl-web-backend .
./trpl-web-backend
```

### Running Complete Stack

**Terminal 1 - SSO Auth Service:**
```bash
cd /path/to/sso-auth
go run .
```

**Terminal 2 - TRPL Backend (this service):**
```bash
cd /path/to/TRPL-Web/backend
go run .
```

**Terminal 3 - Frontend:**
```bash
cd /path/to/TRPL-Web/frontend
npm run dev
```

Access: `http://localhost:5173`

## 📝 Notes

- Service ini TIDAK handle authentication, hanya validasi JWT
- Database SQLite di-share dengan sso-auth service
- JWT_SECRET harus identik dengan sso-auth
- Project hanya bisa dibuat oleh Mahasiswa (role check via JWT claims)
- JWT token valid selama 24 jam (diatur oleh sso-auth)
- Gunakan middleware `AuthMiddleware` untuk protected routes
