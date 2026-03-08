# TRPL-Web Frontend

React frontend untuk TRPL Project Showcase dengan Vite, Tailwind CSS, dan Google OAuth2.

## 🚀 Tech Stack

- **Framework**: React 18 + Vite
- **Styling**: Tailwind CSS
- **Routing**: React Router v6
- **HTTP Client**: Axios
- **State Management**: React Context API

## ⚙️ Setup Instructions

### 1. Install Dependencies

```bash
cd frontend
npm install
```

### 2. Configure Environment

```bash
cp .env.example .env
```

Edit `.env`:
```env
VITE_API_URL=http://localhost:8080/api
```

### 3. Run Development Server

```bash
npm run dev
```

Frontend akan jalan di `http://localhost:5173`

## 🔐 Authentication Flow

1. User clicks "Sign in with Google"
2. Frontend calls `/api/auth/google` to get OAuth URL
3. User redirects to Google login
4. After login, Google redirects to backend callback
5. Backend generates JWT and redirects to `/auth/callback?token=xxx`
6. Frontend saves token to localStorage
7. All subsequent API calls include `Authorization: Bearer <token>` header

## 🛠️ Available Scripts

```bash
npm run dev          # Start development server
npm run build        # Build for production
npm run preview      # Preview production build
```
