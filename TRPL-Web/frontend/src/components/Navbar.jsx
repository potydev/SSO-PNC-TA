import React from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export default function Navbar() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();
  const [isOpen, setIsOpen] = React.useState(false);

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  if (!user) return null;

  const getInitials = (name) => {
    return name
      .split(' ')
      .map((n) => n[0])
      .join('')
      .toUpperCase()
      .slice(0, 2);
  };

  const isActive = (path) => location.pathname === path;

  return (
    <nav className="bg-white/80 backdrop-blur-xl border-b border-gray-200/80 sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center gap-12">
            <Link to="/dashboard" className="flex items-center gap-3 group">
              <div className="w-8 h-8 bg-gray-900 rounded-lg flex items-center justify-center">
                <svg className="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth="2">
                  <path strokeLinecap="round" strokeLinejoin="round" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                </svg>
              </div>
              <span className="text-[17px] font-semibold text-gray-900">TRPL Web</span>
            </Link>
            
            <div className="hidden md:flex items-center gap-1">
              <Link
                to="/dashboard"
                className={`px-4 py-2 rounded-lg text-[15px] font-medium transition-colors ${
                  isActive('/dashboard')
                    ? 'bg-gray-100 text-gray-900'
                    : 'text-gray-600 hover:text-gray-900'
                }`}
              >
                Dashboard
              </Link>
              <Link
                to="/projects"
                className={`px-4 py-2 rounded-lg text-[15px] font-medium transition-colors ${
                  isActive('/projects')
                    ? 'bg-gray-100 text-gray-900'
                    : 'text-gray-600 hover:text-gray-900'
                }`}
              >
                Projects
              </Link>
              {user.role === 'Mahasiswa' && (
                <Link
                  to="/my-projects"
                  className={`px-4 py-2 rounded-lg text-[15px] font-medium transition-colors ${
                    isActive('/my-projects')
                      ? 'bg-gray-100 text-gray-900'
                      : 'text-gray-600 hover:text-gray-900'
                  }`}
                >
                  My Projects
                </Link>
              )}
            </div>
          </div>

          <div className="flex items-center">
            <div className="relative">
              <button
                onClick={() => setIsOpen(!isOpen)}
                className="flex items-center gap-3 hover:opacity-80 transition-opacity"
              >
                {user.foto_profile ? (
                  <img
                    src={user.foto_profile}
                    alt={user.name}
                    className="w-8 h-8 rounded-full"
                  />
                ) : (
                  <div className="w-8 h-8 rounded-full bg-gray-900 text-white flex items-center justify-center text-[13px] font-semibold">
                    {getInitials(user.name)}
                  </div>
                )}
                <span className="hidden md:block text-[15px] font-medium text-gray-900">{user.name}</span>
                <svg
                  className={`w-4 h-4 text-gray-500 transition-transform ${
                    isOpen ? 'rotate-180' : ''
                  }`}
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  strokeWidth="2"
                >
                  <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7"/>
                </svg>
              </button>

              {isOpen && (
                <>
                  <div 
                    className="fixed inset-0 z-10" 
                    onClick={() => setIsOpen(false)}
                  ></div>
                  <div className="absolute right-0 mt-3 w-72 bg-white/95 backdrop-blur-xl rounded-2xl shadow-lg border border-gray-200/80 py-2 z-20">
                    <div className="px-4 py-3 border-b border-gray-100">
                      <p className="text-[15px] font-semibold text-gray-900">{user.name}</p>
                      <p className="text-[13px] text-gray-500 mt-1">{user.email}</p>
                      <span className="inline-block mt-2 px-3 py-1 bg-gray-100 rounded-full text-[12px] font-medium text-gray-700">
                        {user.role}
                      </span>
                    </div>
                    
                    <div className="md:hidden py-2 border-b border-gray-100">
                      <Link
                        to="/dashboard"
                        className="block px-4 py-2 text-[15px] text-gray-900 hover:bg-gray-50 font-medium"
                        onClick={() => setIsOpen(false)}
                      >
                        Dashboard
                      </Link>
                      <Link
                        to="/projects"
                        className="block px-4 py-2 text-[15px] text-gray-900 hover:bg-gray-50 font-medium"
                        onClick={() => setIsOpen(false)}
                      >
                        Projects
                      </Link>
                      {user.role === 'Mahasiswa' && (
                        <Link
                          to="/my-projects"
                          className="block px-4 py-2 text-[15px] text-gray-900 hover:bg-gray-50 font-medium"
                          onClick={() => setIsOpen(false)}
                        >
                          My Projects
                        </Link>
                      )}
                    </div>

                    <button
                      onClick={handleLogout}
                      className="w-full text-left px-4 py-3 text-[15px] text-red-600 hover:bg-red-50 font-medium transition-colors"
                    >
                      Sign Out
                    </button>
                  </div>
                </>
              )}
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
}
