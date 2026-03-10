# Google OAuth Setup untuk TA-SIPTA-MARIAINE

## Error yang Muncul
```
Error 400: redirect_uri_mismatch
```

## Penyebab
Redirect URI yang dikirim aplikasi (`http://localhost:8000/auth/google/callback`) **belum terdaftar** di Google Cloud Console.

## Solusi - Tambahkan Redirect URI di Google Cloud Console

### Langkah 1: Buka Google Cloud Console
1. Buka: https://console.cloud.google.com/apis/credentials
2. Login dengan akun Google yang digunakan untuk membuat project

### Langkah 2: Edit OAuth 2.0 Client ID
1. Cari OAuth 2.0 Client yang sedang digunakan (lihat Client ID di file `.env`)
2. Klik pada nama client untuk edit

### Langkah 3: Tambahkan Authorized Redirect URIs
Di bagian **"Authorized redirect URIs"**, tambahkan:

**Untuk Development:**
```
http://localhost:8000/auth/google/callback
```

**Saat ini yang terdaftar (jangan dihapus):**
```
http://localhost:9090/api/auth/google/callback  (untuk sso-auth)
```

**Setelah ditambahkan, akan ada 2 redirect URIs:**
- `http://localhost:9090/api/auth/google/callback` - untuk SSO Auth Service
- `http://localhost:8000/auth/google/callback` - untuk TA-SIPTA-MARIAINE

### Langkah 4: Save dan Tunggu
1. Klik **SAVE**
2. Tunggu 1-2 menit agar perubahan ter-propagate
3. Coba login lagi

## Verifikasi Konfigurasi Laravel

Pastikan `.env` sudah benar:
```env
GOOGLE_CLIENT_ID=your-google-client-id.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=GOCSPX-your-client-secret
GOOGLE_REDIRECT_URL=http://localhost:8000/auth/google/callback
ALLOWED_DOMAIN=pnc.ac.id
```

## Test Login
1. Buka: http://localhost:8000/login
2. Klik "Sign in with Google"
3. Pilih akun dengan email `@pnc.ac.id`
4. Seharusnya berhasil redirect ke dashboard

## Troubleshooting

### Jika masih error setelah setup:
```bash
# Clear cache
php artisan config:clear
php artisan cache:clear
php artisan route:clear

# Restart server
# Ctrl+C di terminal
php artisan serve
```

### Jika muncul error "Token Mismatch":
```bash
php artisan key:generate
```

### Jika profil mahasiswa tidak dibuat otomatis:
Cek di database, user baru otomatis dibuat dengan:
- Role: `Mahasiswa`
- Email dari Google
- Foto profil dari Google
- Profil mahasiswa dengan NIM kosong (akan diisi manual nanti)
