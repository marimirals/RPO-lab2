import { useTheme } from '../../hooks/useTheme';

export default function ThemeToggle({ size = 'default' }) {
  const { theme, toggleTheme } = useTheme();

  return (
    <button
      className={`theme-toggle theme-toggle--${size}`}
      onClick={toggleTheme}
      title={theme === 'light' ? 'Тёмная тема' : 'Светлая тема'}
      style={{
        padding: size === 'small' ? '0.3rem' : '0.5rem',
        background: 'transparent',
        border: '1px solid var(--border)',
        borderRadius: 'var(--radius-sm)',
        cursor: 'pointer',
        fontSize: size === 'small' ? '1rem' : '1.2rem',
        color: 'var(--text-primary)',
        transition: 'all var(--transition)',
      }}
    >
      {theme === 'light' ? '🌙' : '☀️'}
    </button>
  );
}