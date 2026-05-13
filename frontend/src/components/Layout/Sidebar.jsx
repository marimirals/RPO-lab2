import { NavLink } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';

export default function Sidebar() {
  const { user } = useAuth();

  const navStyle = {
    display: 'flex',
    flexDirection: 'column',
    padding: '1rem',
    gap: '0.5rem',
    minWidth: '200px',
    background: '#f8f9fa',
    borderRight: '1px solid #dee2e6',
  };

  const linkStyle = ({ isActive }) => ({
    padding: '0.75rem 1rem',
    textDecoration: 'none',
    color: isActive ? '#2c3e50' : '#6c757d',
    background: isActive ? '#e9ecef' : 'transparent',
    borderRadius: '4px',
    fontWeight: isActive ? 'bold' : 'normal',
  });

  return (
    <nav style={navStyle}>
      <NavLink to="/" style={linkStyle}>📊 Главная</NavLink>
      <NavLink to="/cards" style={linkStyle}>💳 Карты</NavLink>
      <NavLink to="/terminals" style={linkStyle}>🏪 Терминалы</NavLink>
      <NavLink to="/transactions" style={linkStyle}>📋 Транзакции</NavLink>
      
      {user?.is_admin && (
        <>
          <hr style={{ margin: '0.5rem 0', border: 'none', borderTop: '1px solid #dee2e6' }} />
          <NavLink to="/keys" style={linkStyle}>🔑 Ключи</NavLink>
          <NavLink to="/users" style={linkStyle}>👥 Пользователи</NavLink>
        </>
      )}
      
      <hr style={{ margin: '0.5rem 0', border: 'none', borderTop: '1px solid #dee2e6' }} />
      <NavLink to="/terminal-test" style={linkStyle}>🧪 Тест терминала</NavLink>
    </nav>
  );
}