import api from './client';

export const transactionsApi = {
  getById: (id) => api.get(`/transactions/${id}`),
  getByCard: (cardId) => api.get(`/transactions/card/${cardId}`),
  updateStatus: (id, status) => api.put(`/transactions/${id}/status?status=${status}`),
};