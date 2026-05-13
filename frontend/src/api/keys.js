import api from './client';

export const keysApi = {
  getAll: () => api.get('/keys'),
  getById: (id) => api.get(`/keys/${id}`),
  create: (data) => api.post('/keys', data),
  update: (id, data) => api.put(`/keys/${id}`, data),
  delete: (id) => api.delete(`/keys/${id}`),
};