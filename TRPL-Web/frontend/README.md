# TRPL-Web Frontend

Frontend aplikasi TRPL Project Showcase dengan React, Vite, dan Tailwind CSS v4. Desain mengikuti macOS design language (minimalis, clean, whitespace maksimal).

## 🚀 Tech Stack

- **Framework**: React 19
- **Build Tool**: Vite 7
- **Styling**: Tailwind CSS v4 (dengan `@tailwindcss/postcss`)
- **Router**: React Router v7
- **HTTP Client**: Axios
- **State Management**: Context API (AuthContext)

## 📁 Project Structure

```
frontend/
├── src/
│   ├── components/       # Reusable components
│   │   ├── Navbar.jsx
│   │   └── ProtectedRoute.jsx
│   ├── context/          # Global state
│   │   └── AuthContext.jsx
│   ├── pages/            # Route pages
│   │   ├── Login.jsx
│   │   ├── Dashboard.jsx
│   │   └── AuthCallback.jsx
│   ├── services/         # API services
│   │   └── projectService.js
│   ├── App.jsx           # Main app component
│   ├── main.jsx          # Entry point
│   └── index.css         # Global styles (Tailwind v4)
├── public/               # Static assets
├── .env.example          # Environment template
├── tailwind.config.js    # Tailwind configuration
├── postcss.config.js     # PostCSS with Tailwind plugin
├── vite.config.js        # Vite configuration
└── package.json          # Dependencies
```

## ⚙️ Setup Instructions

### 1. Install Dependencies

```bash
cd TRPL-Web/frontend
npm install
```

### 2. Configure Environment

```bash
cp .env.example .env
```

Edit `.env`:
```env
VITE_API_URL=http://localhost:8080/api          # TRPL Backend
VITE_AUTH_API_URL=http://localhost:9090/api     # SSO Auth Service
```

### 3. Run Development Server

```bash
npm run dev
```

Aplikasi akan berjalan di `http://localhost:5173`

### 4. Build for Production

```bash
npm run build
```

Output ada di folder `dist/`

## 🎨 Design System (macOS-inspired)

### Color Palette
- **Background**: `#f5f5f7` (Apple gray)
- **Cards**: `#ffffff` dengan border `gray-200/60`
- **Text**: `gray-900` (primary), `gray-600` (secondary)
- **Accent**: `blue-600` (buttons, links)

### Typography
- **Font Sizes**: `13px`, `15px`, `17px`, `22px`, `32px`, `40px`, `48px`
- **Font Weight**: `font-medium` (500), `font-semibold` (600)
- **Tracking**: `tracking-tight` untuk heading

### Spacing
- **Padding**: `px-8` (32px) untuk container
- **Gap**: `gap-6` (24px) untuk grid
- **Margin**: `mb-16` (64px) untuk section spacing

### Components
- **Border Radius**: `rounded-2xl` (18px), `rounded-xl` (12px)
- **Borders**: `border-gray-200/60` (subtle)
- **Shadows**: `shadow-sm` only (minimal)
- **Transitions**: `transition-colors`, `duration-200`

## 🔐 Authentication Flow

```
1. User buka /login
   ↓
2. Klik "Sign in with Google"
   ↓
3. Frontend hit: GET http://localhost:9090/api/auth/google
   ↓
4. Redirect ke Google OAuth (sso-auth service)
   ↓
5. User login dengan @pnc.ac.id
   ↓
6. Google callback ke: http://localhost:9090/api/auth/google/callback
   ↓
7. SSO service generate JWT token
   ↓
8. Redirect ke: http://localhost:5173/auth/callback?token=<jwt>
   ↓
9. AuthCallback.jsx extract token → save to localStorage
   ↓
10. Redirect ke /dashboard (authenticated)
```

## 📡 API Integration

### Auth Service (port 9090)
```javascript
// services/projectService.js - authApi instance
authService.getGoogleAuthURL()    // Get OAuth URL
authService.getCurrentUser()      // Get user profile
```

### Backend Service (port 8080)
```javascript
// services/projectService.js - api instance
projectService.getProjects()      // List projects
projectService.getProjectByID(id) // Get detail
projectService.getStats()         // Get statistics
projectService.createProject(data)// Create (auth required)
projectService.updateProject(id, data)
projectService.deleteProject(id)
projectService.getMyProjects()
```

### Axios Interceptors
- **Request**: Auto-attach JWT token dari localStorage
- **Response**: Handle 401 → redirect ke login

## 🧩 Key Components

### AuthContext
Provides global auth state:
- `user`: Current user object
- `loading`: Loading state
- `login(token)`: Save token & fetch user
- `logout()`: Clear token & user

### ProtectedRoute
Wrapper untuk routes yang butuh auth. Redirect ke `/login` jika tidak authenticated.

### Navbar
- Logo + navigation links
- User dropdown dengan avatar
- Active route highlighting
- Mobile responsive

### Pages

**Login.jsx**
- macOS-style design
- Google OAuth button
- Loading state saat redirect

**Dashboard.jsx**
- Hero section dengan greeting
- Stats cards (Projects, Students, Categories)
- Latest projects grid
- Quick actions (untuk Mahasiswa)
- Skeleton loaders

**AuthCallback.jsx**
- Extract token dari URL query
- Save to localStorage
- Redirect ke dashboard

## 🔧 Development

### Hot Module Replacement (HMR)
Vite auto-reload saat file berubah

### Lint
```bash
npm run lint
```

### Preview Production Build
```bash
npm run preview
```

## 🏗️ Complete Stack

### Running All Services

**Terminal 1 - SSO Auth:**
```bash
cd /path/to/sso-auth
go run .
```

**Terminal 2 - Backend:**
```bash
cd /path/to/TRPL-Web/backend
go run .
```

**Terminal 3 - Frontend:**
```bash
cd /path/to/TRPL-Web/frontend
npm run dev
```

**Access**: `http://localhost:5173`

## 📝 Notes

- **Tailwind v4**: Menggunakan `@import "tailwindcss"` syntax (bukan `@tailwind` directives)
- **CSS Reset**: Body default Vite sudah dihapus agar layout tidak terpusat
- **Token Storage**: JWT disimpan di `localStorage` dengan key `token`
- **Protected Routes**: Gunakan `<ProtectedRoute>` wrapper
- **API Separation**: Auth API (9090) terpisah dari Project API (8080)
- **Domain Restriction**: Hanya email `@pnc.ac.id` yang bisa login (enforced di backend)
