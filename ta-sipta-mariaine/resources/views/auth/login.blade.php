<x-guest-layout>
    <!-- Logo -->
    <div class="flex justify-center mb-6 mt-6">
        <img src="{{ url('img/jkb.png') }}" alt="PNC Logo" class="w-20 h-20">
    </div>

    <!-- Session Status -->
    <x-auth-session-status class="mb-4" :status="session('status')" />
    
    @if(session('error'))
        <div class="mb-4 p-3 bg-red-100 text-red-700 rounded-md text-sm">
            {{ session('error') }}
        </div>
    @endif

    <h2 class="text-xl font-black text-center text-gray-700">SIPTA JKB PNC</h2>
    <p class="text-xs text-center text-gray-500 mb-6">Silakan login terlebih dahulu!</p>
    
    <!-- Google Login Button -->
    <div class="mb-4">
        <a href="{{ route('auth.google') }}" class="w-full flex items-center justify-center gap-3 h-12 bg-white border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50 transition shadow-sm">
            <svg class="w-5 h-5" viewBox="0 0 24 24">
                <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
            </svg>
            <span class="font-medium">Sign in with Google</span>
        </a>
    </div>

    <!-- Divider -->
    <div class="relative mb-4">
        <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-gray-300"></div>
        </div>
        <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-white text-gray-500">Atau login dengan email</span>
        </div>
    </div>

    <form method="POST" action="{{ route('login') }}">
        @csrf
        <!-- Email Address -->
        <div>
            <x-input-label for="email" :value="('Email')" />
            <x-text-input id="email" class="block mt-1 w-full h-10 p-3 border rounded-md focus:ring focus:ring-blue-300"
                type="email" name="email" :value="old('email')" required autofocus autocomplete="username" />
            <x-input-error :messages="$errors->get('email')" class="mt-2" />
        </div>

        <!-- Password -->
        <div class="mt-4">
            <x-input-label for="password" :value="('Password')" />
            <x-text-input id="password" class="block mt-1 h-10 w-full p-3 border rounded-md focus:ring focus:ring-blue-300"
                type="password" name="password" required autocomplete="current-password" />
            <x-input-error :messages="$errors->get('password')" class="mt-2" />
        </div>

        <!-- Remember Me -->
        <div class="flex items-center justify-between mt-4">
            <label for="remember_me" class="flex items-center">
                <input id="remember_me" type="checkbox" class="text-blue-600 border-gray-300 rounded focus:ring-blue-500" name="remember">
                <span class="ml-2 text-sm text-gray-600">{{ __('Remember me') }}</span>
            </label>

            @if (Route::has('password.request'))
                <a class="text-sm text-blue-600 hover:underline" href="{{ route('password.request') }}">
                    {{ __('Forgot your password?') }}
                </a>
            @endif
        </div>

        <div class="mt-6 mb-6">
            <button type="submit" class="w-full h-10 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition">
                {{ __('Log in') }}
            </button>
        </div>
    </form>
</x-guest-layout>
