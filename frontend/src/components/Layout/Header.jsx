import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';
import { useTheme } from '../../hooks/useTheme';
import ThemeToggle from '../UI/ThemeToggle';
import BurgerMenu from './BurgerMenu';
import '../../styles/components/header.css';

export default function Header({ showBurger = true }) {
  const { user, logout } = useAuth();
  const { theme } = useTheme();
  const navigate = useNavigate();
  const [isBurgerOpen, setIsBurgerOpen] = useState(false);

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <>
      <header className="header">
        <Link to="/" className="header-logo">
          Transport Card Auth
        </Link>
        
        <div className="header-actions">
          <ThemeToggle size="small" />
          
          {user && (
            <div className="header-user">
              <span>{user.name}</span>
              {user.is_admin && <span className="admin-badge">👑</span>}
            </div>
          )}
          
          {user && (
            <button className="header-btn secondary" onClick={handleLogout}>
              Выйти
            </button>
          )}
          
          {showBurger && (
            <button 
              className="burger-toggle" 
              onClick={() => setIsBurgerOpen(true)}
              aria-label="Открыть меню"
            >
              <span></span>
              <span></span>
              <span></span>
            </button>
          )}
        </div>
      </header>
      
      <BurgerMenu isOpen={isBurgerOpen} onClose={() => setIsBurgerOpen(false)} />
    </>
  );
}