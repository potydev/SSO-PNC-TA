# TRPL-Web Backend

Backend API untuk TRPL Project Showcase dengan Golang (Gin Framework) dan Google OAuth2 authentication.

## 🚀 Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin
- **Database**: MySQL
- **ORM**: GORM
- **Authentication**: Google OAuth2 + JWT

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
cd backend
go mod download
```

### 2. Configure Environment

```bash
cp .env.example .env
```

Edit `.env` dan sesuaikan:
- Database credentials
- Google OAuth2 credentials (dari Google Cloud Console)
- JWT secret

### 3. Setup Google OAuth2

1. Buka [Google Cloud Console](https://console.cloud.google.com/)
2. Buat project baru atau pilih existing
3. Enable **Google+ API**
4. Credentials → Create OAuth 2.0 Client ID
5. Authorized redirect URIs: `http://localhost:8080/api/auth/google/callback`
6. Copy **Client ID** dan **Client Secret** ke `.env`

### 4. Create Database

```bash
mysql -u root -p
CREATE DATABASE trpl_web CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EXIT;
```

### 5. Run Application

```bash
go run main.go
```

Server akan jalan di `http://localhost:8080`

## 📡 API Endpoints

### Authentication
- `GET /api/auth/google` - Redirect ke Google OAuth
- `GET /api/auth/google/callback` - OAuth callback
- `GET /api/auth/me` - Get current user (protected)
- `POST /api/auth/logout` - Logout

### Projects
- `GET /api/projects` - List all published projects
- `GET /api/projects/:id` - Get project detail
- `GET /api/projects/stats` - Get statistics
- `POST /api/projects` - Create project (Mahasiswa only)
- `PUT /api/projects/:id` - Update project (Owner only)
- `DELETE /api/projects/:id` - Delete project (Owner only)
- `GET /api/my-projects` - Get user's projects (Protected)

## 🔒 Authentication Flow

1. Frontend redirect user ke `/api/auth/google`
2. User login dengan Google (@pnc.ac.id)
3. Callback ke `/api/auth/google/callback`
4. Backend generate JWT token
5. Redirect ke frontend dengan token
6. Frontend simpan token di localStorage
7. Setiap request ke protected endpoints, kirim header:
   ```
   Authorization: Bearer <jwt_token>
   ```

## 🗄️ Database Models

- **users**: User accounts (email, role, google_id)
- **mahasiswa**: Student profiles (NIM, tempat/tgl lahir, jenis kelamin, etc)
- **dosen**: Lecturer profiles (NIDN, NIP, etc)
- **program_studi**: Study programs
- **tahun_ajaran**: Academic years
- **projects**: Project showcase

## 🔧 Development

### Hot reload (optional)

```bash
go install github.com/cosmtrek/air@latest
air
```

### Build for production

```bash
go build -o trpl-web-backend main.go
./trpl-web-backend
```

## 📝 Notes

- Hanya email dengan domain `@pnc.ac.id` yang bisa login
- Default role untuk user baru adalah `Mahasiswa`
- Project hanya bisa dibuat oleh Mahasiswa
- JWT token valid selama 24 jam
