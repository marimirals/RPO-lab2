import { useState, useEffect } from 'react';
import { terminalsApi } from '../api/terminals';

export default function Terminals() {
  const [terminals, setTerminals] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [form, setForm] = useState({
    serial_number: '',
    name: '',
    address: '',
    location: '',
    is_active: true,
  });
  const [editingId, setEditingId] = useState(null);

  // 🔥 Объявляем безопасный список ЗДЕСЬ, до return
  const terminalsList = terminals || [];

  useEffect(() => {
    loadTerminals();
  }, []);

  const loadTerminals = async () => {
    try {
      const response = await terminalsApi.getAll();
      const data = response?.data;
      // Безопасное извлечение массива
      const list = Array.isArray(data?.data) ? data.data : (Array.isArray(data) ? data : []);
      setTerminals(list);
    } catch (err) {
      console.error('Load terminals error:', err);
      setError(err.response?.data?.error || 'Ошибка загрузки терминалов');
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    try {
      if (editingId) {
        await terminalsApi.update(editingId, form);
      } else {
        await terminalsApi.create(form);
      }
      setForm({ serial_number: '', name: '', address: '', location: '', is_active: true });
      setEditingId(null);
      loadTerminals();
    } catch (err) {
      setError(err.response?.data?.error || 'Ошибка сохранения');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Удалить терминал?')) return;
    try {
      await terminalsApi.delete(id);
      loadTerminals();
    } catch (err) {
      setError('Ошибка удаления');
    }
  };

  if (loading) return <div style={{ padding: '2rem' }}>Загрузка...</div>;

  return (
    <div style={{ padding: '2rem', display: 'flex', gap: '2rem' }}>
      {/* Форма */}
      <div style={{ flex: 1, maxWidth: '400px' }}>
        <h3>{editingId ? 'Редактировать терминал' : 'Новый терминал'}</h3>
        {error && <p style={{ color: 'red', background: '#ffe0e0', padding: '0.5rem', borderRadius: '4px' }}>{error}</p>}
        <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
          <div>
            <label style={{ display: 'block', marginBottom: '0.25rem', fontWeight: '500' }}>Серийный номер *</label>
            <input type="text" placeholder="Например: TERM-001" value={form.serial_number} onChange={(e) => setForm({ ...form, serial_number: e.target.value })} required style={{ padding: '0.5rem', width: '100%', boxSizing: 'border-box' }} />
          </div>
          <div>
            <label style={{ display: 'block', marginBottom: '0.25rem', fontWeight: '500' }}>Название *</label>
            <input type="text" placeholder="Например: Метро Проспект Мира" value={form.name} onChange={(e) => setForm({ ...form, name: e.target.value })} required style={{ padding: '0.5rem', width: '100%', boxSizing: 'border-box' }} />
          </div>
          <div>
            <label style={{ display: 'block', marginBottom: '0.25rem', fontWeight: '500' }}>Адрес</label>
            <input type="text" placeholder="Например: ул. Примерная, д. 1" value={form.address} onChange={(e) => setForm({ ...form, address: e.target.value })} style={{ padding: '0.5rem', width: '100%', boxSizing: 'border-box' }} />
          </div>
          <div>
            <label style={{ display: 'block', marginBottom: '0.25rem', fontWeight: '500' }}>Локация</label>
            <input type="text" placeholder="Например: вестибюль 1" value={form.location} onChange={(e) => setForm({ ...form, location: e.target.value })} style={{ padding: '0.5rem', width: '100%', boxSizing: 'border-box' }} />
          </div>
          <label style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', marginTop: '0.5rem' }}>
            <input type="checkbox" checked={form.is_active} onChange={(e) => setForm({ ...form, is_active: e.target.checked })} />
            <span>Активен</span>
          </label>
          <button type="submit" style={{ padding: '0.75rem', cursor: 'pointer', marginTop: '1rem' }}>
            {editingId ? '💾 Сохранить' : '➕ Создать'}
          </button>
          {editingId && (
            <button type="button" onClick={() => { setEditingId(null); setForm({ serial_number: '', name: '', address: '', location: '', is_active: true }); }} style={{ padding: '0.5rem', cursor: 'pointer', marginTop: '0.5rem', background: '#6c757d', color: 'white', border: 'none', borderRadius: '4px' }}>
              ✖️ Отмена
            </button>
          )}
        </form>
      </div>

      {/* Список */}
      <div style={{ flex: 2 }}>
        <h3>📋 Список терминалов</h3>
        {terminalsList.length === 0 ? (
          <p style={{ color: '#6c757d' }}>Нет терминалов. Создайте первый!</p>
        ) : (
          <div style={{ overflowX: 'auto' }}>
            <table style={{ width: '100%', borderCollapse: 'collapse', background: 'white', borderRadius: '8px', overflow: 'hidden', boxShadow: '0 2px 4px rgba(0,0,0,0.1)' }}>
              <thead>
                <tr style={{ background: '#f8f9fa' }}>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Серийный №</th>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Название</th>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Адрес</th>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Статус</th>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Действия</th>
                </tr>
              </thead>
              <tbody>
                {terminalsList.map((t) => (
                  <tr key={t.id} style={{ borderBottom: '1px solid #eee' }}>
                    <td style={{ padding: '0.75rem 1rem', fontFamily: 'monospace' }}>{t.serial_number}</td>
                    <td style={{ padding: '0.75rem 1rem' }}>{t.name}</td>
                    <td style={{ padding: '0.75rem 1rem' }}>{t.address || '—'}</td>
                    <td style={{ padding: '0.75rem 1rem' }}>{t.is_active ? '✅' : '❌'}</td>
                    <td style={{ padding: '0.75rem 1rem', display: 'flex', gap: '0.5rem' }}>
                      <button onClick={() => { setForm(t); setEditingId(t.id); }} style={{ padding: '0.25rem 0.5rem', cursor: 'pointer', fontSize: '0.9rem' }}>✏️</button>
                      <button onClick={() => handleDelete(t.id)} style={{ padding: '0.25rem 0.5rem', cursor: 'pointer', fontSize: '0.9rem', background: '#e74c3c', color: 'white', border: 'none', borderRadius: '4px' }}>🗑️</button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}