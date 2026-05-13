import api from './client';

export const terminalsApi = {
  getAll: () => api.get('/terminals'),
  getById: (id) => api.get(`/terminals/${id}`),
  create: (data) => api.post('/terminals', data),
  update: (id, data) => api.put(`/terminals/${id}`, data),
  delete: (id) => api.delete(`/terminals/${id}`),
};