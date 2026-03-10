<?php

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use App\Models\User;
use App\Models\Mahasiswa;
use App\Models\Dosen;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Laravel\Socialite\Facades\Socialite;

class GoogleAuthController extends Controller
{
    /**
     * Redirect to Google OAuth
     */
    public function redirectToGoogle()
    {
        return Socialite::driver('google')
            ->redirectUrl(config('services.google.redirect'))
            ->with(['hd' => env('ALLOWED_DOMAIN', 'pnc.ac.id')])
            ->redirect();
    }

    /**
     * Handle Google OAuth callback
     */
    public function handleGoogleCallback()
    {
        try {
            $googleUser = Socialite::driver('google')->user();
            
            // Validasi domain email
            $allowedDomain = env('ALLOWED_DOMAIN', 'pnc.ac.id');
            $email = $googleUser->getEmail();
            
            if (!str_ends_with($email, '@' . $allowedDomain)) {
                return redirect()->route('login')->with('error', 'Hanya email dengan domain @' . $allowedDomain . ' yang diperbolehkan.');
            }

            // Cari atau buat user
            $user = User::where('email', $email)->first();

            if (!$user) {
                // Buat user baru dengan default role Mahasiswa
                $user = User::create([
                    'name' => $googleUser->getName(),
                    'email' => $email,
                    'password' => Hash::make(uniqid()), // Random password (tidak digunakan)
                    'role' => 'Mahasiswa',
                    'foto_profile' => $googleUser->getAvatar(),
                    'email_verified_at' => now(),
                ]);

                // Buat profil mahasiswa default
                Mahasiswa::create([
                    'user_id' => $user->id,
                    'nim' => '', // Akan diisi manual nanti
                    'nama' => $googleUser->getName(),
                    'tempat_lahir' => '',
                    'tanggal_lahir' => null,
                    'jenis_kelamin' => '',
                    'program_studi_id' => null,
                    'tahun_ajaran_id' => null,
                ]);
            } else {
                // Update foto profil jika berubah
                if ($user->foto_profile !== $googleUser->getAvatar()) {
                    $user->update([
                        'foto_profile' => $googleUser->getAvatar(),
                    ]);
                }
            }

            // Login user
            Auth::login($user, true);

            // Simpan waktu login terakhir
            session()->put('last_activity', now()->timestamp);

            return redirect()->intended(route('dashboard'));

        } catch (\Exception $e) {
            return redirect()->route('login')->with('error', 'Login gagal: ' . $e->getMessage());
        }
    }
}
