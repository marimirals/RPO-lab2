import { useState, useEffect } from 'react';
import { cardsApi } from '../api/cards';

export default function Cards() {
  const [cards, setCards] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [form, setForm] = useState({
    card_number: '',
    balance: 0,
    owner_name: '',
    key_id: null,
  });
  const [editingId, setEditingId] = useState(null);

  useEffect(() => {
    loadCards();
  }, []);

  const loadCards = async () => {
    try {
      const response = await cardsApi.getAll();
      // 🔥 Безопасное извлечение данных
      const data = response?.data;
      const cardsList = Array.isArray(data?.data) ? data.data : (Array.isArray(data) ? data : []);
      setCards(cardsList);
    } catch (err) {
      console.error('Load cards error:', err);
      setError(err.response?.data?.error || 'Ошибка загрузки карт');
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    try {
      if (editingId) {
        await cardsApi.update(editingId, form);
      } else {
        await cardsApi.create(form);
      }
      setForm({ card_number: '', balance: 0, owner_name: '', key_id: null });
      setEditingId(null);
      loadCards();
    } catch (err) {
      setError(err.response?.data?.error || 'Ошибка сохранения');
    }
  };

  const handleEdit = (card) => {
    setForm(card);
    setEditingId(card.id);
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Удалить карту?')) return;
    try {
      await cardsApi.delete(id);
      loadCards();
    } catch (err) {
      setError('Ошибка удаления');
    }
  };

  // 🔥 Защита от null в рендере
  const cardsList = cards || [];

  if (loading) return <div style={{ padding: '2rem' }}>Загрузка...</div>;

  return (
    <div style={{ padding: '2rem', display: 'flex', gap: '2rem' }}>
      {/* Форма */}
      <div style={{ flex: 1, maxWidth: '400px' }}>
        <h3>{editingId ? 'Редактировать карту' : 'Новая карта'}</h3>
        {error && <p style={{ color: 'red', background: '#ffe0e0', padding: '0.5rem', borderRadius: '4px' }}>{error}</p>}
        <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
          <div>
            <label style={{ display: 'block', marginBottom: '0.25rem', fontWeight: '500' }}>Номер карты *</label>
            <input
              type="text"
              placeholder="Например: 1234567890123456"
              value={form.card_number}
              onChange={(e) => setForm({ ...form, card_number: e.target.value })}
              required
              style={{ padding: '0.5rem', width: '100%', boxSizing: 'border-box' }}
            />
          </div>
          <div>
            <label style={{ display: 'block', marginBottom: '0.25rem', fontWeight: '500' }}>Баланс (копейки)</label>
            <input
              type="number"
              placeholder="Например: 1000"
              value={form.balance}
              onChange={(e) => setForm({ ...form, balance: parseInt(e.target.value) || 0 })}
              style={{ padding: '0.5rem', width: '100%', boxSizing: 'border-box' }}
            />
          </div>
          <div>
            <label style={{ display: 'block', marginBottom: '0.25rem', fontWeight: '500' }}>Имя владельца</label>
            <input
              type="text"
              placeholder="Например: Иван Иванов"
              value={form.owner_name}
              onChange={(e) => setForm({ ...form, owner_name: e.target.value })}
              style={{ padding: '0.5rem', width: '100%', boxSizing: 'border-box' }}
            />
          </div>
          <button type="submit" style={{ padding: '0.75rem', cursor: 'pointer', marginTop: '0.5rem' }}>
            {editingId ? '💾 Сохранить' : '➕ Создать'}
          </button>
          {editingId && (
            <button type="button" onClick={() => { setEditingId(null); setForm({ card_number: '', balance: 0, owner_name: '', key_id: null }); }} style={{ padding: '0.5rem', cursor: 'pointer', marginTop: '0.5rem', background: '#6c757d' }}>
              ✖️ Отмена
            </button>
          )}
        </form>
      </div>

      {/* Список */}
      <div style={{ flex: 2 }}>
        <h3>📋 Список карт</h3>
        {cardsList.length === 0 ? (
          <p style={{ color: '#6c757d' }}>Нет карт. Создайте первую!</p>
        ) : (
          <div style={{ overflowX: 'auto' }}>
            <table style={{ width: '100%', borderCollapse: 'collapse', background: 'white', borderRadius: '8px', overflow: 'hidden', boxShadow: '0 2px 4px rgba(0,0,0,0.1)' }}>
              <thead>
                <tr style={{ background: '#f8f9fa' }}>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Номер</th>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Баланс</th>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Владелец</th>
                  <th style={{ padding: '0.75rem 1rem', textAlign: 'left', borderBottom: '1px solid #dee2e6' }}>Действия</th>
                </tr>
              </thead>
              <tbody>
                {cardsList.map((card) => (
                  <tr key={card.id} style={{ borderBottom: '1px solid #eee' }}>
                    <td style={{ padding: '0.75rem 1rem', fontFamily: 'monospace' }}>{card.card_number}</td>
                    <td style={{ padding: '0.75rem 1rem' }}>{card.balance} коп.</td>
                    <td style={{ padding: '0.75rem 1rem' }}>{card.owner_name || '—'}</td>
                    <td style={{ padding: '0.75rem 1rem', display: 'flex', gap: '0.5rem' }}>
                      <button onClick={() => handleEdit(card)} style={{ padding: '0.25rem 0.5rem', cursor: 'pointer', fontSize: '0.9rem' }}>✏️</button>
                      <button onClick={() => handleDelete(card.id)} style={{ padding: '0.25rem 0.5rem', cursor: 'pointer', fontSize: '0.9rem', background: '#e74c3c' }}>🗑️</button>
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