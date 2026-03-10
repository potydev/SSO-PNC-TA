import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { projectService } from '../services/projectService';
import { useAuth } from '../context/AuthContext';

export default function Dashboard() {
  const { user } = useAuth();
  const [stats, setStats] = useState(null);
  const [projects, setProjects] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const [statsData, projectsData] = await Promise.all([
        projectService.getStats(),
        projectService.getProjects(),
      ]);
      setStats(statsData);
      setProjects(projectsData.slice(0, 6));
    } catch (error) {
      console.error('Failed to fetch data:', error);
    } finally {
      setLoading(false);
    }
  };

  const StatsSkeleton = () => (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
      {[1, 2, 3].map((i) => (
        <div key={i} className="bg-white rounded-2xl border border-gray-200/60 p-8 animate-pulse">
          <div className="h-4 bg-gray-200 rounded w-24 mb-4"></div>
          <div className="h-10 bg-gray-300 rounded w-20"></div>
        </div>
      ))}
    </div>
  );

  const ProjectsSkeleton = () => (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {[1, 2, 3].map((i) => (
        <div key={i} className="bg-white rounded-2xl border border-gray-200/60 overflow-hidden animate-pulse">
          <div className="w-full h-56 bg-gray-200"></div>
          <div className="p-6">
            <div className="h-6 bg-gray-300 rounded w-3/4 mb-3"></div>
            <div className="h-4 bg-gray-200 rounded w-full mb-2"></div>
            <div className="h-4 bg-gray-200 rounded w-2/3"></div>
          </div>
        </div>
      ))}
    </div>
  );

  return (
    <div className="min-h-screen bg-[#f5f5f7]">
      <div className="bg-white border-b border-gray-200/60">
        <div className="max-w-7xl mx-auto px-8 py-16">
          <h1 className="text-[48px] font-semibold text-gray-900 tracking-tight mb-3">
            Welcome back, {user?.name?.split(' ')[0]}
          </h1>
          <p className="text-[17px] text-gray-600 font-normal">
            {user?.role === 'Mahasiswa' ? 'Showcase your best work' : 'Explore student innovations'}
          </p>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-8 py-12">
        {loading ? (
          <StatsSkeleton />
        ) : stats ? (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-16">
            <div className="bg-white rounded-2xl border border-gray-200/60 p-8 hover:border-gray-300 transition-colors">
              <p className="text-[13px] font-medium text-gray-500 uppercase tracking-wide mb-2">Total Projects</p>
              <p className="text-[40px] font-semibold text-gray-900 tracking-tight">{stats.total_projects}</p>
            </div>
            <div className="bg-white rounded-2xl border border-gray-200/60 p-8 hover:border-gray-300 transition-colors">
              <p className="text-[13px] font-medium text-gray-500 uppercase tracking-wide mb-2">Students</p>
              <p className="text-[40px] font-semibold text-gray-900 tracking-tight">{stats.total_mahasiswa}</p>
            </div>
            <div className="bg-white rounded-2xl border border-gray-200/60 p-8 hover:border-gray-300 transition-colors">
              <p className="text-[13px] font-medium text-gray-500 uppercase tracking-wide mb-2">Categories</p>
              <p className="text-[40px] font-semibold text-gray-900 tracking-tight">{stats.categories?.length || 0}</p>
            </div>
          </div>
        ) : null}

        <div className="mb-16">
          <div className="flex items-center justify-between mb-8">
            <h2 className="text-[32px] font-semibold text-gray-900 tracking-tight">Latest Projects</h2>
            <Link to="/projects" className="text-[15px] font-medium text-blue-600 hover:text-blue-700 transition-colors">
              View All
            </Link>
          </div>

          {loading ? (
            <ProjectsSkeleton />
          ) : projects.length > 0 ? (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {projects.map((project) => (
                <Link key={project.id} to={`/projects/${project.id}`} className="group bg-white rounded-2xl border border-gray-200/60 overflow-hidden hover:border-gray-300 hover:shadow-sm transition-all">
                  {project.url_gambar ? (
                    <div className="relative overflow-hidden h-56 bg-gray-100">
                      <img src={project.url_gambar} alt={project.judul} className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" />
                    </div>
                  ) : (
                    <div className="h-56 bg-gray-100 flex items-center justify-center">
                      <svg className="w-12 h-12 text-gray-300" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth="1.5">
                        <path strokeLinecap="round" strokeLinejoin="round" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                      </svg>
                    </div>
                  )}
                  <div className="p-6">
                    <h3 className="text-[17px] font-semibold text-gray-900 mb-2 line-clamp-1">{project.judul}</h3>
                    <p className="text-[15px] text-gray-600 mb-4 line-clamp-2 leading-relaxed">{project.deskripsi}</p>
                    <div className="flex items-center justify-between text-[13px]">
                      <span className="text-gray-500 font-medium">{project.mahasiswa?.user?.name}</span>
                      <span className="bg-gray-100 text-gray-700 px-3 py-1 rounded-full font-medium">{project.kategori}</span>
                    </div>
                  </div>
                </Link>
              ))}
            </div>
          ) : (
            <div className="bg-white rounded-2xl border border-gray-200/60 p-16 text-center">
              <div className="inline-flex items-center justify-center w-16 h-16 bg-gray-100 rounded-full mb-4">
                <svg className="w-8 h-8 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" strokeWidth="1.5">
                  <path strokeLinecap="round" strokeLinejoin="round" d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
                </svg>
              </div>
              <h3 className="text-[20px] font-semibold text-gray-900 mb-2">No Projects Yet</h3>
              <p className="text-[15px] text-gray-600">Be the first to create a project</p>
            </div>
          )}
        </div>

        {user?.role === 'Mahasiswa' && (
          <div className="bg-white rounded-2xl border border-gray-200/60 p-8">
            <h3 className="text-[22px] font-semibold text-gray-900 mb-6">Quick Actions</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Link to="/my-projects/new" className="bg-blue-600 text-white px-6 py-4 rounded-xl hover:bg-blue-700 transition-colors text-center font-medium text-[15px]">
                Create New Project
              </Link>
              <Link to="/my-projects" className="bg-gray-100 text-gray-900 px-6 py-4 rounded-xl hover:bg-gray-200 transition-colors text-center font-medium text-[15px]">
                Manage My Projects
              </Link>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
