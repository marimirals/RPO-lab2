import { NavLink, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { useTheme } from '../../hooks/useTheme';
import ThemeToggle from '../UI/ThemeToggle';

export default function BurgerMenu({ isOpen, onClose }) {
  const { user, logout } = useAuth();
  const { theme, toggleTheme } = useTheme();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    onClose();
    navigate('/login');
  };

  const getInitial = (name) => name?.charAt(0).toUpperCase() || '?';

  if (!isOpen) return null;

  return (
    <>
      <div className="burger-overlay" onClick={onClose} />
      <nav className="burger-menu open">
        <div className="burger-header">
          <span className="burger-title">Меню</span>
          <button className="burger-close" onClick={onClose}>&times;</button>
        </div>
        
        <div className="burger-body">
          <div className="burger-section">
            <div className="burger-section-title">Навигация</div>
            <div className="burger-nav">
              <NavLink to="/" onClick={onClose}>
                <span className="icon">📊</span>
                Главная
              </NavLink>
              <NavLink to="/cards" onClick={onClose}>
                <span className="icon">💳</span>
                Карты
              </NavLink>
              <NavLink to="/terminals" onClick={onClose}>
                <span className="icon">🏪</span>
                Терминалы
              </NavLink>
              <NavLink to="/transactions" onClick={onClose}>
                <span className="icon">📋</span>
                Транзакции
              </NavLink>
              {user?.is_admin && (
                <>
                  <NavLink to="/keys" onClick={onClose}>
                    <span className="icon">🔑</span>
                    Ключи
                  </NavLink>
                  <NavLink to="/users" onClick={onClose}>
                    <span className="icon">👥</span>
                    Пользователи
                  </NavLink>
                </>
              )}
            </div>
          </div>
          
          <div className="burger-section">
            <div className="burger-section-title">Настройки</div>
            <div className="burger-nav">
              <NavLink to="/profile" onClick={onClose}>
                <span className="icon">👤</span>
                Профиль
              </NavLink>
              <button 
                className="burger-nav-item" 
                onClick={() => { toggleTheme(); onClose(); }}
                style={{ 
                  display: 'flex', 
                  alignItems: 'center', 
                  gap: '0.75rem', 
                  padding: '0.6rem 0.8rem', 
                  background: 'transparent', 
                  border: 'none', 
                  color: 'var(--text-primary)', 
                  borderRadius: 'var(--radius-sm)',
                  cursor: 'pointer',
                  width: '100%',
                  textAlign: 'left'
                }}
              >
                <span className="icon">{theme === 'light' ? '🌙' : '☀️'}</span>
                Тема: {theme === 'light' ? 'Светлая' : 'Тёмная'}
              </button>
            </div>
          </div>
        </div>
        
        <div className="burger-footer">
          <div className="burger-user">
            <div className="burger-avatar">{getInitial(user?.name)}</div>
            <div className="burger-user-info">
              <div className="burger-user-name">{user?.name}</div>
              <div className="burger-user-role">
                {user?.is_admin ? 'Администратор' : 'Пользователь'}
              </div>
            </div>
          </div>
          <button className="burger-logout danger" onClick={handleLogout}>
            Выйти
          </button>
        </div>
      </nav>
    </>
  );
}