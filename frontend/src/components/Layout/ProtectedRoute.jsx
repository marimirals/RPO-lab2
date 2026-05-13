import { Navigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';

export default function ProtectedRoute({ children, adminOnly = false }) {
  const { user, loading } = useAuth();

  if (loading) return <div style={{ padding: '2rem' }}>Загрузка...</div>;
  if (!user) return <Navigate to="/login" />;
  if (adminOnly && !user.is_admin) return <Navigate to="/" />;

  return children;
}