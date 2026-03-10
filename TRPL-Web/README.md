# TRPL-Web - Project Showcase Platform

Platform showcase project mahasiswa TRPL dengan Google OAuth2 SSO (@pnc.ac.id). Desain mengikuti macOS design language yang clean dan minimalis.

## 🚀 Tech Stack

**Backend:**
- Golang (Gin Framework)
- SQLite (shared dengan sso-auth)
- GORM
- JWT Validation (OAuth handled by sso-auth)

**Frontend:**
- React 19
- Vite 7
- Tailwind CSS v4 (macOS-inspired design)
- React Router v7
- Axios

**Authentication:**
- Separate `sso-auth` service (Go + Gin + Google OAuth2)
- JWT tokens (24-hour validity)
- Domain restriction: @pnc.ac.id only

## 📁 Project Structure

```
TRPL-Web/
├── backend/              # Golang API server (port 8080)
│   ├── config/           # Database config (SQLite)
│   ├── controllers/      # Request handlers (Project CRUD)
│   ├── middleware/       # JWT validation middleware
│   ├── models/           # Database models (Project)
│   ├── routes/           # API routes
│   ├── utils/            # JWT validation helpers
│   ├── main.go           # Entry point
│   ├── .env.example      # Environment template
│   └── README.md         # Backend documentation
│
└── frontend/             # React app (port 5173)
    ├── src/
    │   ├── components/   # Navbar, ProtectedRoute
    │   ├── pages/        # Login, Dashboard, AuthCallback
    │   ├── services/     # API services (dual endpoints)
    │   ├── context/      # AuthContext (global state)
    │   ├── App.jsx       # Main app with routes
    │   ├── main.jsx      # Entry point
    │   └── index.css     # Tailwind v4 styles
    ├── .env.example      # Environment template
    └── README.md         # Frontend documentation

Note: Authentication logic ada di /home/potydev/SSO-PNC-TA/sso-auth (port 9090)
```

## ⚙️ Quick Setup

### Prerequisites
- Go 1.25+
- Node.js 18+
- SSO Auth Service running (port 9090)

### 1. SSO Auth Service (Run First!)

```bash
cd /home/potydev/SSO-PNC-TA/sso-auth

# Setup environment
cp .env.example .env
# Edit .env dengan Google OAuth credentials

# Install dependencies
go mod tidy

# Run service
go run .
```

SSO service akan jalan di `http://localhost:9090`

### 2. Backend Setup

```bash
cd backend

# Install dependencies
go mod download

# Setup environment
cp .env.example .env
# Edit: JWT_SECRET harus sama dengan sso-auth!
# Edit: DB_DATABASE path ke SQLite (shared dengan sso-auth)

# Run server
go run .
```

Backend akan jalan di `http://localhost:8080`

### 3. Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Setup environment
cp .env.example .env
# Edit jika perlu (default sudah benar)

# Run dev server
npm run dev
```

Frontend akan jalan di `http://localhost:5173`

### 4. Access Application

Buka browser: `http://localhost:5173`

**Login Flow:**
1. Klik "Sign in with Google"
2. Login dengan email @pnc.ac.id
3. Redirect ke Dashboard

## 🎯 Features

### Completed ✅
- **Authentication**: Google OAuth2 via SSO service (@pnc.ac.id only)
- **JWT Authorization**: Token-based with 24-hour validity
- **User Management**: Mahasiswa & Dosen profiles
- **Dashboard**: Statistics cards + latest projects grid
- **Protected Routes**: Role-based access control
- **macOS Design**: Clean, minimalis, whitespace maksimal
- **Responsive UI**: Mobile-friendly dengan Tailwind v4
- **Skeleton Loaders**: Smooth loading states

### In Progress 🚧
- Project CRUD operations
- Project detail page
- My Projects page (Mahasiswa only)

### Planned ⏳
- Image upload for project covers
- Project listing with filters & search
- Profile page synced with ta-sipta
- Pagination & infinite scroll
- View counter for projects
- Dark mode toggle

## 📡 API Endpoints

