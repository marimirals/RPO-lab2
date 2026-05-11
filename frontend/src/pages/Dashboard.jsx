import { useAuth } from '../contexts/AuthContext';
import { useNavigate } from 'react-router-dom';

export default function Dashboard() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  return (
    <div style={{ padding: '2rem' }}>
      <h2>Добро пожаловать, {user?.name}!</h2>
      <p>Логин: {user?.login} | Роль: {user?.is_admin ? 'Администратор' : 'Пользователь'}</p>
      <button 
        onClick={() => { logout(); navigate('/login'); }}
        style={{ padding: '0.5rem 1rem', marginTop: '1rem', cursor: 'pointer' }}
      >
        Выйти
      </button>
    </div>
  );
}