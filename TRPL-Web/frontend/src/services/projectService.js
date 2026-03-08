import api from './api';

export const authService = {
  getGoogleAuthURL: async () => {
    const response = await api.get('/auth/google');
    return response.data.url;
  },

  getCurrentUser: async () => {
    const response = await api.get('/auth/me');
    return response.data.user;
  },

  logout: () => {
    localStorage.removeItem('token');
    window.location.href = '/login';
  },
};

export const projectService = {
  getProjects: async (filters = {}) => {
    const response = await api.get('/projects', { params: filters });
    return response.data.projects;
  },

  getProjectById: async (id) => {
    const response = await api.get(`/projects/${id}`);
    return response.data.project;
  },

  createProject: async (data) => {
    const response = await api.post('/projects', data);
    return response.data.project;
  },

  updateProject: async (id, data) => {
    const response = await api.put(`/projects/${id}`, data);
    return response.data.project;
  },

  deleteProject: async (id) => {
    await api.delete(`/projects/${id}`);
  },

  getMyProjects: async () => {
    const response = await api.get('/my-projects');
    return response.data.projects;
  },

  getStats: async () => {
    const response = await api.get('/projects/stats');
    return response.data;
  },
};
