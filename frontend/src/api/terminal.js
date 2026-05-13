import axios from 'axios';

const TERMINAL_URL = import.meta.env.VITE_API_URL || 'https://localhost:8888/api/v1';
const TERMINAL_TOKEN = import.meta.env.VITE_TERMINAL_TOKEN || 'terminal-secret-token';

const terminalApi = axios.create({
  baseURL: TERMINAL_URL,
  headers: {
    'Content-Type': 'application/json',
    'X-Terminal-Token': TERMINAL_TOKEN,
  },
});

export const terminalAuthApi = {
  authorize: (data) => terminalApi.post('/terminal/auth', data),
  downloadKeys: () => terminalApi.get('/terminal/keys'),
};