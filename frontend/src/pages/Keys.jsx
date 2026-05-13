import { useState, useEffect } from 'react';
import { keysApi } from '../api/keys';
import { useAuth } from '../contexts/AuthContext';
import { Navigate } from 'react-router-dom';

export default function Keys() {
  const { user } = useAuth();
  const [keys, setKeys] = useState([]);
  const [loading, setLoading] = useState(true);
  const [form, setForm] = useState({ key_value: '', description: '' });

  useEffect(() => {
    if (user?.is_admin) loadKeys();
  }, [user]);

  const loadKeys = async () => {
    try {
      const { data } = await keysApi.getAll();
      setKeys(data.success ? data.data : []);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await keysApi.create(form);
      setForm({ key_value: '', description: '' });
      loadKeys();
    } catch (err) {
      alert(err.response?.data?.error || 'Ошибка');
    }
  };

  if (!user?.is_admin) return <Navigate to="/" />;
  if (loading) return <div style={{ padding: '2rem' }}>Загрузка...</div>;

  return (
    <div style={{ padding: '2rem' }}>
      <h3>Управление ключами</h3>
      <form onSubmit={handleSubmit} style={{ display: 'flex', gap: '1rem', marginBottom: '2rem' }}>
        <input type="text" placeholder="Значение ключа *" value={form.key_value} onChange={(e) => setForm({ ...form, key_value: e.target.value })} required style={{ padding: '0.5rem' }} />
        <input type="text" placeholder="Описание" value={form.description} onChange={(e) => setForm({ ...form, description: e.target.value })} style={{ padding: '0.5rem' }} />
        <button type="submit" style={{ padding: '0.5rem 1rem', cursor: 'pointer' }}>Добавить</button>
      </form>
      <ul>
        {keys.map((key) => (
          <li key={key.id} style={{ padding: '0.5rem', borderBottom: '1px solid #eee' }}>
            <strong>{key.key_value}</strong> — {key.description || 'без описания'}
          </li>
        ))}
      </ul>
    </div>
  );
}