### SSO Auth Service (port 9090)
- `GET /api/auth/google` - Get Google OAuth URL
- `GET /api/auth/google/callback` - OAuth callback handler
- `GET /api/auth/me` - Get current user (protected)
- `POST /api/auth/verify` - Verify JWT token
- `POST /api/auth/logout` - Logout user

### TRPL Backend (port 8080)
- `GET /health` - Health check

**Projects (Public):**
- `GET /api/projects` - List published projects
- `GET /api/projects/:id` - Get project detail
- `GET /api/projects/stats` - Get statistics

**Projects (Protected):**
- `POST /api/projects` - Create project (Mahasiswa only)
- `PUT /api/projects/:id` - Update project (Owner only)
- `DELETE /api/projects/:id` - Delete project (Owner only)
- `GET /api/my-projects` - Get user's projects

## 🏗️ Architecture

```
┌──────────────────────────────────────────────────┐
│           Frontend (React + Vite)                │
│           http://localhost:5173                  │
└────────────┬──────────────────┬──────────────────┘
             │                  │
    Auth API │                  │ Project API
             ▼                  ▼
  ┌──────────────────┐   ┌──────────────────┐
  │  SSO Auth        │   │  TRPL Backend    │
  │  Port: 9090      │   │  Port: 8080      │
  │                  │   │                  │
  │ - Google OAuth   │   │ - Project CRUD   │
  │ - JWT Generate   │   │ - JWT Validate   │
  │ - User Mgmt      │   │ - Public API     │
  └────────┬─────────┘   └────────┬─────────┘
           │                      │
           └──────────┬───────────┘
                      │
                      ▼
            ┌──────────────────┐
            │ SQLite Database  │
            │    (Shared)      │
            └──────────────────┘
```

## 🔒 Authentication

**Restriction:** Only email addresses with `@pnc.ac.id` domain can login.

**Default Role:** New users are assigned **Mahasiswa** role automatically.

**JWT Token:** Valid for **24 hours**.

### Authentication Flow

```
1. User clicks "Login with Google"
   ↓
2. Frontend → GET /api/auth/google (SSO Service)
   ↓
3. Redirect to Google OAuth consent screen
   ↓
4. User approves → Google redirects to /auth/google/callback
   ↓
5. SSO Service validates email domain (@pnc.ac.id)
   ↓
6. SSO Service creates/updates user in SQLite
   ↓
7. SSO Service generates JWT token (signed with JWT_SECRET)
   ↓
8. Frontend receives token + user data
   ↓
9. Frontend stores token in localStorage
   ↓
10. Frontend requests protected routes → TRPL Backend
    ↓
11. TRPL Backend validates JWT (using same JWT_SECRET)
    ↓
12. Access granted to protected resources
```

**Critical:** `JWT_SECRET` must be identical in both SSO auth service and TRPL backend for token validation to work.

## 🗄️ Database Schema

**Technology:** SQLite (shared file)

**Location:** `backend/database/trpl_web.sqlite`

**Shared Access:** Both SSO auth service and TRPL backend use the same SQLite file.

### Tables Managed by SSO Auth Service
- `users` - User accounts (email, role, google_id, foto_profile)
- `mahasiswa` - Student profiles (NIM, tempat/tanggal lahir, jenis kelamin, prodi, tahun_ajaran)
- `dosen` - Lecturer profiles (NIDN, NIP, jabatan)
- `program_studi` - Study programs
- `tahun_ajaran` - Academic years

### Tables Managed by TRPL Backend
- `users` - User accounts (email, role, google_id, foto_profile)
- `mahasiswa` - Student profiles (NIM, tempat/tanggal lahir, jenis kelamin, prodi, tahun_ajaran)
- `dosen` - Lecturer profiles (NIDN, NIP, jabatan)
- `program_studi` - Study programs
- `tahun_ajaran` - Academic years
- `projects` - Project showcase (judul, deskripsi, kategori, teknologi, URL demo/repo)

## 🔧 Development

### Running Complete Stack

You need **3 terminal windows**:

**Terminal 1 - SSO Auth Service:**
```bash
cd SSO-PNC-TA/sso-auth
air  # or: go run main.go
```

