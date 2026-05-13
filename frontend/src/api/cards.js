import api from './client';

export const cardsApi = {
  getAll: async () => {
    const response = await api.get('/cards');
    // Бэкенд может вернуть {success: true, data: [...]} или просто [...]
    return response;
  },
  getById: (id) => api.get(`/cards/${id}`),
  getByNumber: (number) => api.get(`/cards/number/${number}`),
  create: (data) => api.post('/cards', data),
  update: (id, data) => api.put(`/cards/${id}`, data),
  delete: (id) => api.delete(`/cards/${id}`),
};