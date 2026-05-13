import { useState } from 'react';
import { transactionsApi } from '../api/transactions';

export default function Transactions() {
  const [cardId, setCardId] = useState('');
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  const handleSearch = async (e) => {
    e.preventDefault();
    if (!cardId) return;
    
    setLoading(true);
    setError('');
    try {
      const { data } = await transactionsApi.getByCard(cardId);
      setTransactions(data.success ? data.data : []);
    } catch (err) {
      setError(err.response?.data?.error || 'Ошибка поиска');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ padding: '2rem' }}>
      <h3>📋 Транзакции</h3>
      
      <form onSubmit={handleSearch} style={{ marginBottom: '2rem', display: 'flex', gap: '1rem', maxWidth: '500px' }}>
        <input
          type="number"
          placeholder="ID карты для поиска"
          value={cardId}
          onChange={(e) => setCardId(e.target.value)}
          style={{ flex: 1, padding: '0.5rem' }}
        />
        <button type="submit" style={{ padding: '0.5rem 1rem' }}>
          {loading ? 'Поиск...' : '🔍 Найти'}
        </button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      {transactions.length > 0 && (
        <table style={{ width: '100%', borderCollapse: 'collapse' }}>
          <thead>
            <tr style={{ background: '#f8f9fa' }}>
              <th style={{ padding: '0.75rem', textAlign: 'left', border: '1px solid #dee2e6' }}>ID</th>
              <th style={{ padding: '0.75rem', textAlign: 'left', border: '1px solid #dee2e6' }}>Сумма</th>
              <th style={{ padding: '0.75rem', textAlign: 'left', border: '1px solid #dee2e6' }}>Тип</th>
              <th style={{ padding: '0.75rem', textAlign: 'left', border: '1px solid #dee2e6' }}>Статус</th>
              <th style={{ padding: '0.75rem', textAlign: 'left', border: '1px solid #dee2e6' }}>Время</th>
            </tr>
          </thead>
          <tbody>
            {transactions.map((t) => (
              <tr key={t.id}>
                <td style={{ padding: '0.75rem', border: '1px solid #dee2e6' }}>{t.id}</td>
                <td style={{ padding: '0.75rem', border: '1px solid #dee2e6' }}>{t.amount} коп.</td>
                <td style={{ padding: '0.75rem', border: '1px solid #dee2e6' }}>{t.transaction_type}</td>
                <td style={{ padding: '0.75rem', border: '1px solid #dee2e6' }}>{t.status}</td>
                <td style={{ padding: '0.75rem', border: '1px solid #dee2e6' }}>{new Date(t.transaction_time).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}