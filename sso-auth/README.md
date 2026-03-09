# SSO Auth Service

Service autentikasi terpisah untuk semua aplikasi di workspace (`TRPL-Web`, `ta-sipta-mariaine`, dll).

## Fitur
- Google OAuth2 login (`@pnc.ac.id`)
- JWT token issuance
- Endpoint `me` dan `verify`
- SQLite (default share file dengan TRPL backend untuk fase transisi)

## Endpoint
- `GET /health`
- `GET /api/auth/google`
- `GET /api/auth/google/callback`
- `GET /api/auth/me` (Bearer token)
- `GET /api/auth/verify` (Bearer token)
- `POST /api/auth/logout` (Bearer token)

## Menjalankan
```bash
cd /home/potydev/SSO-PNC-TA/sso-auth
cp .env.example .env
go mod tidy
go run .
```

Server default: `http://localhost:9090`

## Catatan Integrasi
- Set `GOOGLE_REDIRECT_URL` ke `http://localhost:9090/api/auth/google/callback`
- Set frontend `VITE_AUTH_API_URL=http://localhost:9090/api`
- Set backend TRPL dan `sso-auth` memakai `JWT_SECRET` yang sama
