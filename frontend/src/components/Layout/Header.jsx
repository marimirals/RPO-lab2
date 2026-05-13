import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';

export default function Header() {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <header style={{ 
      padding: '1rem 2rem', 
      background: '#2c3e50', 
      color: 'white', 
      display: 'flex', 
      justifyContent: 'space-between',
      alignItems: 'center'
    }}>
      <Link to="/" style={{ color: 'white', textDecoration: 'none', fontSize: '1.25rem', fontWeight: 'bold' }}>
        🚇 Transport Card Auth
      </Link>
      
      <div style={{ display: 'flex', gap: '1rem', alignItems: 'center' }}>
        {user && <span>👤 {user.name} {user.is_admin && '👑'}</span>}
        {user && (
          <button onClick={handleLogout} style={{ 
            padding: '0.5rem 1rem', 
            background: '#e74c3c', 
            color: 'white', 
            border: 'none', 
            borderRadius: '4px',
            cursor: 'pointer'
          }}>
            Выйти
          </button>
        )}
      </div>
    </header>
  );
}