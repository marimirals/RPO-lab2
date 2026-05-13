import { useState } from 'react';
import { terminalAuthApi } from '../api/terminal';

export default function TerminalTest() {
  const [form, setForm] = useState({ card_number: '', amount: 0, terminal_id: 1 });
  const [result, setResult] = useState(null);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setResult(null);
    try {
      const { data } = await terminalAuthApi.authorize(form);
      setResult(data);
    } catch (err) {
      setError(err.response?.data?.error || 'Ошибка авторизации');
    }
  };

  return (
    <div style={{ padding: '2rem', maxWidth: '500px', margin: '0 auto' }}>
      <h3>🧪 Тест авторизации платежа</h3>
      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
        <div>
          <label style={{ display: 'block', marginBottom: '0.5rem', fontWeight: 'bold' }}>
            Номер карты:
          </label>
          <input 
            type="text" 
            placeholder="Например: 1234567890123456" 
            value={form.card_number} 
            onChange={(e) => setForm({ ...form, card_number: e.target.value })} 
            required 
            style={{ padding: '0.5rem', width: '100%' }} 
          />
        </div>
        
        <div>
          <label style={{ display: 'block', marginBottom: '0.5rem', fontWeight: 'bold' }}>
            Сумма платежа (копейки):
          </label>
          <input 
            type="number" 
            placeholder="Например: 100" 
            value={form.amount} 
            onChange={(e) => setForm({ ...form, amount: parseInt(e.target.value) || 0 })} 
            required 
            style={{ padding: '0.5rem', width: '100%' }} 
          />
        </div>
        
        <div>
          <label style={{ display: 'block', marginBottom: '0.5rem', fontWeight: 'bold' }}>
            ID терминала:
          </label>
          <input 
            type="number" 
            placeholder="Например: 1" 
            value={form.terminal_id} 
            onChange={(e) => setForm({ ...form, terminal_id: parseInt(e.target.value) || 1 })} 
            style={{ padding: '0.5rem', width: '100%' }} 
          />
        </div>
        
        <button type="submit" style={{ padding: '0.75rem', cursor: 'pointer', marginTop: '1rem' }}>
          Авторизовать
        </button>
      </form>
      
      {error && <p style={{ color: 'red', background: '#ffe0e0', padding: '0.5rem', borderRadius: '4px' }}>{error}</p>}
      
      {result && (
        <div style={{ marginTop: '1rem', padding: '1rem', background: result.authorized ? '#d4edda' : '#f8d7da', borderRadius: '4px' }}>
          <strong>Результат:</strong> {result.authorized ? '✅ Одобрено' : '❌ Отклонено'}
          {result.message && <p>{result.message}</p>}
          {result.balance !== undefined && <p>Остаток на карте: {result.balance} коп.</p>}
        </div>
      )}
    </div>
  );
}