**Terminal 2 - TRPL Backend:**
```bash
cd TRPL-Web/backend
air  # or: go run main.go
```

**Terminal 3 - Frontend:**
```bash
cd TRPL-Web/frontend
npm run dev
```

### Hot Reload Setup

```bash
go install github.com/air-verse/air@latest
```

Then use `air` command instead of `go run main.go` for auto-reload on file changes.

### Build for Production

**SSO Auth:**
```bash
cd SSO-PNC-TA/sso-auth
go build -o sso-auth main.go
./sso-auth
```

**Backend:**
```bash
cd TRPL-Web/backend
go build -o trpl-web-backend main.go
./trpl-web-backend
```

**Frontend:**
```bash
cd TRPL-Web/frontend
npm run build
npm run preview
```

## 📝 Environment Variables

### SSO Auth Service (.env)
```env
PORT=9090
JWT_SECRET=your-secret-key-must-match-backend

GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_REDIRECT_URL=http://localhost:9090/api/auth/google/callback
ALLOWED_DOMAIN=pnc.ac.id

DB_PATH=../TRPL-Web/backend/database/trpl_web.sqlite
FRONTEND_URL=http://localhost:5173
```

### TRPL Backend (.env)
```env
PORT=8080
JWT_SECRET=your-secret-key-must-match-backend

DB_PATH=./database/trpl_web.sqlite
FRONTEND_URL=http://localhost:5173
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080/api
VITE_AUTH_API_URL=http://localhost:9090/api
```

**Important Notes:**
- `JWT_SECRET` must be **identical** in both SSO auth and TRPL backend
- `DB_PATH` in sso-auth points to shared SQLite file in TRPL-Web/backend/database/
- Frontend has **two** API URLs: one for auth (9090), one for projects (8080)

## 🐛 Troubleshooting

### Service Not Running
**Check if all services are up:**
```bash
curl http://localhost:9090/health  # SSO Auth
curl http://localhost:8080/health  # TRPL Backend
curl http://localhost:5173         # Frontend
```

### CORS Error
Pastikan `FRONTEND_URL` di **kedua** service (.env) sesuai dengan URL frontend (`http://localhost:5173`).

### JWT Validation Failed / 401 Unauthorized
**Possible causes:**
1. Token expired (24 hours) → Logout and login again
2. `JWT_SECRET` mismatch → Ensure **identical** secret in both sso-auth and backend .env files
3. Token not sent → Check if `Authorization: Bearer <token>` header present

### Database Connection Failed
**Check SQLite file:**
```bash
ls -la TRPL-Web/backend/database/trpl_web.sqlite
```
- File must exist and be readable
- Both services need read/write permission
- Check `DB_PATH` in both .env files

### Login Redirect Fails
**Verify Google OAuth settings:**
1. `GOOGLE_REDIRECT_URL` must match Google Console callback URL
2. Must be `http://localhost:9090/api/auth/google/callback` (port 9090, not 8080)
3. Check if email domain is `@pnc.ac.id`

### Frontend Can't Connect to Backend
**Check environment variables:**
- `VITE_API_URL=http://localhost:8080/api` (TRPL Backend)
- `VITE_AUTH_API_URL=http://localhost:9090/api` (SSO Auth)
- Restart frontend after changing .env file

### Database Locked Error
Multiple services accessing SQLite simultaneously. This is normal and handled by SQLite's locking mechanism. If persistent:
```bash
fuser TRPL-Web/backend/database/trpl_web.sqlite  # Check which processes using file
```

## 📚 Documentation

- Backend API docs: `backend/README.md`
- Frontend docs: `frontend/README.md`

## 👥 Roles & Permissions

- **Mahasiswa**: Can create, edit, delete own projects
- **Dosen**: Can view all projects (read-only)

## 🚦 Next Steps

1. Implementasi project CRUD lengkap
2. Sinkronisasi data mahasiswa dari ta-sipta-mariaine
3. Upload gambar project ke storage
4. Search & filter functionality
5. Deploy to production

## 📄 License

MIT License
