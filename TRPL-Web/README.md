# TRPL-Web - Project Showcase Platform

Platform showcase project mahasiswa TRPL dengan SSO Google OAuth2 (@pnc.ac.id).

## 🚀 Tech Stack

**Backend:**
- Golang (Gin Framework)
- MySQL
- GORM
- JWT Authentication
- Google OAuth2

**Frontend:**
- React 18
- Vite
- Tailwind CSS
- React Router
- Axios

## 📁 Project Structure

```
TRPL-Web/
├── backend/              # Golang API server
│   ├── config/           # Database config
│   ├── controllers/      # Request handlers
│   ├── middleware/       # Auth middleware
│   ├── models/           # Database models
│   ├── routes/           # API routes
│   ├── utils/            # JWT helpers
│   ├── main.go           # Entry point
│   ├── .env.example      # Environment template
│   └── README.md         # Backend docs
│
└── frontend/             # React app
    ├── src/
    │   ├── components/   # UI components
    │   ├── pages/        # Page components
    │   ├── services/     # API services
    │   ├── context/      # Auth context
    │   └── App.jsx       # Main app
    ├── .env.example      # Environment template
    └── README.md         # Frontend docs
```

## ⚙️ Quick Setup

### 1. Backend Setup

```bash
cd backend

# Install dependencies
go mod download

# Setup environment
cp .env.example .env
# Edit .env with your credentials

# Create database
mysql -u root -p
CREATE DATABASE trpl_web CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EXIT;

# Run server
go run main.go
```

Backend will run on `http://localhost:8080`

### 2. Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Setup environment
cp .env.example .env
# Edit .env if needed

# Run dev server
npm run dev
```

Frontend will run on `http://localhost:5173`

### 3. Google OAuth2 Setup

1. Buka [Google Cloud Console](https://console.cloud.google.com/)
2. Buat project baru
3. Enable **Google+ API**
4. Credentials → Create OAuth 2.0 Client ID
5. Authorized redirect URIs: `http://localhost:8080/api/auth/google/callback`
6. Copy **Client ID** dan **Client Secret** ke backend `.env`

## 🎯 Features

### Completed ✅
- Google OAuth2 authentication (@pnc.ac.id only)
- JWT-based authorization
- User management (Mahasiswa & Dosen)
- Dashboard with statistics
- Protected routes
- Responsive UI with Tailwind CSS

### TODO ⏳
- Project CRUD (Create, Read, Update, Delete)
- Project listing with filters & search
- Project detail page with view counter
- Image upload for project covers
- My Projects page (Mahasiswa only)
- Profile page with data from ta-sipta
- Pagination & infinite scroll
- Dark mode toggle

## 📡 API Endpoints

### Auth
- `GET /api/auth/google` - Get Google OAuth URL
- `GET /api/auth/google/callback` - OAuth callback
- `GET /api/auth/me` - Get current user (protected)
- `POST /api/auth/logout` - Logout

### Projects
- `GET /api/projects` - List published projects
- `GET /api/projects/:id` - Get project detail
- `GET /api/projects/stats` - Get statistics
- `POST /api/projects` - Create project (Mahasiswa)
- `PUT /api/projects/:id` - Update project (Owner)
- `DELETE /api/projects/:id` - Delete project (Owner)
- `GET /api/my-projects` - Get user's projects (Protected)

## 🔒 Authentication

Hanya email dengan domain **@pnc.ac.id** yang bisa login.

Default role untuk user baru adalah **Mahasiswa**.

JWT token valid selama **24 jam**.

## 🗄️ Database Schema

**Tables:**
- `users` - User accounts (email, role, google_id, foto_profile)
- `mahasiswa` - Student profiles (NIM, tempat/tanggal lahir, jenis kelamin, prodi, tahun_ajaran)
- `dosen` - Lecturer profiles (NIDN, NIP, jabatan)
- `program_studi` - Study programs
- `tahun_ajaran` - Academic years
- `projects` - Project showcase (judul, deskripsi, kategori, teknologi, URL demo/repo)

## 🔧 Development

### Run Backend with Hot Reload

```bash
go install github.com/cosmtrek/air@latest
air
```

### Run Frontend Dev Server

```bash
npm run dev
```

### Build for Production

**Backend:**
```bash
go build -o trpl-web-backend main.go
./trpl-web-backend
```

**Frontend:**
```bash
npm run build
npm run preview
```

## 📝 Environment Variables

**Backend (.env):**
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=trpl_web

PORT=8080
JWT_SECRET=your-secret-key

GOOGLE_CLIENT_ID=your-client-id
GOOGLE_CLIENT_SECRET=your-client-secret
GOOGLE_REDIRECT_URL=http://localhost:8080/api/auth/google/callback
ALLOWED_DOMAIN=pnc.ac.id

FRONTEND_URL=http://localhost:5173
```

**Frontend (.env):**
```env
VITE_API_URL=http://localhost:8080/api
```

## 🐛 Troubleshooting

### CORS Error
Pastikan `FRONTEND_URL` di backend `.env` sesuai dengan URL frontend.

### Database Connection Failed
Cek credentials database di `.env` dan pastikan MySQL service running.

### 401 Unauthorized
Token expired. Logout dan login kembali.

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
